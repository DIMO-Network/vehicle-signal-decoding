package api

import (
	"context"

	"github.com/pkg/errors"

	//"errors"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"strings"
	"syscall"

	usersapi "github.com/DIMO-Network/users-api/pkg/grpc"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/utils"
	jwtware "github.com/gofiber/contrib/jwt"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/gateways"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/middleware/owner"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/controllers"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/services"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/shared/middleware/metrics"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"

	fiberrecover "github.com/gofiber/fiber/v2/middleware/recover"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
)

func Run(ctx context.Context, logger zerolog.Logger, settings *config.Settings) {
	//db
	pdb := db.NewDbConnectionFromSettings(ctx, &settings.DB, true)
	pdb.WaitForDB(logger)

	s3Client := getS3ServiceClient(ctx, settings, logger)
	conns, err := getGRPCConnections(settings)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to dial connection to at least one grpc service")
	}

	go StartGrpcServer(logger, pdb.DBS, settings, s3Client)
	startWebAPI(logger, settings, pdb, conns)
	startMonitoringServer(logger, settings)

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent with length of 1
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel
	<-c                                             // This blocks the main thread until an interrupt is received
	logger.Info().Msg("Gracefully shutting down and running cleanup tasks...")
	_ = ctx.Done()
	_ = pdb.DBS().Writer.Close()
	_ = pdb.DBS().Reader.Close()
	_ = conns.usersAPIConn.Close()
	_ = conns.definitionsAPIConn.Close()
	_ = conns.devicesAPIConn.Close()
}

func startMonitoringServer(logger zerolog.Logger, settings *config.Settings) {
	logger = logger.With().Str("fiber-app", "monitoring").Logger()
	monApp := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return ErrorHandler(c, err, logger)
		},
		DisableStartupMessage: true,
	})
	monApp.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("healthy")
	})
	monApp.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	go func() {
		// 8888 is our standard port for exposing metrics in DIMO infra
		if err := monApp.Listen(":" + settings.MonitoringPort); err != nil {
			logger.Fatal().Err(err).Str("port", settings.MonitoringPort).Msg("Failed to start monitoring web server.")
		}
	}()

	logger.Info().Str("port", settings.MonitoringPort).Msg("started monitoring web server.")
}

func startWebAPI(logger zerolog.Logger, settings *config.Settings, database db.Store, conns *grpcServiceConnections) *fiber.App {
	usersClient := usersapi.NewUserServiceClient(conns.usersAPIConn) // this is a raw grpc client
	// these are an abstractions over the grpc client
	userDeviceSvc := services.NewUserDevicesService(conns.devicesAPIConn)
	deviceDefsvc := services.NewDeviceDefinitionsService(conns.definitionsAPIConn)

	identityAPI := gateways.NewIdentityAPIService(&logger, settings, nil)
	deviceTemplatesvc := services.NewDeviceTemplateService(database.DBS().Writer.DB, logger, settings, identityAPI)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return ErrorHandler(c, err, logger)
		},
		DisableStartupMessage: true,
		ReadBufferSize:        16000,
	})

	jwtAuth := jwtware.New(jwtware.Config{
		JWKSetURLs: []string{settings.JwtKeySetURL},
		ErrorHandler: func(_ *fiber.Ctx, err error) error {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid JWT. "+err.Error())
		},
	})

	app.Use(metrics.HTTPMetricsMiddleware)

	app.Use(fiberrecover.New(fiberrecover.Config{
		Next:              nil,
		EnableStackTrace:  true,
		StackTraceHandler: nil,
	}))
	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("healthy")
	})
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("healthy")
	})

	deviceConfigController := controllers.NewDeviceConfigController(settings, &logger, database.DBS, userDeviceSvc,
		deviceDefsvc, deviceTemplatesvc, identityAPI)
	jobsController := controllers.NewJobsController(settings, &logger, database.DBS().Reader.DB, userDeviceSvc, deviceDefsvc)

	v1 := app.Group("/v1")

	deviceMw := owner.New(usersClient, userDeviceSvc, &logger)

	v1.Get("/swagger/*", swagger.HandlerDefault)

	// resolve what templates to use for my car/device
	v1.Get("/device-config/eth-addr/:ethAddr/urls", deviceConfigController.GetConfigURLsFromEthAddr)
	v1.Get("/device-config/vin/:vin/urls", deviceConfigController.GetConfigURLsFromVIN)

	// endpoints that serve up the template contents
	v1.Get("/device-config/pids/:templateName", deviceConfigController.GetPIDsByTemplate)
	v1.Get("/device-config/settings/:name", deviceConfigController.GetDeviceSettingsByName)
	v1.Get("/device-config/dbc/:templateName", deviceConfigController.GetDBCFileByTemplateName)
	// for backwards compatibility - remove after a month or 2
	v1.Get("/device-config/:name/device-settings", deviceConfigController.GetDeviceSettingsByName)

	// device to template and fw status
	v1.Get("/device-config/eth-addr/:ethAddr/status", deviceConfigController.GetConfigStatusByEthAddr)

	// EC recover authentication middleware
	etherSigAuth := func(c *fiber.Ctx) error {
		ethAddr := c.Params("ethAddr")

		// get signature from header
		signature := c.Get("Signature")
		if signature == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Signature not found in request header",
			})
		}

		ok, err := utils.VerifySignature(c.Body(), signature, ethAddr)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to recover an address from the signature: %s", ethAddr))
		} else if !ok {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		// If EC recover authentication succeeds, we should skip authorization middleware since
		// we already authenticated on behalf of the device
		return c.Next()
	}

	// Jwt authentication
	v1.Patch("/device-config/eth-addr/:ethAddr/status", jwtAuth, deviceMw, deviceConfigController.PatchConfigStatusByEthAddr)

	// Signature authentication
	v1.Patch("/device-config/eth-addr/:ethAddr/hw/status", etherSigAuth, deviceConfigController.PatchHwConfigStatusByEthAddr)

	// Jobs endpoints
	v1.Get("/device-config/eth-addr/:ethAddr/jobs", etherSigAuth, jobsController.GetJobsFromEthAddr)
	v1.Get("/device-config/eth-addr/:ethAddr/jobs/pending", etherSigAuth, jobsController.GetJobsPendingFromEthAddr)
	v1.Patch("/device-config/eth-addr/:ethAddr/jobs/:jobId/:status", etherSigAuth, jobsController.PatchJobsFromEthAddr)

	go func() {
		if err := app.Listen(":" + settings.Port); err != nil {
			logger.Fatal().Err(err).Str("port", settings.Port).Msg("Failed to start web server.")
		}
	}()

	logger.Info().Str("port", settings.Port).Msg("Started api web server")

	return app
}

// ErrorHandler handles errors returned from fiber handlers / controllers
func ErrorHandler(c *fiber.Ctx, err error, logger zerolog.Logger) error {
	// Default error info
	code := fiber.StatusInternalServerError
	message := "An error occurred while processing your request."

	var e *fiber.Error
	var evnf *exceptions.NotFoundError
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	} else if errors.As(err, &evnf) {
		code = fiber.StatusNotFound
		message = e.Message
	}

	logger.Err(err).Int("httpStatusCode", code).
		Str("path", strings.TrimPrefix(c.Path(), "/")).
		Str("method", c.Method()).
		Str("stack", string(debug.Stack())).Msg("Failed request.")

	return c.Status(code).JSON(CodeResp{Code: code, Message: message})
}

type CodeResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type grpcServiceConnections struct {
	devicesAPIConn     *grpc.ClientConn
	definitionsAPIConn *grpc.ClientConn
	usersAPIConn       *grpc.ClientConn
}

// getGRPCConnections establishes and returns the connections for our GRPC dependencies
func getGRPCConnections(settings *config.Settings) (*grpcServiceConnections, error) {
	devicesConn, err := grpc.NewClient(settings.DeviceGRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to dial at: %s", settings.DeviceGRPCAddr)
	}
	definitionsConn, err := grpc.NewClient(settings.DefinitionsGRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to dial at: %s", settings.DefinitionsGRPCAddr)
	}
	usersConn, err := grpc.NewClient(settings.UsersGRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to dial at: %s", settings.UsersGRPCAddr)
	}
	return &grpcServiceConnections{
		devicesAPIConn:     devicesConn,
		definitionsAPIConn: definitionsConn,
		usersAPIConn:       usersConn,
	}, nil
}

// getS3ServiceClient instantiates a new default config and then a new s3 services client if not already set. Takes context in, although it could likely use a context from container passed in on instantiation
func getS3ServiceClient(ctx context.Context, settings *config.Settings, logger zerolog.Logger) *s3.Client {
	cfg, err := awsconfig.LoadDefaultConfig(ctx,
		awsconfig.WithRegion(settings.AWSRegion))
	// Below code is used when working with locally running S3 simulator. These options have been deprecated and could not quickly find new way to do this.
	//awsconfig.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
	//	func(_, _ string, _ ...interface{}) (aws.Endpoint, error) {
	//
	//		if settings.Environment == "local" {
	//			return aws.Endpoint{PartitionID: "aws", URL: settings.CandumpsAWSEndpoint, SigningRegion: settings.AWSRegion}, nil // The SigningRegion key was what's was missing! D'oh.
	//		}
	//
	//		// returning EndpointNotFoundError will allow the service to fallback to its default resolution
	//		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	//	})))

	if err != nil {
		logger.Fatal().Err(err).Msg("Could not load aws config, terminating")
	}

	s3ServiceClient := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = settings.AWSRegion
		o.Credentials = credentials.NewStaticCredentialsProvider(settings.CandumpsAWSAccessKeyID, settings.CandumpsAWSSecretAccessKey, "")
	})

	return s3ServiceClient
}

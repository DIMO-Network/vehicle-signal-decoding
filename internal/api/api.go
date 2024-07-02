package api

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"strings"
	"syscall"

	pb "github.com/DIMO-Network/users-api/pkg/grpc"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/gateways"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/middleware/owner"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/utils"
	jwtware "github.com/gofiber/contrib/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/commands"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/controllers"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/services"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/kafka"
	"github.com/Shopify/sarama"

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

	go StartGrpcServer(logger, pdb.DBS, settings, s3Client)
	startWebAPI(logger, settings, pdb)
	startVehicleSignalConsumer(logger, settings, pdb)
	startMonitoringServer(logger, settings)

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent with length of 1
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel
	<-c                                             // This blocks the main thread until an interrupt is received
	logger.Info().Msg("Gracefully shutting down and running cleanup tasks...")
	_ = ctx.Done()
	_ = pdb.DBS().Writer.Close()
	_ = pdb.DBS().Reader.Close()
}

func startVehicleSignalConsumer(logger zerolog.Logger, settings *config.Settings, pdb db.Store) {

	if len(settings.KafkaBrokers) == 0 {
		return
	}

	clusterConfig := sarama.NewConfig()
	clusterConfig.Version = sarama.V2_8_1_0
	clusterConfig.Consumer.Offsets.Initial = sarama.OffsetNewest

	cfg := &kafka.Config{
		ClusterConfig:   clusterConfig,
		BrokerAddresses: strings.Split(settings.KafkaBrokers, ","),
		Topic:           settings.DBCDecodingTopic,
		GroupID:         "vehicle-signal-decoding",
		MaxInFlight:     int64(5),
	}
	consumer, err := kafka.NewConsumer(cfg, &logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not start credential update consumer")
	}

	userDeviceService := services.NewUserDeviceService(settings)
	handler := commands.NewRunTestSignalCommandHandler(pdb.DBS, logger, userDeviceService)
	service := NewWorkerListenerService(logger, handler)

	consumer.Start(context.Background(), service.ProcessWorker)

	logger.Info().Msg("Vehicle Signal Decoding consumer started")
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

	logger.Info().Str("port", settings.MonitoringPort).Msg("Started monitoring web server.")
}

func startWebAPI(logger zerolog.Logger, settings *config.Settings, database db.Store) *fiber.App {
	//Create gRPC connection
	userDeviceSvc := services.NewUserDeviceService(settings)
	deviceDefsvc := services.NewDeviceDefinitionsService(settings)
	deviceTemplatesvc := services.NewDeviceTemplateService(database.DBS().Writer.DB, deviceDefsvc, logger, settings)
	identityAPI := gateways.NewIdentityAPIService(&logger)
	// todo: this is messy - we open the connection but are never closing it, or wrapping this in a class that handles the connection for us
	usersClient := getUsersClient(logger, settings.UsersGRPCAddr)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return ErrorHandler(c, err, logger)
		},
		DisableStartupMessage: true,
		ReadBufferSize:        16000,
	})

	// secured paths
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

	deviceConfigController := controllers.NewDeviceConfigController(settings, &logger, database.DBS().Reader.DB, userDeviceSvc,
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

	// jwt authentication wrapper, which also calls another authentication method if jwt fails
	jwtAuthWrapper := func(c *fiber.Ctx) error {
		if err := jwtAuth(c); err != nil {
			// If JWT authentication fails, call the next middleware function
			return c.Next()
		}

		// If JWT authentication succeeds, don't call the next middleware function
		// Instead, skip to the handler after the second authentication middleware
		return deviceMw(c)
	}

	// EC recover authentication middleware
	ecRecoverAuth := func(c *fiber.Ctx) error {
		ethAddr := c.Params("ethAddr")

		// get signature from header
		signature := c.Get("Signature")
		if signature == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Signature not found in request header",
			})
		}

		ok, err := utils.VerifySignature(c.Body(), signature, ethAddr)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to recover an address from the signature: %s", ethAddr))
		} else if !ok {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		// If EC recover authentication succeeds, call the next middleware function
		return c.Next()
	}

	v1.Patch("/device-config/eth-addr/:ethAddr/status", jwtAuthWrapper, ecRecoverAuth, deviceMw, deviceConfigController.PatchConfigStatusByEthAddr)

	// Jobs endpoint
	v1.Get("/device-config/eth-addr/:ethAddr/jobs", jobsController.GetJobsFromEthAddr)
	v1.Get("/device-config/eth-addr/:ethAddr/jobs/pending", jobsController.GetJobsPendingFromEthAddr)
	v1.Patch("/device-config/eth-addr/:ethAddr/jobs/:jobId/:status", jobsController.PatchJobsFromEthAddr)

	go func() {
		if err := app.Listen(":" + settings.Port); err != nil {
			logger.Fatal().Err(err).Str("port", settings.Port).Msg("Failed to start monitoring web server.")
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

// getS3ServiceClient instantiates a new default config and then a new s3 services client if not already set. Takes context in, although it could likely use a context from container passed in on instantiation
func getS3ServiceClient(ctx context.Context, settings *config.Settings, logger zerolog.Logger) *s3.Client {
	cfg, err := awsconfig.LoadDefaultConfig(ctx,
		awsconfig.WithRegion(settings.AWSRegion),
		// Comment the below out if not using localhost
		awsconfig.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(_, _ string, _ ...interface{}) (aws.Endpoint, error) {

				if settings.Environment == "local" {
					return aws.Endpoint{PartitionID: "aws", URL: settings.CandumpsAWSEndpoint, SigningRegion: settings.AWSRegion}, nil // The SigningRegion key was what's was missing! D'oh.
				}

				// returning EndpointNotFoundError will allow the service to fallback to its default resolution
				return aws.Endpoint{}, &aws.EndpointNotFoundError{}
			})))

	if err != nil {
		logger.Fatal().Err(err).Msg("Could not load aws config, terminating")
	}

	s3ServiceClient := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = settings.AWSRegion
		o.Credentials = credentials.NewStaticCredentialsProvider(settings.CandumpsAWSAccessKeyID, settings.CandumpsAWSSecretAccessKey, "")
	})

	return s3ServiceClient
}

func getUsersClient(logger zerolog.Logger, usersAPIGRPCAddr string) pb.UserServiceClient {
	usersConn, err := grpc.Dial(usersAPIGRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal().Err(err).Msgf("Failed to dial users-api at %s", usersAPIGRPCAddr)
	}

	return pb.NewUserServiceClient(usersConn)
}

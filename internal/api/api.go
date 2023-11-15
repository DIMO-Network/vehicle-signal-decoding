package api

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"strings"
	"syscall"

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
)

func Run(ctx context.Context, logger zerolog.Logger, settings *config.Settings) {

	//db
	pdb := db.NewDbConnectionFromSettings(ctx, &settings.DB, true)
	pdb.WaitForDB(logger)

	go StartGrpcServer(logger, pdb.DBS, settings)

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

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return ErrorHandler(c, err, logger)
		},
		DisableStartupMessage: true,
		ReadBufferSize:        16000,
	})

	app.Use(metrics.HTTPMetricsMiddleware)

	app.Use(fiberrecover.New(fiberrecover.Config{
		Next:              nil,
		EnableStackTrace:  true,
		StackTraceHandler: nil,
	}))
	app.Use(cors.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("healthy")
	})

	deviceConfigController := controllers.NewDeviceConfigController(settings, &logger, database.DBS().Reader.DB, userDeviceSvc, deviceDefsvc)

	v1 := app.Group("/v1")

	v1.Get("/swagger/*", swagger.HandlerDefault)

	v1.Get("/device-config/vin/:vin/urls", deviceConfigController.GetConfigURLsFromVIN)
	v1.Get("/device-config/eth-addr/:ethAddr/urls", deviceConfigController.GetConfigURLsFromEthAddr)

	v1.Get("/device-config/:templateName/pids", deviceConfigController.GetPIDsByTemplate)
	v1.Get("/device-config/:templateName/device-settings", deviceConfigController.GetDeviceSettingsByTemplate)
	v1.Get("/device-config/:templateName/dbc", deviceConfigController.GetDBCFileByTemplateName)

	go func() {
		if err := app.Listen(":" + settings.Port); err != nil {
			logger.Fatal().Err(err).Str("port", settings.Port).Msg("Failed to start monitoring web server.")
		}
	}()

	logger.Info().Str("port", settings.Port).Msg("Started api web server")

	return app
}

// Code below copied from device-data-api/main.go
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

	logger.Err(err).Int("code", code).Str("path", strings.TrimPrefix(c.Path(), "/")).Msg("Failed request.")

	return c.Status(code).JSON(CodeResp{Code: code, Message: message})
}

type CodeResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

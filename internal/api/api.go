package api

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/commands"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/services"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/kafka"
	"github.com/Shopify/sarama"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"

	"github.com/DIMO-Network/device-data-api/internal/middleware/metrics"
	fiberrecover "github.com/gofiber/fiber/v2/middleware/recover"
)

type Response struct {
	PIDURL   string `json:"pidsUrl"`
	PowerURL string `json:"powerUrl"`
	DBCURL   string `json:"dbcUrl"`
}

func Run(ctx context.Context, logger zerolog.Logger, settings *config.Settings) {

	//db
	pdb := db.NewDbConnectionFromSettings(ctx, &settings.DB, true)
	pdb.WaitForDB(logger)

	go StartGrpcServer(logger, pdb.DBS, settings)

	startMonitoringServer(logger, settings)
	startVehicleSignalConsumer(logger, settings, pdb)

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent with length of 1
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel
	<-c                                             // This blocks the main thread until an interrupt is received
	logger.Info().Msg("Gracefully shutting down and running cleanup tasks...")
	_ = ctx.Done()
	_ = pdb.DBS().Writer.Close()
	_ = pdb.DBS().Reader.Close()
}

// startMonitoringServer start server for monitoring endpoints. Could likely be moved to shared lib.
func startMonitoringServer(logger zerolog.Logger, settings *config.Settings) {
	monApp := fiber.New(fiber.Config{DisableStartupMessage: true})
	monApp.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("healthy")
	})

	monApp.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	// New endpoints
	monApp.Get("/device-config/vin/:vin", getDefaultConfigHandler)
	monApp.Get("/device-config/serial/:serial", getDefaultConfigHandler)

	go func() {
		// 8888 is our standard port for exposing metrics in DIMO infra
		if err := monApp.Listen(":" + settings.MonitoringPort); err != nil {
			logger.Fatal().Err(err).Str("port", settings.MonitoringPort).Msg("Failed to start monitoring web server.")
		}
	}()

	logger.Info().Str("port", "8888").Msg("Started monitoring web server.")
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

func startConfigAPI(logger zerolog.Logger, settings *config.Settings) *fiber.App {
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

	app.Get("/v1/swagger/*", swagger.HandlerDefault)

	app.Get("/v1/default-config/:identifier", getDefaultConfigHandler)

	return app
}

func getDefaultConfigHandler(c *fiber.Ctx) error {
	identifier := c.Params("identifier")

	defaultConfig := Response{
		PIDURL:   fmt.Sprintf("https://something/default/pid-config/%s", identifier),
		PowerURL: fmt.Sprintf("https://something/default/power-config/%s", identifier),
		DBCURL:   fmt.Sprintf("https://something/default/dbc-config/%s", identifier),
	}

	// Desired response payload format
	responsePayload := fiber.Map{
		"pidsUrl":  defaultConfig.PIDURL,
		"powerUrl": defaultConfig.PowerURL,
		"dbcUrl":   defaultConfig.DBCURL,
	}

	return c.JSON(responsePayload)
}

// Code below copied from device-data-api/main.go
func ErrorHandler(c *fiber.Ctx, err error, logger zerolog.Logger) error {
	code := fiber.StatusInternalServerError // Default 500 statuscode
	message := "Internal error."

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	}

	logger.Err(err).Int("code", code).Str("path", strings.TrimPrefix(c.Path(), "/")).Msg("Failed request.")

	return c.Status(code).JSON(CodeResp{Code: code, Message: message})
}

type CodeResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

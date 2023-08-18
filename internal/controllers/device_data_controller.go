package controllers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type DeviceDataController struct {
	Settings *config.Settings
	log      *zerolog.Logger
}

type Response struct {
	PIDURL   string `json:"pidsUrl"`
	PowerURL string `json:"powerUrl"`
	DBCURL   string `json:"dbcUrl"`
}
type CodeResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewDeviceDataController constructor
func NewDeviceDataController(settings *config.Settings, logger *zerolog.Logger) DeviceDataController {
	return DeviceDataController{
		Settings: settings,
		log:      logger,
	}
}

// GetHistoricalRaw godoc
// @Description  Get all historical data for a userDeviceID, within start and end range
// @Tags         device-data
// @Produce      json
// @Success      200
// @Param        userDeviceID  path   string  true   "user id"
// @Param        startDate     query  string  false  "startDate eg 2022-01-02. if empty two weeks back"
// @Param        endDate       query  string  false  "endDate eg 2022-03-01. if empty today"
// @Security     BearerAuth
// @Router       /user/device-data/{userDeviceID}/historical [get]
func (d *DeviceDataController) GetDefaultConfigHandler(c *fiber.Ctx) error {
	vin := c.Params("vin")

	defaultConfig := Response{
		PIDURL:   fmt.Sprintf("https://something/default/pid-config/%s", vin),
		PowerURL: fmt.Sprintf("https://something/default/power-config/%s", vin),
		DBCURL:   fmt.Sprintf("https://something/default/dbc-config/%s", vin),
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

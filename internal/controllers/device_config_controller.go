package controllers

import (
	"fmt"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type DeviceConfigController struct {
	Settings *config.Settings
	log      *zerolog.Logger
}

type PIDConfig struct {
	Name            string `json:"name"`
	Header          int    `json:"header"`
	Mode            int    `json:"mode"`
	PID             int    `json:"pid"`
	Formula         string `json:"formula"`
	IntervalSeconds int    `json:"intervalSeconds"`
}

type PowerConfig struct {
	Battery struct {
		CriticalLevel struct {
			Voltage string `json:"voltage"`
		} `json:"critical_level"`
	} `json:"battery"`

	SafetyCutOut struct {
		Voltage string `json:"voltage"`
	} `json:"safety_cut-out"`

	SleepTimer struct {
		EventDriven struct {
			Interval string `json:"interval"`
			Period   string `json:"period"`
		} `json:"event_driven"`

		InactivityAfterSleep struct {
			Interval string `json:"interval"`
		} `json:"inactivity_after_sleep"`

		InactivityFallback struct {
			Interval string `json:"interval"`
		} `json:"inactivity_fallback"`
	} `json:"sleep_timer"`

	WakeTrigger struct {
		VoltageLevel string `json:"voltage_level"`
	} `json:"wake_trigger"`
}

// NewDeviceConfigController constructor
func NewDeviceConfigController(settings *config.Settings, logger *zerolog.Logger) DeviceConfigController {
	return DeviceConfigController{
		Settings: settings,
		log:      logger,
	}
}

// / GetPIDConfig godoc
// @Description  Retrieve the PID configuration based on a given VIN
// @Tags         vehicle-signal-decoding
// @Produce      json
// @Success      200 {object} PIDConfig
// @Param        vin  path   string  true   "vehicle identification number (VIN)"
// @Router       /device-config/{vin}/pid [get]
func (d *DeviceConfigController) GetPIDConfig(c *fiber.Ctx) error {
	vin := c.Params("vin")
	pidConfig := PIDConfig{
		Name:            vin,
		Header:          2015,
		Mode:            9,
		PID:             2,
		Formula:         "ascii: 3|17 X",
		IntervalSeconds: 5,
	}
	return c.JSON(pidConfig)
}

// GetPowerConfig godoc
// @Description  Retrieve the power configuration based on a given VIN
// @Tags         vehicle-signal-decoding
// @Produce      json
// @Success      200 {object} PowerConfig
// @Param        vin  path   string  true   "vehicle identification number (VIN)"
// @Router       /device-config/{vin}/power [get]
func (d *DeviceConfigController) GetPowerConfig(c *fiber.Ctx) error {
	// Example hardcoded power config
	vin := c.Params("vin")
	d.log.Info().Msg("recieved vin" + vin)
	powerConfig := PowerConfig{
		Battery: struct {
			CriticalLevel struct {
				Voltage string `json:"voltage"`
			} `json:"critical_level"`
		}{
			CriticalLevel: struct {
				Voltage string `json:"voltage"`
			}{
				Voltage: "200V",
			},
		},
		SafetyCutOut: struct {
			Voltage string `json:"voltage"`
		}{
			Voltage: "180V",
		},
		SleepTimer: struct {
			EventDriven struct {
				Interval string `json:"interval"`
				Period   string `json:"period"`
			} `json:"event_driven"`
			InactivityAfterSleep struct {
				Interval string `json:"interval"`
			} `json:"inactivity_after_sleep"`
			InactivityFallback struct {
				Interval string `json:"interval"`
			} `json:"inactivity_fallback"`
		}{
			EventDriven: struct {
				Interval string `json:"interval"`
				Period   string `json:"period"`
			}{
				Interval: "5m",
				Period:   "10m",
			},
			InactivityAfterSleep: struct {
				Interval string `json:"interval"`
			}{
				Interval: "30m",
			},
			InactivityFallback: struct {
				Interval string `json:"interval"`
			}{
				Interval: "1h",
			},
		},
		WakeTrigger: struct {
			VoltageLevel string `json:"voltage_level"`
		}{
			VoltageLevel: "210V",
		},
	}

	return c.JSON(powerConfig)
}

// GetDBCFile godoc
// @Description  Retrieve the URL pointing to the DBC file for a given VIN
// @Tags         vehicle-signal-decoding
// @Produce      json
// @Success      200 {string} string
// @Param        vin  path   string  true   "vehicle identification number (VIN)"
// @Router       /device-config/{vin}/dbc [get]
func (d *DeviceConfigController) GetDBCFile(c *fiber.Ctx) error {
	baseURL := d.Settings.DeploymentURL
	dbcURL := fmt.Sprintf("%s/default/dbc-config/%s.dbc", baseURL, c.Params("vin"))
	return c.JSON(fiber.Map{"dbcFileUrl": dbcURL})
}

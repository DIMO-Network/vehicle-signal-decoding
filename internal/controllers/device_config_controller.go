package controllers

import (
	"database/sql"
	"fmt"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

type DeviceConfigController struct {
	Settings *config.Settings
	log      *zerolog.Logger
	db       *sql.DB
}

// NewDeviceConfigController constructor
func NewDeviceConfigController(settings *config.Settings, logger *zerolog.Logger, database *sql.DB) DeviceConfigController {
	return DeviceConfigController{
		Settings: settings,
		log:      logger,
		db:       database,
	}
}

// InitializeDatabaseConnection connect to postgres driver
func InitializeDatabaseConnection(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}

// resolveTemplateName get template name based on VIN
func resolveTemplateName(vin string, db *sql.DB) (string, error) {
	var templateName string
	query := "SELECT template_name FROM vin_to_template WHERE vin=$1"
	err := db.QueryRow(query, vin).Scan(&templateName)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("No template found for VIN: %s", vin)
		}
		return "", err
	}
	return templateName, nil
}

// Struct definitions
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

//Endpoints:

/*
//I don't think we need hard coded GetPIDConfig anymore
// / GetPIDConfig godoc
// @Description  Retrieve the PID configuration based on a given VIN
// @Tags         vehicle-signal-decoding
// @Produce      json
// @Success      200 {object} PIDConfig
// @Param        vin  path   string  true   "vehicle identification number (VIN)"
// @Router       /device-config/:vin/pid [get]
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
*/

// I think this is what we want instead
// GetPIDSByTemplate (device-config/pid/:template_name) fetches PID config from DB based on template name provided in URL
func (d *DeviceConfigController) GetPIDsByTemplate(c *fiber.Ctx) error {
	templateName := c.Params("template_name")
	var pids []PIDConfig

	// Query the database to get the PIDs based on the template name
	query := "SELECT * FROM pid_configs WHERE template_name = $1"
	rows, err := d.db.Query(query, templateName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve PID data",
		})
	}
	defer rows.Close()

	for rows.Next() {
		var pid PIDConfig
		err := rows.Scan(&pid.Name, &pid.Header, &pid.Mode, &pid.PID, &pid.Formula, &pid.IntervalSeconds)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to process PID data",
			})
		}
		pids = append(pids, pid)
	}

	return c.JSON(pids)
}

/*
// I don't think we need hard coded GetPowerConfig anymore
// GetPowerConfig godoc
// @Description  Retrieve the power configuration based on a given VIN
// @Tags         vehicle-signal-decoding
// @Produce      json
// @Success      200 {object} PowerConfig
// @Param        vin  path   string  true   "vehicle identification number (VIN)"
// @Router       /device-config/:vin/power [get]
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
*/

// GetPowerByTemplate assumes DB stores power config flat. Dependencies: power_configs table with columns matching fields of DbPowerConfig struct
func (d *DeviceConfigController) GetPowerByTemplate(c *fiber.Ctx) error {
	templateName := c.Params("template_name")
	var powerConfigs []PowerConfig

	// Query the database to get the PowerConfigs based on the template name
	query := "SELECT battery_critical_level_voltage, safety_cut_out_voltage, sleep_timer_event_driven_interval, sleep_timer_event_driven_period, sleep_timer_inactivity_after_sleep_interval, sleep_timer_inactivity_fallback_interval, wake_trigger_voltage_level FROM power_configs WHERE template_name = $1"
	rows, err := d.db.Query(query, templateName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve Power Config data",
		})
	}
	defer rows.Close()

	for rows.Next() {
		var pc PowerConfig
		err := rows.Scan(&pc.Battery.CriticalLevel.Voltage, &pc.SafetyCutOut.Voltage, &pc.SleepTimer.EventDriven.Interval, &pc.SleepTimer.EventDriven.Period, &pc.SleepTimer.InactivityAfterSleep.Interval, &pc.SleepTimer.InactivityFallback.Interval, &pc.WakeTrigger.VoltageLevel)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to process Power Config data",
			})
		}
		powerConfigs = append(powerConfigs, pc)
	}

	return c.JSON(powerConfigs)
}

// GetDBCFile godoc
// @Description  Retrieve the URL pointing to the DBC file for a given VIN
// @Tags         vehicle-signal-decoding
// @Produce      json
// @Success      200 {string} string
// @Param        vin  path   string  true   "vehicle identification number (VIN)"
// @Router       /device-config/:vin/dbc [get]
func (d *DeviceConfigController) GetDBCFile(c *fiber.Ctx) error {
	baseURL := d.Settings.DeploymentURL
	dbcURL := fmt.Sprintf("%s/default/dbc-config/%s.dbc", baseURL, c.Params("vin"))
	return c.JSON(fiber.Map{"dbcFileUrl": dbcURL})
}

// GetConfigURLs godoc
// @Description  Retrieve the URLs for PID, Power, and DBC configuration based on a given VIN
// @Tags         vehicle-signal-decoding
// @Produce      json
// @Success      200 {object} map[string]string
// @Param        vin  path   string  true   "vehicle identification number (VIN)"
// @Router       /device-config/:vin/urls [get]
func (d *DeviceConfigController) GetConfigURLs(c *fiber.Ctx) error {
	baseURL := d.Settings.DeploymentURL
	vin := c.Params("vin")

	// Resolve template name using VIN
	templateName, err := resolveTemplateName(vin, d.db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to retrieve template name for VIN: %s", vin),
		})
	}
	pidURL := fmt.Sprintf("/device-config/%s/pid", templateName)
	powerURL := fmt.Sprintf("/device-config/%s/power", templateName)
	dbcURL := fmt.Sprintf("/device-config/%s/dbc", vin) //TODO: implement with templateName instead of VIN

	return c.JSON(fiber.Map{
		"pidURL":   baseURL + pidURL,
		"powerURL": baseURL + powerURL,
		"dbcURL":   baseURL + dbcURL,
	})
}

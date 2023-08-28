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

// resolveTemplateName retrieves associated template and parent given a serial
func resolveTemplateName(serial string, db *sql.DB) (string, string, error) {
	var templateName, parentTemplateName string
	query := `
		SELECT t.template_name, t.parent_template_name
		FROM templates t
		JOIN serial_to_template_overrides sto ON t.template_name = sto.template_name
		WHERE sto.serial = $1
	`
	err := db.QueryRow(query, serial).Scan(&templateName, &parentTemplateName)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", "", fmt.Errorf("No template found for serial: %s", serial)
		}
		return "", "", err
	}
	return templateName, parentTemplateName, nil
}

func getConfigurationVersion(configType string, templateName string, db *sql.DB) (string, error) {
	query := fmt.Sprintf("SELECT version FROM %s_configs WHERE template_name = $1", configType)
	var version string
	err := db.QueryRow(query, templateName).Scan(&version)
	if err != nil {
		return "", err
	}
	return version, nil
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

type DBCFile struct {
	FilePath     string
	TemplateName string
	Version      string
	CreatedAt    string
	UpdatedAt    string
}

//Endpoints:

// GetPIDSByTemplate godoc
// @Description  Retrieves a list of PID configurations from the database given a template name
// @Tags         vehicle-signal-decoding
// @Produce      json
// @Success      200 {array} PIDConfig "Successfully retrieved PID Configurations"
// @Param        template_name  path   string  true   "template name"
// @Router       /device-config/:template_name/pids [get]
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

// GetPowerByTemplate godoc
// @Description  Fetches the power configurations from power_configs table given a template name
// @Tags         vehicle-signal-decoding
// @Produce      json
// @Success      200 {array} PowerConfig "Successfully retrieved Power Configurations"
// @Param        template_name  path   string  true   "template name"
// @Router       /device-config/:template_name/power [get]
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

// GetDBCFilePathByTemplateName godoc
// @Description  Fetches the DBC file path from the dbc_files table given a template name
// @Tags         vehicle-signal-decoding
// @Produce      plain
// @Success      200 {string} string "Successfully retrieved DBC file path"
// @Param        template_name  path   string  true   "template name"
// @Router       /device-config/:template_name/dbc-file-path [get]
func (d *DeviceConfigController) GetDBCFilePathByTemplateName(c *fiber.Ctx) error { // getDBCFilePathByTemplateName fetches dbc_file_path from dbc_file table given template_name

	templateName := c.Params("template_name")

	query := "SELECT dbc_file_path FROM dbc_files WHERE template_name = $1"
	var filePath string
	err := d.db.QueryRow(query, templateName).Scan(&filePath)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).SendString(fmt.Sprintf("No DBC file found for template name: %s", templateName))
		}
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to retrieve DBC file path for template: %s, Error: %s", templateName, err.Error()))
	}

	// returns DBC file path as a plain text response
	return c.Status(fiber.StatusOK).SendString(filePath)
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
	templateName, parentTemplateName, err := resolveTemplateName(vin, d.db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to retrieve template name for VIN: %s", vin),
		})
	}

	pidTemplateName := templateName
	powerTemplateName := parentTemplateName

	if powerTemplateName == "" {
		powerTemplateName = templateName
	}

	//Versioning

	pidVersion, err := getConfigurationVersion("pid", pidTemplateName, d.db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to retrieve PID version for template: %s", pidTemplateName),
		})
	}

	powerVersion, err := getConfigurationVersion("power", powerTemplateName, d.db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to retrieve Power version for template: %s", powerTemplateName),
		})
	}

	pidURL := fmt.Sprintf("%s/device-config/pid/%s", baseURL, pidTemplateName)
	powerURL := fmt.Sprintf("%s/device-config/power/%s", baseURL, powerTemplateName)
	dbcURL := fmt.Sprintf("%s/device-config/dbc/%s", baseURL, pidTemplateName)

	return c.JSON(fiber.Map{
		"pidUrl":       pidURL,
		"powerUrl":     powerURL,
		"dbcURL":       dbcURL, //TODO: implement
		"pidVersion":   pidVersion,
		"powerVersion": powerVersion,
	})
}

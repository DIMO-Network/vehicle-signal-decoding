package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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
	ID              int64     `json:"id"`
	TemplateName    string    `json:"template_name,omitempty"`
	Header          int       `json:"header"`
	Mode            int       `json:"mode"`
	Pid             int       `json:"pid"`
	Formula         string    `json:"formula"`
	IntervalSeconds int       `json:"interval_seconds"`
	Version         string    `json:"version,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type PowerConfig struct {
	ID                                     int64     `json:"id"`
	Version                                string    `json:"version,omitempty"`
	TemplateName                           string    `json:"template_name"`
	BatteryCriticalLevelVoltage            string    `json:"battery_critical_level_voltage"`
	SafetyCutOutVoltage                    string    `json:"safety_cut_out_voltage"`
	SleepTimerEventDrivenInterval          string    `json:"sleep_timer_event_driven_interval"`
	SleepTimerEventDrivenPeriod            string    `json:"sleep_timer_event_driven_period"`
	SleepTimerInactivityAfterSleepInterval string    `json:"sleep_timer_inactivity_after_sleep_interval"`
	SleepTimerInactivityFallbackInterval   string    `json:"sleep_timer_inactivity_fallback_interval"`
	WakeTriggerVoltageLevel                string    `json:"wake_trigger_voltage_level"`
	CreatedAt                              time.Time `json:"created_at"`
	UpdatedAt                              time.Time `json:"updated_at"`
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

	/// Query the database to get the PIDs based on the template name using SQLBoiler
	pidConfigs, err := models.PidConfigs(
		models.PidConfigWhere.TemplateName.EQ(null.StringFrom(templateName)),
	).All(context.Background(), d.db)

	// Error handling
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fiber.NewError(fiber.StatusNotFound, "No PID data found for the given template name.")
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve PID data",
		})
	}
	// Convert SQLBoiler model
	for _, pidConfig := range pidConfigs {
		pid := PIDConfig{
			ID:              pidConfig.ID,
			TemplateName:    pidConfig.TemplateName.String,
			Header:          pidConfig.Header,
			Mode:            pidConfig.Mode,
			Pid:             pidConfig.Pid,
			Formula:         pidConfig.Formula,
			IntervalSeconds: pidConfig.IntervalSeconds,
			Version:         pidConfig.Version.String,
			CreatedAt:       pidConfig.CreatedAt,
			UpdatedAt:       pidConfig.UpdatedAt,
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

	// Query the database to get the PowerConfigs based on the template name using SQLBoiler
	dbPowerConfigs, err := models.PowerConfigs(
		models.PowerConfigWhere.TemplateName.EQ(templateName),
	).All(context.Background(), d.db)

	// Error handling
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fiber.NewError(fiber.StatusNotFound, "No Power Config data found for the given template name.")
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve Power Config data",
		})
	}

	var apiPowerConfigs []PowerConfig

	for _, dbPowerConfig := range dbPowerConfigs {
		apiPowerConfig := PowerConfig{
			ID:                                     dbPowerConfig.ID,
			Version:                                dbPowerConfig.Version.String,
			TemplateName:                           dbPowerConfig.TemplateName,
			BatteryCriticalLevelVoltage:            dbPowerConfig.BatteryCriticalLevelVoltage,
			SafetyCutOutVoltage:                    dbPowerConfig.SafetyCutOutVoltage,
			SleepTimerEventDrivenInterval:          dbPowerConfig.SleepTimerEventDrivenInterval,
			SleepTimerEventDrivenPeriod:            dbPowerConfig.SleepTimerEventDrivenPeriod,
			SleepTimerInactivityAfterSleepInterval: dbPowerConfig.SleepTimerInactivityAfterSleepInterval,
			SleepTimerInactivityFallbackInterval:   dbPowerConfig.SleepTimerInactivityFallbackInterval,
			WakeTriggerVoltageLevel:                dbPowerConfig.WakeTriggerVoltageLevel,
			CreatedAt:                              dbPowerConfig.CreatedAt,
			UpdatedAt:                              dbPowerConfig.UpdatedAt,
		}

		apiPowerConfigs = append(apiPowerConfigs, apiPowerConfig)
	}

	return c.JSON(apiPowerConfigs)

}

// GetDBCFilePathByTemplateName godoc
// @Description  Fetches the DBC file path from the dbc_files table given a template name
// @Tags         vehicle-signal-decoding
// @Produce      plain
// @Success      200 {string} string "Successfully retrieved DBC file path"
// @Param        template_name  path   string  true   "template name"
// @Router       /device-config/:template_name/dbc-file-path [get]
func (d *DeviceConfigController) GetDBCFilePathByTemplateName(c *fiber.Ctx) error {
	templateName := c.Params("template_name")

	// Query the database using SQLBoiler
	dbResult, err := models.DBCFiles(qm.Where("template_name=?", templateName)).One(context.Background(), d.db)

	// Error handling
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).SendString(fmt.Sprintf("No DBC file found for template name: %s", templateName))
		}
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to retrieve DBC file path for template: %s, Error: %s", templateName, err.Error()))
	}

	// Return the DBC file path as plain text
	return c.Status(fiber.StatusOK).SendString(dbResult.DBCFilePath)
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

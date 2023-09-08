package controllers

import (
	"context"
	"database/sql"
	"encoding/binary"
	"fmt"
	"log"
	"time"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq" //nolint
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
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

type DeviceConfigResponse struct {
	PidURL           string `json:"pidUrl"`
	DeviceSettingURL string `json:"deviceSettingUrl"`
	DbcURL           string `json:"dbcURL"`
	Version          string `json:"version"`
}

// resolveTemplateName retrieves associated template and parent given a serial
func resolveTemplateName(serial string, db *sql.DB) (string, string, error) {
	// Create empty Template and SerialToTemplateOverride structs to hold query results
	var template models.Template

	// Create query modifiers
	queryMods := []qm.QueryMod{
		qm.Select("t.template_name", "t.parent_template_name"),
		qm.From("templates t"),
		qm.InnerJoin("serial_to_template_overrides sto ON t.template_name = sto.template_name"),
		qm.Where("sto.serial = ?", serial),
	}

	// Execute the query and bind the results to the Template struct
	err := models.Templates(queryMods...).Bind(context.Background(), db, &template)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", "", fmt.Errorf("No template found for serial: %s", serial)
		}
		return "", "", err
	}

	parentTemplateName := ""
	if template.ParentTemplateName.Valid {
		parentTemplateName = template.ParentTemplateName.String
	}

	return template.TemplateName, parentTemplateName, nil
}

func getVersionByTemplateName(templateName string, db boil.ContextExecutor) (string, error) {
	template, err := models.Templates(models.TemplateWhere.TemplateName.EQ(templateName)).One(context.Background(), db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("No version found for template name: %s", templateName)
		}
		return "", err
	}
	return template.Version, nil
}

type PIDConfig struct {
	ID              int64  `json:"id"`
	TemplateName    string `json:"template_name,omitempty"`
	Header          []byte `json:"header"`
	Mode            []byte `json:"mode"`
	Pid             []byte `json:"pid"`
	Formula         string `json:"formula"`
	IntervalSeconds int    `json:"interval_seconds"`
	Protocol        string `json:"protocol,omitempty"`
}

type DeviceSetting struct {
	ID                                     int64     `json:"id"`
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

// ProtobufToJSON converts a Protobuf message to its JSON representation.
func ProtobufToJSON(message proto.Message) (string, error) {
	marshaler := protojson.MarshalOptions{
		UseProtoNames: true,
	}
	bytes, err := marshaler.Marshal(message)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
func bytesToUint32(b []byte) (uint32, error) {
	if len(b) != 4 {
		return 0, errors.New("invalid length for uint32 conversion")
	}
	return binary.LittleEndian.Uint32(b), nil
}

// GetPIDsByTemplate godoc
// @Description  Retrieves a list of PID configurations from the database given a template name
// @Tags         vehicle-signal-decoding
// @Produce      json
// @Success      200 {array} PIDConfig "Successfully retrieved PID Configurations"
// @Failure 404 "No PID Config data found for the given template name."
// @Param        template_name  path   string  true   "template name"
// @Router       /device-config/:template_name/pids [get]
func (d *DeviceConfigController) GetPIDsByTemplate(c *fiber.Ctx) error {
	templateName := c.Params("template_name")

	/// Query the database to get the PIDs based on the template name using SQLBoiler
	pidConfigs, err := models.PidConfigs(
		models.PidConfigWhere.TemplateName.EQ(templateName),
	).All(c.Context(), d.db)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fiber.NewError(fiber.StatusNotFound, "No PID data found for the given template name.")
		}
		return errors.Wrap(err, "Failed to retrieve PID Configs")
	}

	acceptHeader := c.Get("Accept", "application/json")

	if acceptHeader == "application/x-protobuf" {
		protoPIDs := &grpc.PIDRequests{
			TemplateName: templateName,
			Version:      "",
		}

		for _, pidConfig := range pidConfigs {
			headerUint32, err := bytesToUint32(pidConfig.Header)
			if err != nil {
				continue
			}

			modeUint32, err := bytesToUint32(pidConfig.Mode)
			if err != nil {
				continue
			}

			pidUint32, err := bytesToUint32(pidConfig.Pid)
			if err != nil {
				continue
			}
			pid := &grpc.PIDConfig{
				Name:            pidConfig.TemplateName,
				Header:          headerUint32,
				Mode:            modeUint32,
				Pid:             pidUint32,
				Formula:         pidConfig.Formula,
				IntervalSeconds: int32(pidConfig.IntervalSeconds),
				Protocol:        pidConfig.Protocol,
			}
			protoPIDs.Requests = append(protoPIDs.Requests, pid)
		}

		out, err := proto.Marshal(protoPIDs)
		if err != nil {
			return errors.Wrap(err, "Failed to serialize to protobuf")
		}

		// Debugging
		jsonStr, err := ProtobufToJSON(protoPIDs)
		if err != nil {
			log.Printf("Failed to convert Protobuf to JSON: %v", err)
		} else {
			log.Printf("Protobuf as JSON: %s", jsonStr)
		}

		c.Set("Content-Type", "application/x-protobuf")

		return c.Send(out)
	}
	pids := make([]PIDConfig, len(pidConfigs))
	for i, pidConfig := range pidConfigs {
		pid := PIDConfig{
			ID:              pidConfig.ID,
			Header:          pidConfig.Header,
			Mode:            pidConfig.Mode,
			Pid:             pidConfig.Pid,
			Formula:         pidConfig.Formula,
			IntervalSeconds: pidConfig.IntervalSeconds,
			Protocol:        pidConfig.Protocol,
		}
		pids[i] = pid
	}
	return c.JSON(pids)

}

// GetDeviceSettingsByTemplate godoc
// @Description  Fetches the device settings configurations from device_settings table given a template name
// @Tags         vehicle-signal-decoding
// @Produce      json
// @Success      200 {object} DeviceSetting "Successfully retrieved Device Settings"
// @Failure 404 "No Device Settings data found for the given template name."
// @Param        template_name  path   string  true   "template name"
// @Router       /device-config/:template_name/deviceSettings [get]
func (d *DeviceConfigController) GetDeviceSettingsByTemplate(c *fiber.Ctx) error {
	templateName := c.Params("template_name")

	// Query the database to get the Device Settings based on the template name using SQLBoiler
	dbDeviceSettings, err := models.DeviceSettings(
		models.DeviceSettingWhere.TemplateName.EQ(templateName),
	).One(c.Context(), d.db)

	// Error handling
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fiber.NewError(fiber.StatusNotFound, "No Device Settings data found for the given template name.")
		}
		return errors.Wrap(err, "Failed to retrieve Device Settings")
	}

	apiDeviceSettings := DeviceSetting{

		ID:                                     dbDeviceSettings.ID,
		BatteryCriticalLevelVoltage:            dbDeviceSettings.BatteryCriticalLevelVoltage,
		SafetyCutOutVoltage:                    dbDeviceSettings.SafetyCutOutVoltage,
		SleepTimerEventDrivenInterval:          dbDeviceSettings.SleepTimerEventDrivenInterval,
		SleepTimerEventDrivenPeriod:            dbDeviceSettings.SleepTimerEventDrivenPeriod,
		SleepTimerInactivityAfterSleepInterval: dbDeviceSettings.SleepTimerInactivityAfterSleepInterval,
		SleepTimerInactivityFallbackInterval:   dbDeviceSettings.SleepTimerInactivityFallbackInterval,
		WakeTriggerVoltageLevel:                dbDeviceSettings.WakeTriggerVoltageLevel,
	}

	return c.JSON(apiDeviceSettings)

}

// GetDBCFileByTemplateName godoc
// @Description  Fetches the DBC file from the dbc_files table given a template name
// @Tags         vehicle-signal-decoding
// @Produce      plain
// @Success      200 {string} string "Successfully retrieved DBC file"
// @Failure 404 "No DBC file found for the given template name."
// @Param        template_name  path   string  true   "template name"
// @Router       /device-config/:template_name/dbc-file [get]
func (d *DeviceConfigController) GetDBCFileByTemplateName(c *fiber.Ctx) error {
	templateName := c.Params("template_name")

	// Query the database using SQLBoiler
	//use same logic as above
	dbResult, err := models.DBCFiles(
		models.DBCFileWhere.TemplateName.EQ(templateName)).One(c.Context(), d.db)

	// Error handling
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).SendString(fmt.Sprintf("No DBC file found for template name: %s", templateName))
		}
		return errors.Wrap(err, "Failed to retrieve DBC File")
	}

	// Return the DBC file itself
	if c.Accepts("text/plain") == "text/plain" {
		c.Status(fiber.StatusOK).Set("Content-Type", "text/plain")
		return c.SendString(dbResult.DBCFile)
	}
	return c.Status(fiber.StatusNotAcceptable).SendString("Not Acceptable")

}

// GetConfigURLs godoc
// @Description  Retrieve the URLs for PID, DeviceSettings, and DBC configuration based on a given VIN
// @Tags         vehicle-signal-decoding
// @Produce      json
// @Success      200 {object} DeviceConfigResponse "Successfully retrieved configuration URLs"
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

	//Versioning
	version, err := getVersionByTemplateName(templateName, d.db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to retrieve version for template: %s", templateName),
		})
	}

	response := DeviceConfigResponse{
		PidURL:           fmt.Sprintf("%s/device-config/pid/%s", baseURL, templateName),
		DeviceSettingURL: fmt.Sprintf("%s/device-config/deviceSetting/%s", baseURL, parentTemplateName),
		DbcURL:           fmt.Sprintf("%s/device-config/dbc/%s", baseURL, templateName),
		Version:          version,
	}

	return c.JSON(response)
}

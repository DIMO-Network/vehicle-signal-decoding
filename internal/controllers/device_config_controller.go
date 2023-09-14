package controllers

import (
	"context"
	"database/sql"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/services"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq" //nolint
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type DeviceConfigController struct {
	settings      *config.Settings
	log           *zerolog.Logger
	db            *sql.DB
	userDeviceSvc services.UserDeviceService
}

// NewDeviceConfigController constructor
func NewDeviceConfigController(settings *config.Settings, logger *zerolog.Logger, database *sql.DB, userDeviceSvc services.UserDeviceService) DeviceConfigController {
	return DeviceConfigController{
		settings:      settings,
		log:           logger,
		db:            database,
		userDeviceSvc: userDeviceSvc,
	}

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
type DeviceConfigResponse struct {
	PidURL           string `json:"pidUrl"`
	DeviceSettingURL string `json:"deviceSettingUrl"`
	DbcURL           string `json:"dbcURL"`
	Version          string `json:"version"`
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
// @Success      200 {object} grpc.PIDRequests "Successfully retrieved PID Configurations"
// @Failure 404 "No PID Config data found for the given template name."
// @Param        templateName  path   string  true   "template name"
// @Router       /device-config/{templateName}/pids [get]
func (d *DeviceConfigController) GetPIDsByTemplate(c *fiber.Ctx) error {
	templateName := c.Params("templateName")

	template, err := models.FindTemplate(c.Context(), d.db, templateName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("No template with name: %s found", templateName))
		}
		return errors.Wrapf(err, "Failed to retrieve Template %s", templateName)
	}

	pidConfigs, err := models.PidConfigs(
		models.PidConfigWhere.TemplateName.EQ(templateName),
	).All(c.Context(), d.db)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fiber.NewError(fiber.StatusNotFound, "No PID data found for the given template name.")
		}
		return errors.Wrap(err, "Failed to retrieve PID Configs")
	}

	protoPIDs := &grpc.PIDRequests{
		TemplateName: templateName,
	}
	if template != nil {
		protoPIDs.Version = template.Version
	}

	for _, pidConfig := range pidConfigs {
		headerUint32, err := bytesToUint32(pidConfig.Header)
		if err != nil {
			d.log.Err(err).Send()
		}

		modeUint32, err := bytesToUint32(pidConfig.Mode)
		if err != nil {
			d.log.Err(err).Send()
		}

		pidUint32, err := bytesToUint32(pidConfig.Pid)
		if err != nil {
			d.log.Err(err).Send()
		}
		pid := &grpc.PIDConfig{
			Name:            pidConfig.SignalName,
			Header:          headerUint32,
			Mode:            modeUint32,
			Pid:             pidUint32,
			Formula:         pidConfig.Formula,
			IntervalSeconds: int32(pidConfig.IntervalSeconds),
			Protocol:        pidConfig.Protocol,
		}
		protoPIDs.Requests = append(protoPIDs.Requests, pid)
	}

	acceptHeader := c.Get("Accept", "application/json")
	if acceptHeader == "application/x-protobuf" {
		out, err := proto.Marshal(protoPIDs)
		if err != nil {
			return errors.Wrap(err, "Failed to serialize to protobuf")
		}

		c.Set("Content-Type", "application/x-protobuf")

		return c.Send(out)
	}

	return c.JSON(protoPIDs)

}

// GetDeviceSettingsByTemplate godoc
// @Description  Fetches the device settings configurations from device_settings table given a template name
// @Tags         vehicle-signal-decoding
// @Produce      json
// @Success      200 {object} DeviceSetting "Successfully retrieved Device Settings"
// @Failure 404 "No Device Settings data found for the given template name."
// @Param        templateName  path   string  true   "template name"
// @Router       /device-config/{templateName}/deviceSettings [get]
func (d *DeviceConfigController) GetDeviceSettingsByTemplate(c *fiber.Ctx) error {
	templateName := c.Params("templateName")

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
// @Param        templateName  path   string  true   "template name"
// @Router       /device-config/{templateName}/dbc [get]
func (d *DeviceConfigController) GetDBCFileByTemplateName(c *fiber.Ctx) error {
	templateName := c.Params("templateName")

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
// @Failure 404  "Not Found - No templates available for the given parameters"
// @Param        vin  path   string  true   "vehicle identification number (VIN)"
// @Router       /device-config/{vin}/urls [get]
func (d *DeviceConfigController) GetConfigURLs(c *fiber.Ctx) error {
	baseURL := d.settings.DeploymentURL
	vin := c.Params("vin")

	// Set up gRPC call to retrieve UserDevice by VIN
	ud, err := d.userDeviceSvc.GetUserDeviceByVIN(c.Context(), vin)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to retrieve user device for VIN: %s", vin),
		})
	}

	switch ud.CANProtocol {
	case "6":
		ud.CANProtocol = models.CanProtocolTypeCAN11_500
	case "7":
		ud.CANProtocol = models.CanProtocolTypeCAN29_500
	case "":
		ud.CANProtocol = models.CanProtocolTypeCAN11_500
	}

	// Set default for PowerTrainType if empty
	if ud.PowerTrainType == "" {
		ud.PowerTrainType = "ICE"
	}

	// Query templates, filter by protocol and powertrain
	templates, err := models.Templates(
		models.TemplateWhere.Protocol.EQ(ud.CANProtocol),
		models.TemplateWhere.Powertrain.EQ(ud.PowerTrainType),
	).All(context.Background(), d.db)

	if err != nil {
		// Check if error is due to no matching rows
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": fmt.Sprintf("No templates found for protocol: %s and powertrain: %s", ud.CANProtocol, ud.PowerTrainType),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to query templates for protocol: %s and powertrain: %s", ud.CANProtocol, ud.PowerTrainType),
		})
	}

	//Return first match

	templateName := templates[0].TemplateName

	var parentTemplateName string
	if templates[0].ParentTemplateName.Valid {
		parentTemplateName = templates[0].ParentTemplateName.String
	} else {
		parentTemplateName = templateName
	}
	version := templates[0].Version

	response := DeviceConfigResponse{
		PidURL:           fmt.Sprintf("%s/device-config/%s/pids", baseURL, templateName),
		DeviceSettingURL: fmt.Sprintf("%s/device-config/%s/deviceSettings", baseURL, parentTemplateName),
		DbcURL:           fmt.Sprintf("%s/device-config/%s/dbc", baseURL, templateName),
		Version:          version,
	}

	return c.JSON(response)
}

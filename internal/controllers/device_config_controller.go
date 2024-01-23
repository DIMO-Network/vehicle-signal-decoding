package controllers

import (
	"bytes"
	"database/sql"
	"encoding/binary"
	"encoding/json"
	"fmt"
	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"strings"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/common"

	"github.com/volatiletech/sqlboiler/v4/types"

	pb "github.com/DIMO-Network/devices-api/pkg/grpc"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/services"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq" //nolint
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/proto"
)

type DeviceConfigController struct {
	settings              *config.Settings
	log                   *zerolog.Logger
	db                    *sql.DB
	userDeviceSvc         services.UserDeviceService
	deviceDefSvc          services.DeviceDefinitionsService
	deviceTemplateService services.DeviceTemplateService
}

// NewDeviceConfigController constructor
func NewDeviceConfigController(settings *config.Settings, logger *zerolog.Logger, database *sql.DB, userDeviceSvc services.UserDeviceService, deviceDefSvc services.DeviceDefinitionsService, deviceTemplateService services.DeviceTemplateService) DeviceConfigController {
	return DeviceConfigController{
		settings:              settings,
		log:                   logger,
		db:                    database,
		userDeviceSvc:         userDeviceSvc,
		deviceDefSvc:          deviceDefSvc,
		deviceTemplateService: deviceTemplateService,
	}

}

type DeviceTemplateStatusResponse struct {
	// IsTemplateUpdated based on information we have, is the device's template up to date?
	IsTemplateUpdated bool   `json:"is_template_updated"`
	TemplateVersion   string `json:"template_version"`
}

type SettingsData struct {
	SafetyCutOutVoltage             float64 `json:"safety_cut_out_voltage"`
	SleepTimerEventDrivenPeriodSecs float64 `json:"sleep_timer_event_driven_period_secs"`
	WakeTriggerVoltageLevel         float64 `json:"wake_trigger_voltage_level"`
}

func bytesToUint32(b []byte) (uint32, error) {
	u := binary.BigEndian.Uint32(padByteArray(b, 4))
	return u, nil
}

// GetPIDsByTemplate godoc
// @Description  Retrieves a list of PID configurations from the database given a template name
// @Tags         vehicle-signal-decoding
// @Produce      json
// @Produce      application/x-protobuf
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

	// Check if template has a parent and retrieve its PID configs
	if template.ParentTemplateName.Valid {
		pidConfigsParent, err := models.PidConfigs(
			models.PidConfigWhere.TemplateName.EQ(template.ParentTemplateName.String),
		).All(c.Context(), d.db)

		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return errors.Wrap(err, "Failed to retrieve Parent PID Configs")
		}

		// Append the parent pidConfigs to the original pidConfigs
		pidConfigs = append(pidConfigs, pidConfigsParent...)
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
			return fiber.NewError(fiber.StatusInternalServerError, "invalid header bytes configuration: "+err.Error())
		}

		modeUint32, err := bytesToUint32(pidConfig.Mode)
		if err != nil {
			d.log.Err(err).Send()
			return fiber.NewError(fiber.StatusInternalServerError, "invalid mode bytes configuration: "+err.Error())
		}

		pidUint32, err := bytesToUint32(pidConfig.Pid)
		if err != nil {
			d.log.Err(err).Send()
			return fiber.NewError(fiber.StatusInternalServerError, "invalid pid bytes configuration: "+err.Error())
		}

		pid := &grpc.PIDConfig{
			Name:            pidConfig.SignalName,
			Header:          headerUint32,
			Mode:            modeUint32,
			Pid:             pidUint32,
			Formula:         common.PrependFormulaTypeDefault(pidConfig.Formula),
			IntervalSeconds: uint32(pidConfig.IntervalSeconds),
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

// GetDeviceSettingsByName godoc
// @Description  Fetches the device settings configurations from device_settings table given a name. Note that device settings mostly only vary by powertrain and
// @Description  may or may not be attached to a specific template. To return protobuf: "application/x-protobuf"
// @Description  Note that the templateName returned here is actually the device setting name
// @Tags         vehicle-signal-decoding
// @Produce      json
// @Produce      application/x-protobuf
// @Success      200 {object} grpc.DeviceSetting "Successfully retrieved Device Settings"
// @Failure 404 "No Device Settings data found for the given name."
// @Param        name  path   string  true   "name"
// @Router       /device-config/settings/{name} [get]
func (d *DeviceConfigController) GetDeviceSettingsByName(c *fiber.Ctx) error {
	name := c.Params("name")
	if len(name) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "name for settings empty")
	}

	dbDeviceSettings, err := models.FindDeviceSetting(c.Context(), d.db, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fiber.NewError(fiber.StatusNotFound, "No Device Settings data found with name: "+name)
		}
		return errors.Wrap(err, "Failed to retrieve Device Settings")
	}

	// Deserialize the settings JSONB into the SettingsData struct
	var settings SettingsData
	if dbDeviceSettings.Settings.Valid {
		jsonBytes, err := dbDeviceSettings.Settings.MarshalJSON()
		if err != nil {
			return errors.Wrap(err, "Failed to marshal settings JSON")
		}
		if err = json.Unmarshal(jsonBytes, &settings); err != nil {
			return errors.Wrap(err, "Failed to deserialize settings data")
		}
	} else {
		return fiber.NewError(fiber.StatusNotFound, "Settings data is null")
	}

	protoDeviceSettings := &grpc.DeviceSetting{
		TemplateName:                             dbDeviceSettings.Name, // in future add a Name field, once safe to change proto
		SafetyCutOutVoltage:                      float32(settings.SafetyCutOutVoltage),
		SleepTimerEventDrivenPeriodSecs:          float32(settings.SleepTimerEventDrivenPeriodSecs),
		WakeTriggerVoltageLevel:                  float32(settings.WakeTriggerVoltageLevel),
		SleepTimerEventDrivenIntervalSecs:        float32(3600), // not used by Macaron
		SleepTimerInactivityAfterSleepSecs:       float32(21600),
		SleepTimerInactivityFallbackIntervalSecs: float32(21600),
		//TemplateName: dbDeviceSettings.TemplateName.String, // in future we could do this, could be empty
		//Version: "v1.0.1", // for future - once safe to change proto file
	}

	acceptHeader := c.Get("Accept", "application/json")
	if acceptHeader == "application/x-protobuf" {
		out, err := proto.Marshal(protoDeviceSettings)
		if err != nil {
			return errors.Wrap(err, "Failed to serialize to protobuf")
		}
		c.Set("Content-Type", "application/x-protobuf")
		return c.Send(out)
	}

	return c.JSON(protoDeviceSettings)
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

// GetConfigURLsFromVIN godoc
// @Description  Retrieve the URLs for PID, DeviceSettings, and DBC configuration based on a given VIN. These could be empty if not configs available
// @Tags         vehicle-signal-decoding
// @Produce      json
// @Success      200 {object} services.DeviceConfigResponse "Successfully retrieved configuration URLs"
// @Failure 404  "Not Found - No templates available for the given parameters"
// @Param        vin  path   string  true   "vehicle identification number (VIN)"
// @Param        protocol  query   string  false  "CAN Protocol, '6' or '7'"
// @Router       /device-config/vin/{vin}/urls [get]
func (d *DeviceConfigController) GetConfigURLsFromVIN(c *fiber.Ctx) error {
	vin := c.Params("vin")
	protocol := c.Query("protocol", "")

	ud, err := d.userDeviceSvc.GetUserDeviceByVIN(c.Context(), vin)
	if err != nil {
		definitionResp, err := d.deviceDefSvc.DecodeVIN(c.Context(), vin)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("could not decode VIN, contact support if you're sure this is valid VIN: %s", vin)})
		}

		ud = &pb.UserDevice{
			DeviceDefinitionId: definitionResp.DeviceDefinitionId,
			PowerTrainType:     definitionResp.Powertrain,
			CANProtocol:        protocol,
			Vin:                &vin,
		}
		if len(definitionResp.DeviceStyleId) > 0 {
			ud.DeviceStyleId = &definitionResp.DeviceStyleId
		}
	}

	response, err := d.deviceTemplateService.ResolveDeviceConfiguration(c, ud)
	if err != nil {
		return err
	}
	// store the template used
	_, err = d.deviceTemplateService.StoreLastTemplateRequested(c.Context(), *ud.Vin, response.DbcURL, response.PidURL, response.DeviceSettingURL, response.Version)
	if err != nil {
		d.log.Err(err).Str("vin", *ud.Vin).Msg("failed to store template urls used for eth addr")
	}

	return c.JSON(response)
}

// GetConfigURLsFromEthAddr godoc
// @Description  Retrieve the URLs for PID, DeviceSettings, and DBC configuration based on device's Ethereum Address. These could be empty if not configs available
// @Tags         vehicle-signal-decoding
// @Produce      json
// @Success      200 {object} services.DeviceConfigResponse "Successfully retrieved configuration URLs"
// @Failure 404  "Not Found - No templates available for the given parameters"
// @Failure 400  "incorrect eth addr format"
// @Param        ethAddr  path   string  true  "Ethereum Address"
// @Param        protocol  query   string  false  "CAN Protocol, '6' or '7'"
// @Router       /device-config/eth-addr/{ethAddr}/urls [get]
func (d *DeviceConfigController) GetConfigURLsFromEthAddr(c *fiber.Ctx) error {
	ethAddr := c.Params("ethAddr")
	protocol := c.Query("protocol", "")

	ud, err := d.userDeviceSvc.GetUserDeviceByEthAddr(c.Context(), ethAddr)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": fmt.Sprintf("no connected user device found for EthAddr: %s", ethAddr)})
	}

	if protocol != "" {
		ud.CANProtocol = protocol
	}

	response, err := d.deviceTemplateService.ResolveDeviceConfiguration(c, ud)
	if err != nil {
		return err
	}
	// store the template used
	_, err = d.deviceTemplateService.StoreLastTemplateRequested(c.Context(), *ud.Vin, response.DbcURL, response.PidURL, response.DeviceSettingURL, response.Version)
	if err != nil {
		d.log.Err(err).Str("vin", *ud.Vin).Msg("failed to store template urls used for eth addr")
	}

	return c.JSON(response)
}

// GetConfigStatusFromEthAddr godoc
// @Description  Helps client determine if template (pids, dbc, settings) are up to date or not for the device with the given eth addr.
// @Tags         vehicle-signal-decoding
// @Produce      json
// @Success      200 {object} DeviceTemplateStatusResponse "Successfully retrieved configuration URLs"
// @Failure 404  "Not Found - we haven't seen this device yet, assume template not up to date"
// @Failure 400  "incorrect eth addr format"
// @Param        ethAddr  path   string  true  "Ethereum Address"
// @Router       /device-config/eth-addr/{ethAddr}/status [get]
func (d *DeviceConfigController) GetConfigStatusFromEthAddr(c *fiber.Ctx) error {
	ethAddr := c.Params("ethAddr")
	// we need to get the user-device so we can find the VIN that was paired to the eth addr.
	// an improvement here would be for the mobile app to confirm templates were installed to a specific eth addr
	ud, err := d.userDeviceSvc.GetUserDeviceByEthAddr(c.Context(), ethAddr)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": fmt.Sprintf("no connected user device found for EthAddr: %s", ethAddr)})
	}

	deviceConfiguration, err := d.deviceTemplateService.ResolveDeviceConfiguration(c, ud)
	if err != nil {
		return err
	}

	dts, err := models.DeviceTemplateStatuses(models.DeviceTemplateStatusWhere.Vin.EQ(*ud.Vin)).One(c.Context(), d.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fiber.NewError(fiber.StatusNotFound, "haven't seen device with eth addr yet: "+ethAddr)
		}
		return err
	}
	// set the eth addr for record keeping, may come in handy later
	if !bytes.Equal(dts.DeviceEthAddr.Bytes, common2.HexToAddress(ethAddr).Bytes()) {
		dts.DeviceEthAddr = null.BytesFrom(common2.HexToAddress(ethAddr).Bytes())
		_, _ = dts.Update(c.Context(), d.db, boil.Infer())
	}

	isTemplateUpdated := false
	// if all this is true then we know we're up to date
	if dts.TemplateVersion == deviceConfiguration.Version &&
		dts.TemplateDBCURL == deviceConfiguration.DbcURL &&
		dts.TemplatePidURL == deviceConfiguration.PidURL &&
		dts.TemplateSettingURL == deviceConfiguration.DeviceSettingURL {

		isTemplateUpdated = true
	}

	return c.JSON(DeviceTemplateStatusResponse{
		IsTemplateUpdated: isTemplateUpdated,
		TemplateVersion:   deviceConfiguration.Version,
	})
}

func padByteArray(input []byte, targetLength int) []byte {
	if len(input) >= targetLength {
		return input // No need to pad if the input is already longer or equal to the target length
	}

	padded := make([]byte, targetLength-len(input))
	return append(padded, input...)
}

// modelMatch simply returns if the modelName is in the model List
func modelMatch(modelList types.StringArray, modelName string) bool {
	for _, m := range modelList {
		if strings.EqualFold(m, modelName) {
			return true
		}
	}
	return false
}

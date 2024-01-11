package controllers

import (
	"context"
	"database/sql"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/common"

	"github.com/volatiletech/sqlboiler/v4/types"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	p_grpc "github.com/DIMO-Network/device-definitions-api/pkg/grpc"
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
	settings                  *config.Settings
	log                       *zerolog.Logger
	db                        *sql.DB
	userDeviceSvc             services.UserDeviceService
	deviceDefSvc              services.DeviceDefinitionsService
	userDeviceTemplateService services.UserDeviceTemplateService
}

// NewDeviceConfigController constructor
func NewDeviceConfigController(settings *config.Settings, logger *zerolog.Logger, database *sql.DB, userDeviceSvc services.UserDeviceService, deviceDefSvc services.DeviceDefinitionsService, userDeviceTemplateService services.UserDeviceTemplateService) DeviceConfigController {
	return DeviceConfigController{
		settings:                  settings,
		log:                       logger,
		db:                        database,
		userDeviceSvc:             userDeviceSvc,
		deviceDefSvc:              deviceDefSvc,
		userDeviceTemplateService: userDeviceTemplateService,
	}

}

type DeviceConfigResponse struct {
	PidURL           string `json:"pidUrl"`
	DeviceSettingURL string `json:"deviceSettingUrl"`
	DbcURL           string `json:"dbcURL,omitempty"`
	Version          string `json:"version"`
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
// @Success      200 {object} DeviceConfigResponse "Successfully retrieved configuration URLs"
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
		}
		if len(definitionResp.DeviceStyleId) > 0 {
			ud.DeviceStyleId = &definitionResp.DeviceStyleId
		}
	}

	return d.getConfigURLs(c, ud)
}

// GetConfigURLsFromEthAddr godoc
// @Description  Retrieve the URLs for PID, DeviceSettings, and DBC configuration based on device's Ethereum Address. These could be empty if not configs available
// @Tags         vehicle-signal-decoding
// @Produce      json
// @Success      200 {object} DeviceConfigResponse "Successfully retrieved configuration URLs"
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

	return d.getConfigURLs(c, ud)
}

// GetConfigStatusFromEthAddr godoc
// @Description  Retrieve the URLs for PID, DeviceSettings, and DBC configuration based on device's Ethereum Address. These could be empty if not configs available
// @Tags         vehicle-signal-decoding
// @Produce      json
// @Success      200 {object} DeviceConfigResponse "Successfully retrieved configuration URLs"
// @Failure 404  "Not Found - No templates available for the given parameters"
// @Failure 400  "incorrect eth addr format"
// @Param        ethAddr  path   string  true  "Ethereum Address"
// @Router       /device-config/eth-addr/{ethAddr}/status [get]
func (d *DeviceConfigController) GetConfigStatusFromEthAddr(c *fiber.Ctx) error {
	ethAddr := c.Params("ethAddr")

	_, err := d.userDeviceSvc.GetUserDeviceByEthAddr(c.Context(), ethAddr)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": fmt.Sprintf("no connected user device found for EthAddr: %s", ethAddr)})
	}

	return c.Send(nil)
}

func padByteArray(input []byte, targetLength int) []byte {
	if len(input) >= targetLength {
		return input // No need to pad if the input is already longer or equal to the target length
	}

	padded := make([]byte, targetLength-len(input))
	return append(padded, input...)
}

// setCANProtocol converts autopi/macaron style Protocol (6 or 7) to our VSD style protocol, but always returning a default if nothing found
func (d *DeviceConfigController) setCANProtocol(ud *pb.UserDevice) {
	switch ud.CANProtocol {
	case "6":
		ud.CANProtocol = models.CanProtocolTypeCAN11_500
	case "7":
		ud.CANProtocol = models.CanProtocolTypeCAN29_500
	case "":
		ud.CANProtocol = models.CanProtocolTypeCAN11_500
	default:
		d.log.Warn().Str("user_device_id", ud.Id).Msgf("invalid protocol detected: %s", ud.CANProtocol)
		ud.CANProtocol = models.CanProtocolTypeCAN11_500
	}
}

// retrieveAndSetVehicleInfo figures out what if any device definition information corresponds to the UserDevice.
// also calls setPowerTrainType to find and set a default Powertrain
func (d *DeviceConfigController) retrieveAndSetVehicleInfo(ctx context.Context, ud *pb.UserDevice) (string, string, int, error) {

	var ddResponse *p_grpc.GetDeviceDefinitionItemResponse
	deviceDefinitionID := ud.DeviceDefinitionId
	ddResponse, err := d.deviceDefSvc.GetDeviceDefinitionByID(ctx, deviceDefinitionID)
	if err != nil {
		return "", "", 0, fmt.Errorf("failed to retrieve device definition for deviceDefinitionId %s: %w", deviceDefinitionID, err)
	}

	vehicleYear := int(ddResponse.Type.Year)
	vehicleMake := ddResponse.Type.MakeSlug
	vehicleModel := ddResponse.Type.ModelSlug

	d.setPowerTrainType(ddResponse, ud)

	return vehicleMake, vehicleModel, vehicleYear, nil
}

func (d *DeviceConfigController) setPowerTrainType(ddResponse *p_grpc.GetDeviceDefinitionItemResponse, ud *pb.UserDevice) {
	var powerTrainType string
	for _, attribute := range ddResponse.DeviceAttributes {
		if attribute.Name == "powertrain_type" {
			powerTrainType = attribute.Value
			break
		}
	}
	if ud.PowerTrainType == "" {
		ud.PowerTrainType = powerTrainType
		if ud.PowerTrainType == "" {
			ud.PowerTrainType = "ICE"
		}
	}
}

// selectAndFetchTemplate figures out the right template to use based on the protocol, powertrain, year range, make, and /or model.
// Returns default template if nothing found. Requirees ud.CANProtocol and Powertrain to be set to something
func (d *DeviceConfigController) selectAndFetchTemplate(ctx context.Context, ud *pb.UserDevice, vehicleMake, vehicleModel string, vehicleYear int) (*models.Template, error) {
	// guard
	if ud.CANProtocol == "" {
		return nil, fmt.Errorf("CANProtocol is required in the user device")
	}
	if ud.PowerTrainType == "" {
		return nil, fmt.Errorf("PowerTrainType is required in the user device")
	}

	var matchedTemplateName string

	// First, try to find a template based on device definitions
	deviceDefinitions, err := models.TemplateDeviceDefinitions(
		models.TemplateDeviceDefinitionWhere.DeviceDefinitionID.EQ(ud.DeviceDefinitionId),
	).All(ctx, d.db)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("failed to query template device definitions: %w", err)
	}

	if len(deviceDefinitions) > 0 {
		matchedTemplateName = deviceDefinitions[0].TemplateName
	}

	// Second, try to find a template based on Year, then Make & Model
	if matchedTemplateName == "" {
		// compare by year first, then in memory below we'll look for make and/or model
		templateVehicles, err := models.TemplateVehicles(
			models.TemplateVehicleWhere.YearStart.LTE(vehicleYear),
			models.TemplateVehicleWhere.YearEnd.GTE(vehicleYear),
			qm.Load(models.TemplateVehicleRels.TemplateNameTemplate),
		).All(ctx, d.db)

		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("failed to query templates for make: %s, model: %s, year: %d: %w", vehicleMake, vehicleModel, vehicleYear, err)
		}
		// if anything is returned, try finding a match by make and/or model
		for _, tv := range templateVehicles {
			// any matches for year & same protocol
			if tv.R.TemplateNameTemplate.Protocol == ud.CANProtocol {
				matchedTemplateName = tv.TemplateName
				// now any matches for make
				if tv.MakeSlug.String == vehicleMake {
					matchedTemplateName = tv.TemplateName
					// now see if there is also a model match
					if modelMatch(tv.ModelWhitelist, vehicleModel) {
						break
					}
				}
			}
		}

	}

	// Third, fallback to query by protocol and powertrain. Match by protocol first
	if matchedTemplateName == "" {
		templates, err := models.Templates(
			models.TemplateWhere.Protocol.EQ(ud.CANProtocol),
		).All(ctx, d.db)

		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("failed to query templates for protocol: %s and powertrain: %s: %w", ud.CANProtocol, ud.PowerTrainType, err)
		}
		if len(templates) > 0 {
			// match the first one just in case
			matchedTemplateName = templates[0].TemplateName
			// now see if also have a powertrain match
			for _, t := range templates {
				if t.Powertrain == ud.PowerTrainType {
					matchedTemplateName = t.TemplateName
					break
				}
			}
		}
	}

	// Fallback to default template if still no match is found
	if matchedTemplateName == "" {
		defaultTemplates, err := models.Templates(
			qm.Where("template_name like 'default%'"),
		).All(ctx, d.db)

		if err != nil {
			return nil, fmt.Errorf("failed to query for default templates: %w", err)
		}

		if len(defaultTemplates) > 0 {
			matchedTemplateName = defaultTemplates[0].TemplateName
		} else {
			return nil, errors.New("no default templates found")
		}
	}

	// Fetch the template object if a name was found
	matchedTemplate, err := models.Templates(
		models.TemplateWhere.TemplateName.EQ(matchedTemplateName),
		qm.Load(models.TemplateRels.TemplateNameDBCFile),
		qm.Load(models.TemplateRels.TemplateNameDeviceSettings),
	).One(ctx, d.db)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch template by name %s: %w", matchedTemplateName, err)
	}

	return matchedTemplate, nil
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

func (d *DeviceConfigController) getConfigURLs(c *fiber.Ctx, ud *pb.UserDevice) error {
	d.setCANProtocol(ud)

	vehicleMake, vehicleModel, vehicleYear, err := d.retrieveAndSetVehicleInfo(c.Context(), ud)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed to retrieve device definition: %s", ud.DeviceDefinitionId))
	}

	matchedTemplate, err := d.selectAndFetchTemplate(c.Context(), ud, vehicleMake, vehicleModel, vehicleYear)
	if err != nil {
		return err
	}
	if matchedTemplate == nil {
		return errors.New("matched template is nil")
	}
	baseURL := d.settings.DeploymentURL

	response := DeviceConfigResponse{
		PidURL:  fmt.Sprintf("%s/v1/device-config/%s/pids", baseURL, matchedTemplate.TemplateName),
		Version: matchedTemplate.Version,
	}

	// only set dbc url if we have dbc
	if matchedTemplate.R.TemplateNameDBCFile != nil && len(matchedTemplate.R.TemplateNameDBCFile.DBCFile) > 0 {
		response.DbcURL = fmt.Sprintf("%s/v1/device-config/%s/dbc", baseURL, matchedTemplate.TemplateName)
	}

	// set device settings from template, or based on powertrain default
	if len(matchedTemplate.R.TemplateNameDeviceSettings) > 0 {
		response.DeviceSettingURL = fmt.Sprintf("%s/v1/device-config/settings/%s", baseURL, matchedTemplate.R.TemplateNameDeviceSettings[0].Name)
	} else {
		var deviceSetting *models.DeviceSetting
		var dbErr error
		if matchedTemplate.ParentTemplateName.Valid {
			deviceSetting, dbErr = models.DeviceSettings(models.DeviceSettingWhere.TemplateName.EQ(matchedTemplate.ParentTemplateName),
				qm.OrderBy(models.DeviceSettingColumns.Name)).One(c.Context(), d.db)
			if dbErr != nil && !errors.Is(dbErr, sql.ErrNoRows) {
				return errors.Wrap(dbErr, "Failed to retrieve device setting for parent template")
			}
		}

		if deviceSetting == nil {
			var powertrain string
			if ud.PowerTrainType != "" {
				powertrain = ud.PowerTrainType
			} else {
				powertrain = matchedTemplate.Powertrain
			}
			// default will be whatever gets ordered first
			deviceSetting, dbErr = models.DeviceSettings(models.DeviceSettingWhere.Powertrain.EQ(powertrain),
				qm.OrderBy(models.DeviceSettingColumns.Name)).One(c.Context(), d.db)
			if errors.Is(dbErr, sql.ErrNoRows) {
				// grab the first record in db
				deviceSetting, dbErr = models.DeviceSettings(qm.OrderBy(models.DeviceSettingColumns.Name)).One(c.Context(), d.db)
			}
			if dbErr != nil {
				return errors.Wrap(err, fmt.Sprintf("Failed to retrieve device setting. Powertrain: %s", powertrain))
			}
		}
		// device settings have a name key separate from templateName since simpler setup
		response.DeviceSettingURL = fmt.Sprintf("%s/v1/device-config/settings/%s", baseURL, deviceSetting.Name)
	}

	// Associate current template
	currentVersion, err := d.userDeviceTemplateService.AssociateTemplate(c.Context(), "")
	if err != nil {
		return errors.Wrap(err, "Failed to associate template version")
	}

	// todo:?
	if currentVersion.RequiresUpdateVersion {

	}

	return c.JSON(response)
}

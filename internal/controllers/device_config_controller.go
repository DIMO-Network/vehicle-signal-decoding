package controllers

import (
	"database/sql"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/queries"

	"github.com/DIMO-Network/shared/device"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/gateways"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/appmodels"

	gdata "github.com/DIMO-Network/device-data-api/pkg/grpc"

	"github.com/DIMO-Network/shared"
	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/tidwall/gjson"
	"golang.org/x/mod/semver"

	"github.com/volatiletech/sqlboiler/v4/types"

	pb "github.com/DIMO-Network/devices-api/pkg/grpc"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	_ "github.com/DIMO-Network/vehicle-signal-decoding/internal/core/appmodels" // for swagger
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
	dbs                   func() *db.ReaderWriter
	userDeviceSvc         services.UserDeviceService
	deviceDefSvc          services.DeviceDefinitionsService
	deviceTemplateService services.DeviceTemplateService
	fwVersionAPI          shared.HTTPClientWrapper
	identityAPI           gateways.IdentityAPI
}

const latestFirmwareURL = "https://binaries.dimo.zone/DIMO-Network/Macaron/releases/latest"

// NewDeviceConfigController constructor
func NewDeviceConfigController(settings *config.Settings, logger *zerolog.Logger, dbs func() *db.ReaderWriter, userDeviceSvc services.UserDeviceService,
	deviceDefSvc services.DeviceDefinitionsService, deviceTemplateService services.DeviceTemplateService, identityAPI gateways.IdentityAPI) DeviceConfigController {
	fwVersionAPI, _ := shared.NewHTTPClientWrapper(latestFirmwareURL, "", 10*time.Second, nil, true)

	return DeviceConfigController{
		settings:              settings,
		log:                   logger,
		dbs:                   dbs,
		userDeviceSvc:         userDeviceSvc,
		deviceDefSvc:          deviceDefSvc,
		deviceTemplateService: deviceTemplateService,
		fwVersionAPI:          fwVersionAPI,
		identityAPI:           identityAPI,
	}
}

// DeviceTemplateStatusResponse status on template and firmware versions
type DeviceTemplateStatusResponse struct {
	// IsTemplateUpToDate based on information we have, based on what was set last by mobile app
	IsTemplateUpToDate bool   `json:"isTemplateUpToDate"`
	FirmwareVersion    string `json:"firmwareVersion,omitempty"`
	IsFirmwareUpToDate bool   `json:"isFirmwareUpToDate"`
	// Template contains the current urls server has for this device
	Template device.ConfigResponse `json:"template"`
}

func bytesToUint32(b []byte) (uint32, error) {
	u := binary.BigEndian.Uint32(padByteArray(b, 4))
	return u, nil
}

// GetPIDsByTemplate godoc
// @Description  Retrieves a list of PID configurations from the database given a template name
// @Tags         device-config
// @Produce      json
// @Produce      application/x-protobuf
// @Success      200 {object} grpc.PIDRequests "Successfully retrieved PID Configurations"
// @Failure 404 "No PID Config data found for the given template name."
// @Param        templateName  path   string  true   "template name"
// @Router       /device-config/pids/{templateName} [get]
func (d *DeviceConfigController) GetPIDsByTemplate(c *fiber.Ctx) error {
	templateNameWithVersion := c.Params("templateName")
	// split out version
	templateName, _ := parseOutTemplateAndVersion(templateNameWithVersion)
	// ignore version for now since we're not really using it
	pidConfigs, template, err := queries.GetPidsByTemplate(c.Context(), d.dbs, &queries.GetPidsQueryRequest{
		TemplateName: templateName,
	})
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return fiber.ErrNotFound
		}
		return err
	}

	protoPIDs := &grpc.PIDRequests{
		TemplateName: templateName,
		Version:      template.Version,
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
			Formula:         pidConfig.Formula,
			IntervalSeconds: uint32(pidConfig.IntervalSeconds),
		}
		// the pid can override the protocol, otherwise use one at template level.
		if pidConfig.Protocol.Valid {
			pid.Protocol = pidConfig.Protocol.String
		} else {
			pid.Protocol = template.Protocol
		}
		if pidConfig.CanFlowControlClear.Valid {
			pid.CanFlowControlClear = pidConfig.CanFlowControlClear.Bool
		}
		if pidConfig.CanFlowControlIDPair.Valid {
			pid.CanFlowControlIdPair = pidConfig.CanFlowControlIDPair.String
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
// @Tags         device-config
// @Produce      json
// @Produce      application/x-protobuf
// @Success      200 {object} grpc.DeviceSetting "Successfully retrieved Device Settings"
// @Failure 404 "No Device Settings data found for the given name."
// @Param        name  path   string  true   "name"
// @Router       /device-config/settings/{name} [get]
func (d *DeviceConfigController) GetDeviceSettingsByName(c *fiber.Ctx) error {
	nameWithVersion := c.Params("name")
	if len(nameWithVersion) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "name for settings empty")
	}
	name, _ := parseOutTemplateAndVersion(nameWithVersion)
	// ignore version for now since we're not really using it

	dbDeviceSettings, err := models.FindDeviceSetting(c.Context(), d.dbs().Reader, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fiber.NewError(fiber.StatusNotFound, "No Device Settings data found with name: "+name)
		}
		return errors.Wrap(err, "Failed to retrieve Device Settings")
	}

	// Deserialize the settings JSONB into the SettingsData struct
	var settings appmodels.SettingsData
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
	if settings.MinVoltageOBDLoggers == 0 {
		settings.MinVoltageOBDLoggers = 13.3
	}

	if settings.LocationFrequencySecs == 0 {
		settings.LocationFrequencySecs = 20
	}

	protoDeviceSettings := &grpc.DeviceSetting{
		TemplateName:                             dbDeviceSettings.Name, // in future add a Name field, once safe to change proto
		SafetyCutOutVoltage:                      float32(settings.SafetyCutOutVoltage),
		SleepTimerEventDrivenPeriodSecs:          float32(settings.SleepTimerEventDrivenPeriodSecs),
		WakeTriggerVoltageLevel:                  float32(settings.WakeTriggerVoltageLevel),
		SleepTimerEventDrivenIntervalSecs:        float32(3600), // not used by Macaron
		SleepTimerInactivityAfterSleepSecs:       float32(21600),
		SleepTimerInactivityFallbackIntervalSecs: float32(21600),
		MinVoltageObdLoggers:                     float32(settings.MinVoltageOBDLoggers),
		LocationFrequencySecs:                    float32(settings.LocationFrequencySecs),
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
// @Tags         device-config
// @Produce      plain
// @Success      200 {string} string "Successfully retrieved DBC file"
// @Failure 404 "No DBC file found for the given template name."
// @Param        templateName  path   string  true   "template name"
// @Router       /device-config/dbc/{templateName} [get]
func (d *DeviceConfigController) GetDBCFileByTemplateName(c *fiber.Ctx) error {
	templateNameWithVersion := c.Params("templateName")
	templateName, _ := parseOutTemplateAndVersion(templateNameWithVersion)
	// ignore version since not really using right now

	dbResult, err := models.DBCFiles(
		models.DBCFileWhere.TemplateName.EQ(templateName)).One(c.Context(), d.dbs().Reader)

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
// @Tags         device-config
// @Produce      json
// @Success      200 {object} device.ConfigResponse "Successfully retrieved configuration URLs"
// @Failure 404  "Not Found - No templates available for the given parameters"
// @Param        vin  path   string  true   "vehicle identification number (VIN)"
// @Param        protocol  query   string  false  "CAN Protocol, '6' or '7', 8,9,66,77,88,99"
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

	response, strategy, err := d.deviceTemplateService.ResolveDeviceConfiguration(c, ud, nil)
	if err != nil {
		return err
	}
	d.log.Info().Str("vin", *ud.Vin).Msgf(fmt.Sprintf("template configuration urls for VIN %s. strategy: %s. dbc: %s, pids: %s, settings: %s",
		*ud.Vin, strategy, response.DbcURL, response.PidURL, response.DeviceSettingURL))

	return c.JSON(response)
}

// GetConfigURLsFromEthAddr godoc
// @Description  Retrieve the URLs for PID, DeviceSettings, and DBC configuration based on device's Ethereum Address. These could be empty if not configs available
// @Tags         device-config
// @Produce      json
// @Success      200 {object} device.ConfigResponse "Successfully retrieved configuration URLs"
// @Failure 404  "Not Found - No templates available for the given parameters"
// @Failure 400  "incorrect eth addr format"
// @Param        ethAddr  path   string  true  "Ethereum Address"
// @Param        protocol  query   string  false  "CAN Protocol, '6' or '7'"
// @Router       /device-config/eth-addr/{ethAddr}/urls [get]
func (d *DeviceConfigController) GetConfigURLsFromEthAddr(c *fiber.Ctx) error {
	ethAddr := c.Params("ethAddr")
	protocol := c.Query("protocol", "")
	address := common2.HexToAddress(ethAddr)

	// first check for direct mapping
	directConfig := d.deviceTemplateService.FindDirectDeviceToTemplateConfig(c.Context(), address)
	if directConfig != nil {
		d.log.Info().Str("ethAddr", ethAddr).Msgf(fmt.Sprintf("template configuration urls for eth addr %s. strategy: direct. dbc: %s, pids: %s, settings: %s",
			ethAddr, directConfig.DbcURL, directConfig.PidURL, directConfig.DeviceSettingURL))
		return c.JSON(directConfig)
	}

	vehicle, err := d.identityAPI.QueryIdentityAPIForVehicle(address)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": fmt.Sprintf("no minted vehicle for device EthAddr: %s", ethAddr)})
	}
	// we still need this to get the powertrain
	ud, err := d.userDeviceSvc.GetUserDeviceByEthAddr(c.Context(), address)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": fmt.Sprintf("no connected user device found for EthAddr: %s", ethAddr)})
	}

	if protocol != "" {
		ud.CANProtocol = protocol
	}

	response, strategy, err := d.deviceTemplateService.ResolveDeviceConfiguration(c, ud, vehicle)
	if err != nil {
		return err
	}

	d.log.Info().Str("vin", *ud.Vin).Msgf(fmt.Sprintf("template configuration urls for VIN %s and eth Addr: %s. strategy: %s. dbc: %s, pids: %s, settings: %s",
		*ud.Vin, ethAddr, strategy, response.DbcURL, response.PidURL, response.DeviceSettingURL))

	return c.JSON(response)
}

// GetConfigStatusByEthAddr godoc
// @Description  Helps client determine if template (pids, dbc, settings) are up to date or not for the device with the given eth addr.
// @Tags         device-config
// @Produce      json
// @Success      200 {object} DeviceTemplateStatusResponse "Successfully retrieved configuration URLs"
// @Failure 404  "Not Found - we haven't seen this device yet, assume template not up to date"
// @Failure 400  "incorrect eth addr format"
// @Param        ethAddr  path   string  true  "Ethereum Address"
// @Router       /device-config/eth-addr/{ethAddr}/status [get]
func (d *DeviceConfigController) GetConfigStatusByEthAddr(c *fiber.Ctx) error {
	ethAddr := c.Params("ethAddr")
	addr := common2.HexToAddress(ethAddr)

	// we use this to know what the config should be
	ud, err := d.userDeviceSvc.GetUserDeviceByEthAddr(c.Context(), addr)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": fmt.Sprintf("no connected user device found for EthAddr: %s", ethAddr)})
	}

	dts, err := models.DeviceTemplateStatuses(models.DeviceTemplateStatusWhere.DeviceEthAddr.EQ(addr.Bytes())).One(c.Context(), d.dbs().Reader)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
	}
	deviceFWVers := ""
	isTemplateUpdated := false
	if dts != nil {
		deviceFWVers = dts.FirmwareVersion.String
		// figure out what the config should be
		deviceConfiguration, _, err := d.deviceTemplateService.ResolveDeviceConfiguration(c, ud, nil)
		if err != nil {
			return err
		}
		// if all this is true then we know we're up to date
		if dts.TemplateDBCURL.String == deviceConfiguration.DbcURL &&
			dts.TemplatePidURL.String == deviceConfiguration.PidURL &&
			dts.TemplateSettingsURL.String == deviceConfiguration.DeviceSettingURL {

			isTemplateUpdated = true
		}
	}

	if deviceFWVers == "" {
		// get fw version from device if any
		deviceData, err := d.userDeviceSvc.GetRawDeviceData(c.Context(), ud.Id)
		if err != nil {
			d.log.Err(err).Str("device_address", ethAddr).Msg("failed to get device data")
			if dts == nil {
				// if don't get anything from device-data-api and dts is nil, nothing we can do
				return fiber.NewError(fiber.StatusNotFound, "haven't seen device with eth addr yet: "+ethAddr)
			}
		}
		deviceFWVers = parseOutFWVersion(deviceData)
	}
	latestFirmwareStr, err := d.getLatestFWVersion()
	if err != nil {
		return err
	}
	resp := DeviceTemplateStatusResponse{
		IsTemplateUpToDate: isTemplateUpdated,
		IsFirmwareUpToDate: isFwUpToDate(latestFirmwareStr, deviceFWVers),
		FirmwareVersion:    deviceFWVers,
	}
	if dts != nil {
		resp.Template.DbcURL = dts.TemplateDBCURL.String
		resp.Template.PidURL = dts.TemplatePidURL.String
		resp.Template.DeviceSettingURL = dts.TemplateSettingsURL.String
	}
	return c.JSON(resp)
}

// PatchConfigStatusByEthAddr godoc
// @Description  Set what template and/or firmware was applied. None of the properties are required. Will not be set if not passed in.
// @Tags         device-config
// @Produce      json
// @Success      200 "Successfully updated"
// @Failure 500  "unable to parse request or storage failure"
// @Param        ethAddr  path   string  true  "Ethereum Address"
// @Param        config body DeviceTemplateStatusPatch true "set any properties that were updated on the device"
// @Security     BearerAuth
// @Router       /device-config/eth-addr/{ethAddr}/status [patch]
func (d *DeviceConfigController) PatchConfigStatusByEthAddr(c *fiber.Ctx) error {
	ethAddr := c.Params("ethAddr")
	addr := common2.HexToAddress(ethAddr)

	payload := DeviceTemplateStatusPatch{}
	err := c.BodyParser(&payload)
	if err != nil {
		return err
	}

	// control for deprecated properties still used by mobile app for patching
	if payload.PidsURL != "" {
		payload.PidURL = payload.PidsURL
	}
	if payload.SettingsURL != "" {
		payload.DeviceSettingURL = payload.SettingsURL
	}

	_, err = d.deviceTemplateService.StoreDeviceConfigUsed(c.Context(), addr, payload.DbcURL, payload.PidURL, payload.DeviceSettingURL, payload.FirmwareVersionApplied)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

// PatchHwConfigStatusByEthAddr godoc
// @Description  Set what template and/or firmware was applied. None of the properties are required. Will not be set if not passed in. Endpoint is meant only for hardware devices self-reporting their template update.
// @Tags         device-config
// @Produce      json
// @Success      200 "Successfully updated"
// @Failure 500  "unable to parse request or storage failure"
// @Param        ethAddr  path   string  true  "Ethereum Address"
// @Param        config body DeviceTemplateStatusPatch true "set any properties that were updated on the device"
// @Security     SignatureAuth
// @Router       /device-config/eth-addr/{ethAddr}/hw/status [patch]
func (d *DeviceConfigController) PatchHwConfigStatusByEthAddr(c *fiber.Ctx) error {
	return d.PatchConfigStatusByEthAddr(c)
}

func parseOutFWVersion(data *gdata.RawDeviceDataResponse) string {
	for _, item := range data.Items {
		v := gjson.GetBytes(item.SignalsJsonData, "fwVersion.value").Str
		if len(v) > 1 {
			if v[0:1] != "v" {
				return "v" + v
			}
			return v
		}
	}
	return ""
}

func parseOutTemplateAndVersion(templateNameWithVersion string) (string, string) {
	parts := strings.Split(templateNameWithVersion, "@")
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	return parts[0], ""
}

type DeviceTemplateStatusPatch struct {
	device.ConfigResponse
	// FirmwareVersionApplied version of firmware that was confirmed installed on device
	FirmwareVersionApplied string `json:"firmwareVersionApplied"`

	// PidsURL exists for backwards compatibility
	// Deprecated
	PidsURL string `json:"pidsUrl"`

	// SettingsUrl exists for backwards compatibiltiy
	// Deprecated
	SettingsURL string `json:"settingsUrl"`
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

func isFwUpToDate(latest, current string) bool {
	if len(latest) > 1 && len(current) > 1 {
		if latest[0:1] != "v" {
			latest = "v" + latest
		}
		if current[0:1] != "v" {
			current = "v" + current
		}
		if semver.Compare(latest, current) == 0 {
			return true
		}
	}
	return false
}

// calls well known dimo URL to get latest Macaron fw version
func (d *DeviceConfigController) getLatestFWVersion() (string, error) {
	// get latest fw version. at some point will need to know device hw type to know this better
	res, err := d.fwVersionAPI.ExecuteRequest("", "GET", nil)
	if err != nil {
		return "", errors.Wrap(err, "unable to get latest macaron firmware")
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	latestFirmwareStr := gjson.GetBytes(body, "name").Str

	return latestFirmwareStr, nil
}

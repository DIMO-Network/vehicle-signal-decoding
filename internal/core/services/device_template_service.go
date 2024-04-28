package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/DIMO-Network/shared"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/gateways"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/appmodels"
	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/volatiletech/null/v8"

	pb "github.com/DIMO-Network/devices-api/pkg/grpc"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/types"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"database/sql"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
)

//go:generate mockgen -source device_template_service.go -destination mocks/device_template_service_mock.go
type DeviceTemplateService interface {
	StoreDeviceConfigUsed(ctx context.Context, address common2.Address, dbcURL, pidURL, settingURL, firmwareVersion *string) (*models.DeviceTemplateStatus, error)
	ResolveDeviceConfiguration(c *fiber.Ctx, ud *pb.UserDevice, vehicle *gateways.VehicleInfo) (*appmodels.DeviceConfigResponse, error)
	// todo: pass in a ResolveConfigRequest instead of pb.UserDevice - this is not tied to a user device
}

type deviceTemplateService struct {
	db           *sql.DB
	log          zerolog.Logger
	settings     *config.Settings
	deviceDefSvc DeviceDefinitionsService
}

func NewDeviceTemplateService(database *sql.DB, deviceDefSvc DeviceDefinitionsService, log zerolog.Logger, settings *config.Settings) DeviceTemplateService {
	return &deviceTemplateService{
		db:           database,
		log:          log,
		settings:     settings,
		deviceDefSvc: deviceDefSvc,
	}
}

// StoreDeviceConfigUsed stores the configurations that were used by the mobile app to apply onto the device
func (dts *deviceTemplateService) StoreDeviceConfigUsed(ctx context.Context, address common2.Address, dbcURL, pidURL, settingURL, firmwareVersion *string) (*models.DeviceTemplateStatus, error) {

	dt, err := models.DeviceTemplateStatuses(models.DeviceTemplateStatusWhere.DeviceEthAddr.EQ(address.Bytes())).
		One(ctx, dts.db)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if dt != nil {
		// update - only set if not nil
		if settingURL != nil {
			dt.TemplateSettingsURL = null.StringFromPtr(settingURL)
		}
		if dbcURL != nil {
			dt.TemplateDBCURL = null.StringFromPtr(dbcURL)
		}
		if pidURL != nil {
			dt.TemplatePidURL = null.StringFromPtr(pidURL)
		}
		if firmwareVersion != nil {
			fwv := *firmwareVersion
			if len(fwv) > 1 {
				if fwv[0:1] != "v" {
					fwv = "v" + fwv
				}
				dt.FirmwareVersion = null.StringFrom(fwv)
			}
		}

		if _, err = dt.Update(ctx, dts.db, boil.Infer()); err != nil {
			return nil, err
		}
	} else {
		// create
		dt = &models.DeviceTemplateStatus{
			DeviceEthAddr:       address.Bytes(),
			TemplateDBCURL:      null.StringFromPtr(dbcURL),
			TemplatePidURL:      null.StringFromPtr(pidURL),
			TemplateSettingsURL: null.StringFromPtr(settingURL),
			FirmwareVersion:     null.StringFromPtr(firmwareVersion),
		}
		if err = dt.Insert(ctx, dts.db, boil.Infer()); err != nil {
			return nil, err
		}
	}

	return dt, nil
}

func (dts *deviceTemplateService) FindDirectDeviceToTemplateConfig(ctx context.Context, address common2.Address) *appmodels.DeviceConfigResponse {
	deviceToTemplate, err := models.AftermarketDeviceToTemplates(
		models.AftermarketDeviceToTemplateWhere.AftermarketDeviceEthereumAddress.EQ(address.Bytes()),
		qm.Load(models.AftermarketDeviceToTemplateRels)).One(ctx, dts.db)
	if err != nil || deviceToTemplate == nil {
		return nil
	}
	response := appmodels.DeviceConfigResponse{
		PidURL: dts.buildConfigRoute(PIDs, deviceToTemplate.TemplateName, matchedTemplate.Version),
	}

	return &response
}

// ResolveDeviceConfiguration figures out what template to return based on protocol, powertrain, vehicle or definition (vehicle could be nil)
func (dts *deviceTemplateService) ResolveDeviceConfiguration(c *fiber.Ctx, ud *pb.UserDevice, vehicle *gateways.VehicleInfo) (*appmodels.DeviceConfigResponse, error) {
	canProtocl := convertCANProtocol(dts.log, ud.CANProtocol)
	// todo (jreate): what about powertrain at the style level... But ideally it is stored at vehicle level. this could come from oracle?
	powertrain, err := dts.retrievePowertrain(c.Context(), ud.DeviceDefinitionId)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed to retrieve powertrain for ddid: %s", ud.DeviceDefinitionId))
	}

	matchedTemplate, err := dts.selectAndFetchTemplate(c.Context(), canProtocl, powertrain, ud.DeviceDefinitionId, vehicle)
	if err != nil {
		return nil, err
	}
	if matchedTemplate == nil {
		return nil, errors.New("matched template is nil")
	}

	response := appmodels.DeviceConfigResponse{
		PidURL: dts.buildConfigRoute(PIDs, matchedTemplate.TemplateName, matchedTemplate.Version),
	}

	// only set dbc url if we have dbc
	if matchedTemplate.R.TemplateNameDBCFile != nil && len(matchedTemplate.R.TemplateNameDBCFile.DBCFile) > 0 {
		response.DbcURL = dts.buildConfigRoute(DBC, matchedTemplate.TemplateName, matchedTemplate.Version)
	}

	// set device settings from template, or based on powertrain default
	if len(matchedTemplate.R.TemplateNameDeviceSettings) > 0 {
		ds := matchedTemplate.R.TemplateNameDeviceSettings[0]
		response.DeviceSettingURL = dts.buildConfigRoute(Setting, ds.Name, ds.Version)
	} else {
		var deviceSetting *models.DeviceSetting
		var dbErr error
		if matchedTemplate.ParentTemplateName.Valid {
			deviceSetting, dbErr = models.DeviceSettings(models.DeviceSettingWhere.TemplateName.EQ(matchedTemplate.ParentTemplateName),
				qm.OrderBy(models.DeviceSettingColumns.Name)).One(c.Context(), dts.db)
			if dbErr != nil && !errors.Is(dbErr, sql.ErrNoRows) {
				return nil, errors.Wrap(dbErr, "Failed to retrieve device setting for parent template")
			}
		}

		if deviceSetting == nil {
			var pt string
			if ud.PowerTrainType != "" {
				pt = ud.PowerTrainType
			} else {
				pt = matchedTemplate.Powertrain
			}
			// default will be whatever gets ordered first
			deviceSetting, dbErr = models.DeviceSettings(models.DeviceSettingWhere.Powertrain.EQ(pt),
				qm.OrderBy(models.DeviceSettingColumns.Name)).One(c.Context(), dts.db)
			if errors.Is(dbErr, sql.ErrNoRows) {
				// grab the first record in db
				deviceSetting, dbErr = models.DeviceSettings(qm.OrderBy(models.DeviceSettingColumns.Name)).One(c.Context(), dts.db)
			}
			if dbErr != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("Failed to retrieve device setting. Powertrain: %s", pt))
			}
		}
		// device settings have a name key separate from templateName since simpler setup
		response.DeviceSettingURL = dts.buildConfigRoute(Setting, deviceSetting.Name, deviceSetting.Version)
	}

	dts.log.Info().Str("vin", *ud.Vin).Msgf(fmt.Sprintf("template configuration urls for VIN %s, dbc: %s, pids: %s, settings: %s",
		*ud.Vin, response.DbcURL, response.PidURL, response.DeviceSettingURL))

	return &response, nil
}

type configType string

const (
	PIDs    = "pids"
	Setting = "settings"
	DBC     = "dbc"
)

func (dts *deviceTemplateService) buildConfigRoute(ct configType, name, version string) string {
	return fmt.Sprintf("%s/v1/device-config/%s/%s@%s", dts.settings.DeploymentURL, ct, name, version)
}

// retrievePowertrain gets the powertrain for the device definition id from attributes, if empty defaults to ICE
func (dts *deviceTemplateService) retrievePowertrain(ctx context.Context, deviceDefinitionID string) (string, error) {
	ddResponse, err := dts.deviceDefSvc.GetDeviceDefinitionByID(ctx, deviceDefinitionID)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve device definition for deviceDefinitionId %s: %w", deviceDefinitionID, err)
	}

	var powerTrainType string
	for _, attribute := range ddResponse.DeviceAttributes {
		if attribute.Name == "powertrain_type" {
			powerTrainType = attribute.Value
			break
		}
	}
	if powerTrainType == "" {
		powerTrainType = "ICE"
	}

	return powerTrainType, nil
}

// selectAndFetchTemplate figures out the right template to use based on the protocol, powertrain, year range, make, and /or model.
// Returns default template if nothing found. Requirees ud.CANProtocol and Powertrain to be set to something
func (dts *deviceTemplateService) selectAndFetchTemplate(ctx context.Context, canProtocol, powertrain, definitionID string, vehicle *gateways.VehicleInfo) (*models.Template, error) {
	// guard
	if canProtocol == "" {
		return nil, fmt.Errorf("CANProtocol is required in the user device")
	}
	if powertrain == "" {
		return nil, fmt.Errorf("PowerTrainType is required in the user device")
	}

	var matchedTemplateName string

	// First, try to find a template based on device definitions
	deviceDefinitions, err := models.TemplateDeviceDefinitions(
		models.TemplateDeviceDefinitionWhere.DeviceDefinitionID.EQ(definitionID),
	).All(ctx, dts.db)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("failed to query template device definitions: %w", err)
	}

	if len(deviceDefinitions) > 0 {
		matchedTemplateName = deviceDefinitions[0].TemplateName
	}
	year := 0
	mk := ""
	model := ""
	if vehicle == nil {
		definition, err := dts.deviceDefSvc.GetDeviceDefinitionByID(ctx, definitionID)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to query device definition %s", definitionID)
		}
		year = int(definition.Type.Year)
		mk = definition.Type.Make
		model = definition.Type.Model
	} else {
		year = vehicle.Definition.Year
		mk = vehicle.Definition.Make
		model = vehicle.Definition.Model
	}

	// Second, try to find a template based on Year, then Make & Model
	if matchedTemplateName == "" {
		// compare by year first, then in memory below we'll look for make and/or model
		templateVehicles, err := models.TemplateVehicles(
			models.TemplateVehicleWhere.YearStart.LTE(year),
			models.TemplateVehicleWhere.YearEnd.GTE(year),
			qm.Load(models.TemplateVehicleRels.TemplateNameTemplate),
		).All(ctx, dts.db)

		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("failed to query templates for vehicle: %s: %w", fmt.Sprintf("%d %s %s", year, mk, model), err)
		}
		// if anything is returned, try finding a match by make and/or model
		for _, tv := range templateVehicles {
			// any matches for year & same protocol
			if tv.R.TemplateNameTemplate.Protocol == canProtocol {
				matchedTemplateName = tv.TemplateName
				// now any matches for make
				if tv.MakeSlug.String == shared.SlugString(mk) {
					matchedTemplateName = tv.TemplateName
					// now see if there is also a model slug match
					if modelMatch(tv.ModelWhitelist, shared.SlugString(model)) {
						break
					}
				}
			}
		}

	}
	// todo here is the problem - we're not filtering by "default" and powertrain and canprotocol
	// todo thought: iterate again after filtering by powertrain for "default" startword.
	// Third, default templates come into play: fallback to query by protocol, 'default' as first word, and powertrain
	if matchedTemplateName == "" {
		templates, err := models.Templates(
			models.TemplateWhere.Protocol.EQ(canProtocol),
			models.TemplateWhere.Powertrain.EQ(powertrain),
			qm.Where("template_name like 'default%'"),
		).All(ctx, dts.db)

		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("failed to query templates for protocol: %s and powertrain: %s: %w", canProtocol, powertrain, err)
		}
		if len(templates) == 0 {
			return nil, fmt.Errorf("configuration error - no default template found for protocol: %s and powertrain: %s", canProtocol, powertrain)
		}
		if len(templates) > 0 {
			matchedTemplateName = templates[0].TemplateName
		}
		if len(templates) > 1 {
			dts.log.Warn().Msgf("more than one default template found for protocol: %s and powertrain: %s (%d templates found)", canProtocol, powertrain, len(templates))
		}
	}

	// Fetch the template object if a name was found
	matchedTemplate, err := models.Templates(
		models.TemplateWhere.TemplateName.EQ(matchedTemplateName),
		qm.Load(models.TemplateRels.TemplateNameDBCFile),
		qm.Load(models.TemplateRels.TemplateNameDeviceSettings),
	).One(ctx, dts.db)
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

// convertCANProtocol converts autopi/macaron style Protocol (6 or 7) to our VSD style protocol (db enum), but always returning a default if nothing found
func convertCANProtocol(logger zerolog.Logger, canProtocolSimple string) string {
	switch canProtocolSimple {
	case "6":
		return models.CanProtocolTypeCAN11_500
	case "7":
		return models.CanProtocolTypeCAN29_500
	case "8":
		return models.CanProtocolTypeCAN11_250
	case "9":
		return models.CanProtocolTypeCAN29_250
	case "66":
		// car supports UDS vin query
		return models.CanProtocolTypeCAN11_500
	case "77":
		// car supports UDS vin query
		return models.CanProtocolTypeCAN29_500
	case "88":
		// car supports UDS vin query
		return models.CanProtocolTypeCAN11_250
	case "99":
		// car supports UDS vin query
		return models.CanProtocolTypeCAN29_250
	case "":
		return models.CanProtocolTypeCAN11_500
	default:
		logger.Warn().Msgf("invalid protocol detected: %s", canProtocolSimple)
		return models.CanProtocolTypeCAN11_500
	}
}

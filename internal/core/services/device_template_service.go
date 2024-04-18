package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/appmodels"
	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/volatiletech/null/v8"

	pgrpc "github.com/DIMO-Network/device-definitions-api/pkg/grpc"
	pb "github.com/DIMO-Network/devices-api/pkg/grpc"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
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
	ResolveDeviceConfiguration(ctx context.Context, ud *pb.UserDevice, ethAddress common2.Address) (*appmodels.DeviceConfigResponse, error)
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

func (dts *deviceTemplateService) ResolveDeviceConfiguration(ctx context.Context, ud *pb.UserDevice, ethAddress common2.Address) (*appmodels.DeviceConfigResponse, error) {
	dts.setCANProtocol(ud)

	vehicleMake, vehicleModel, vehicleYear, err := dts.retrieveAndSetVehicleInfo(ctx, ud)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed to retrieve device definition: %s", ud.DeviceDefinitionId))
	}

	matchedTemplate, err := dts.selectAndFetchTemplate(ctx, ud, vehicleMake, vehicleModel, vehicleYear)
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
			var powertrain string
			if ud.PowerTrainType != "" {
				powertrain = ud.PowerTrainType
			} else {
				powertrain = matchedTemplate.Powertrain
			}
			// default will be whatever gets ordered first
			deviceSetting, dbErr = models.DeviceSettings(models.DeviceSettingWhere.Powertrain.EQ(powertrain),
				qm.OrderBy(models.DeviceSettingColumns.Name)).One(ctx, dts.db)
			if errors.Is(dbErr, sql.ErrNoRows) {
				// grab the first record in db
				deviceSetting, dbErr = models.DeviceSettings(qm.OrderBy(models.DeviceSettingColumns.Name)).One(ctx, dts.db)
			}
			if dbErr != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("Failed to retrieve device setting. Powertrain: %s", powertrain))
			}
		}
		// device settings have a name key separate from templateName since simpler setup
		response.DeviceSettingURL = dts.buildConfigRoute(Setting, deviceSetting.Name, deviceSetting.Version)
	}

	aftermarketDeviceTemplate, err := models.AftermarketDeviceToTemplates(models.AftermarketDeviceToTemplateWhere.AftermarketDeviceEthereumAddress.EQ(ethAddress.Bytes())).One(ctx, dts.db)

	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve aftermarket device template")
	}

	if aftermarketDeviceTemplate != nil {

		afktemplate, err := models.Templates(models.TemplateWhere.TemplateName.EQ(aftermarketDeviceTemplate.TemplateName)).One(ctx, dts.db)

		if err != nil {
			return nil, errors.Wrap(err, "Failed to retrieve aftermarket device template")
		}

		response.DeviceTemplateURL = dts.buildConfigRoute(Setting, afktemplate.TemplateName, afktemplate.Version)
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

// retrieveAndSetVehicleInfo figures out what if any device definition information corresponds to the UserDevice.
// also calls setPowerTrainType to find and set a default Powertrain, returns Make, Model, Year.
func (dts *deviceTemplateService) retrieveAndSetVehicleInfo(ctx context.Context, ud *pb.UserDevice) (string, string, int, error) {

	var ddResponse *pgrpc.GetDeviceDefinitionItemResponse
	deviceDefinitionID := ud.DeviceDefinitionId
	ddResponse, err := dts.deviceDefSvc.GetDeviceDefinitionByID(ctx, deviceDefinitionID)
	if err != nil {
		return "", "", 0, fmt.Errorf("failed to retrieve device definition for deviceDefinitionId %s: %w", deviceDefinitionID, err)
	}

	vehicleYear := int(ddResponse.Type.Year)
	vehicleMake := ddResponse.Type.MakeSlug
	vehicleModel := ddResponse.Type.ModelSlug

	setPowerTrainType(ddResponse, ud)

	return vehicleMake, vehicleModel, vehicleYear, nil
}

func setPowerTrainType(ddResponse *pgrpc.GetDeviceDefinitionItemResponse, ud *pb.UserDevice) {
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
func (dts *deviceTemplateService) selectAndFetchTemplate(ctx context.Context, ud *pb.UserDevice, vehicleMake, vehicleModel string, vehicleYear int) (*models.Template, error) {
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
	).All(ctx, dts.db)

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
		).All(ctx, dts.db)

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
		).All(ctx, dts.db)

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
		).All(ctx, dts.db)

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

// setCANProtocol converts autopi/macaron style Protocol (6 or 7) to our VSD style protocol, but always returning a default if nothing found
func (dts *deviceTemplateService) setCANProtocol(ud *pb.UserDevice) {
	switch ud.CANProtocol {
	case "6":
		ud.CANProtocol = models.CanProtocolTypeCAN11_500
	case "7":
		ud.CANProtocol = models.CanProtocolTypeCAN29_500
	case "":
		ud.CANProtocol = models.CanProtocolTypeCAN11_500
	default:
		dts.log.Warn().Str("user_device_id", ud.Id).Msgf("invalid protocol detected: %s", ud.CANProtocol)
		ud.CANProtocol = models.CanProtocolTypeCAN11_500
	}
}

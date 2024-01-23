package services

import (
	"context"
	"fmt"
	"strings"

	localmodels "github.com/DIMO-Network/vehicle-signal-decoding/internal/core/models"

	p_grpc "github.com/DIMO-Network/device-definitions-api/pkg/grpc"
	pb "github.com/DIMO-Network/devices-api/pkg/grpc"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/types"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"database/sql"

	models "github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
)

//go:generate mockgen -source device_template_service.go -destination mocks/device_template_service_mock.go
type DeviceTemplateService interface {
	StoreLastTemplateRequested(ctx context.Context, vin, templateDbcURL, templatePidURL, templateSettingURL, version string) (*models.DeviceTemplateStatus, error)
	ResolveDeviceConfiguration(c *fiber.Ctx, ud *pb.UserDevice) (*localmodels.DeviceConfigResponse, error)
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

// StoreLastTemplateRequested stores the last template urls & version requested for a given vin
func (dts *deviceTemplateService) StoreLastTemplateRequested(ctx context.Context, vin, templateDbcURL, templatePidURL, templateSettingURL, version string) (*models.DeviceTemplateStatus, error) {

	dt, err := models.DeviceTemplateStatuses(models.DeviceTemplateStatusWhere.Vin.EQ(vin)).
		One(ctx, dts.db)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if dt != nil && dt.TemplateVersion != version && dt.TemplateSettingURL != templateSettingURL {
		dt.TemplateVersion = version
		dt.TemplateSettingURL = templateSettingURL
		dt.TemplateDBCURL = templateDbcURL
		dt.TemplatePidURL = templatePidURL

		if _, err = dt.Update(ctx, dts.db, boil.Infer()); err != nil {
			return nil, err
		}
	}

	if dt == nil {
		dt = &models.DeviceTemplateStatus{
			Vin:                vin,
			TemplateDBCURL:     templateDbcURL,
			TemplatePidURL:     templatePidURL,
			TemplateSettingURL: templateSettingURL,
		}

		if err = dt.Insert(ctx, dts.db, boil.Infer()); err != nil {
			return nil, err
		}
	}

	return dt, nil
}

func (dts *deviceTemplateService) ResolveDeviceConfiguration(c *fiber.Ctx, ud *pb.UserDevice) (*localmodels.DeviceConfigResponse, error) {
	dts.setCANProtocol(ud)

	vehicleMake, vehicleModel, vehicleYear, err := dts.retrieveAndSetVehicleInfo(c.Context(), ud)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed to retrieve device definition: %s", ud.DeviceDefinitionId))
	}

	matchedTemplate, err := dts.selectAndFetchTemplate(c.Context(), ud, vehicleMake, vehicleModel, vehicleYear)
	if err != nil {
		return nil, err
	}
	if matchedTemplate == nil {
		return nil, errors.New("matched template is nil")
	}
	baseURL := dts.settings.DeploymentURL

	response := localmodels.DeviceConfigResponse{
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
				qm.OrderBy(models.DeviceSettingColumns.Name)).One(c.Context(), dts.db)
			if errors.Is(dbErr, sql.ErrNoRows) {
				// grab the first record in db
				deviceSetting, dbErr = models.DeviceSettings(qm.OrderBy(models.DeviceSettingColumns.Name)).One(c.Context(), dts.db)
			}
			if dbErr != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("Failed to retrieve device setting. Powertrain: %s", powertrain))
			}
		}
		// device settings have a name key separate from templateName since simpler setup
		response.DeviceSettingURL = fmt.Sprintf("%s/v1/device-config/settings/%s", baseURL, deviceSetting.Name)
	}

	dts.log.Info().Str("vin", *ud.Vin).Msgf(fmt.Sprintf("template configuration urls for VIN %s, dbc: %s, pids: %s, settings: %s, version: %s",
		*ud.Vin, response.DbcURL, response.PidURL, response.DeviceSettingURL, response.Version))

	return &response, nil
}

// retrieveAndSetVehicleInfo figures out what if any device definition information corresponds to the UserDevice.
// also calls setPowerTrainType to find and set a default Powertrain, returns Make, Model, Year.
func (dts *deviceTemplateService) retrieveAndSetVehicleInfo(ctx context.Context, ud *pb.UserDevice) (string, string, int, error) {

	var ddResponse *p_grpc.GetDeviceDefinitionItemResponse
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

func setPowerTrainType(ddResponse *p_grpc.GetDeviceDefinitionItemResponse, ud *pb.UserDevice) {
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
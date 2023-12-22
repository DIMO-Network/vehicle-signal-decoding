package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/DIMO-Network/shared/db"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"

	"github.com/volatiletech/null/v8"

	_ "github.com/lib/pq"

	"github.com/volatiletech/sqlboiler/v4/types"

	p_grpc "github.com/DIMO-Network/device-definitions-api/pkg/grpc"

	pb "github.com/DIMO-Network/devices-api/pkg/grpc"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	mock_services "github.com/DIMO-Network/vehicle-signal-decoding/internal/core/services/mocks"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/dbtest"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

const migrationsDirRelPath = "../infrastructure/db/migrations"

type DeviceConfigControllerTestSuite struct {
	suite.Suite
	ctx               context.Context
	pdb               db.Store
	container         testcontainers.Container
	mockCtrl          *gomock.Controller
	logger            *zerolog.Logger
	mockUserDeviceSvc *mock_services.MockUserDeviceService
	mockDeviceDefSvc  *mock_services.MockDeviceDefinitionsService
	controller        *DeviceConfigController
	app               *fiber.App
}

const dbSchemaName = "vehicle_signal_decoding"

func (s *DeviceConfigControllerTestSuite) SetupSuite() {
	s.ctx = context.Background()
	s.pdb, s.container = dbtest.StartContainerDatabase(s.ctx, dbSchemaName, s.T(), migrationsDirRelPath)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	s.logger = &logger
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserDeviceSvc = mock_services.NewMockUserDeviceService(s.mockCtrl)
	s.mockDeviceDefSvc = mock_services.NewMockDeviceDefinitionsService(s.mockCtrl)
	ctrl := NewDeviceConfigController(&config.Settings{Port: "3000", DeploymentURL: "http://localhost:3000"}, s.logger, s.pdb.DBS().Reader.DB, s.mockUserDeviceSvc, s.mockDeviceDefSvc)
	s.controller = &ctrl
	s.app = fiber.New()
}

func (s *DeviceConfigControllerTestSuite) SetupTest() {
	dbtest.TruncateTables(s.pdb.DBS().Writer.DB, dbSchemaName, s.T())
}

func (s *DeviceConfigControllerTestSuite) TearDownSuite() {
	if err := s.container.Terminate(s.ctx); err != nil {
		s.T().Fatal(err)
	}
	s.mockCtrl.Finish()
}

func (s *DeviceConfigControllerTestSuite) TearDownTest() {
	dbtest.TruncateTables(s.pdb.DBS().Writer.DB, dbSchemaName, s.T())
	temps, err := models.Templates().All(context.Background(), s.pdb.DBS().Writer)
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, temp := range temps {
		fmt.Printf("template: %+v\n", temp)
	}
}

func TestDeviceConfigControllerTestSuite(t *testing.T) {
	suite.Run(t, new(DeviceConfigControllerTestSuite))
}

/* Actual Tests */
func (s *DeviceConfigControllerTestSuite) TestGetPIDsByTemplate() {

	template := models.Template{
		TemplateName: "exampleTemplate",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN11_500,
		Powertrain:   "ICE",
	}
	err := template.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	s.Require().NoError(err)

	pc := models.PidConfig{
		ID:              1,
		SignalName:      "odometer",
		TemplateName:    "exampleTemplate",
		Header:          []byte{0x07, 0xdf},
		Mode:            []byte{0x01},
		Pid:             []byte{0xa6},
		Formula:         "A*5",
		IntervalSeconds: 60,
		Protocol:        models.CanProtocolTypeCAN11_500,
	}

	err = pc.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	s.Require().NoError(err)

	s.app.Get("/device-config/:templateName/pids", s.controller.GetPIDsByTemplate)

	request := dbtest.BuildRequest("GET", "/device-config/"+template.TemplateName+"/pids", "")
	response, err := s.app.Test(request)
	s.Require().NoError(err)

	body, err := io.ReadAll(response.Body)
	s.Require().NoError(err)

	s.Equal(fiber.StatusOK, response.StatusCode)
	pids := grpc.PIDRequests{}
	err = json.Unmarshal(body, &pids)
	s.Require().NoError(err)

	s.Equal(1, len(pids.Requests))
	s.Equal(pc.SignalName, pids.Requests[0].Name)
	s.Equal(uint32(2015), pids.Requests[0].Header)
	s.Equal(uint32(1), pids.Requests[0].Mode)
	s.Equal(uint32(166), pids.Requests[0].Pid)
	s.Equal(pc.Formula, pids.Requests[0].Formula)
	s.Equal(pc.IntervalSeconds, int(pids.Requests[0].IntervalSeconds))
	s.Equal(pc.Protocol, pids.Requests[0].Protocol)
	s.Equal(template.Version, pids.Version)
}

func (s *DeviceConfigControllerTestSuite) TestGetDeviceSettingsByName() {

	const name = "default-ice-settings"
	exampleSettingsJSON := []byte(`{
        "safety_cut_out_voltage": 12.5,
        "sleep_timer_event_driven_period_secs": 30,
        "wake_trigger_voltage_level": 3.3
    }`)

	settingsJSON := null.JSONFrom(exampleSettingsJSON)
	ds := models.DeviceSetting{
		Name:       name,
		Settings:   settingsJSON,
		Powertrain: "ICE",
	}

	err := ds.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	s.app.Get("/device-config/:name/settings", s.controller.GetDeviceSettingsByName)

	request := dbtest.BuildRequest("GET", "/device-config/"+name+"/settings", "")
	response, _ := s.app.Test(request)

	assert.Equal(s.T(), fiber.StatusOK, response.StatusCode)

	var receivedDS grpc.DeviceSetting
	body, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &receivedDS)
	assert.NoError(s.T(), err)

	expectedSettings := &grpc.DeviceSetting{
		SafetyCutOutVoltage:             12.5,
		SleepTimerEventDrivenPeriodSecs: 30,
		WakeTriggerVoltageLevel:         3.3,
	}

	assert.Equal(s.T(), expectedSettings.SafetyCutOutVoltage, receivedDS.SafetyCutOutVoltage)
	assert.Equal(s.T(), expectedSettings.SleepTimerEventDrivenPeriodSecs, receivedDS.SleepTimerEventDrivenPeriodSecs)
	assert.Equal(s.T(), expectedSettings.WakeTriggerVoltageLevel, receivedDS.WakeTriggerVoltageLevel)
}

func (s *DeviceConfigControllerTestSuite) TestGetDBCFileByTemplateName() {

	template := models.Template{
		TemplateName: "exampleDBCFileTemplate",
		Version:      "3.0",
		Protocol:     models.CanProtocolTypeCAN11_500,
		Powertrain:   "ICE",
	}
	err := template.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	dbcf := models.DBCFile{
		TemplateName: "exampleDBCFileTemplate",
		DBCFile:      "ThisIsTheDBCFileContent",
	}

	err = dbcf.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	s.app.Get("/device-config/:templateName/dbc-file", s.controller.GetDBCFileByTemplateName)

	request := dbtest.BuildRequest("GET", "/device-config/"+template.TemplateName+"/dbc-file", "")
	response, err := s.app.Test(request)
	require.NoError(s.T(), err)

	require.Equal(s.T(), fiber.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	assert.Equal(s.T(), dbcf.DBCFile, string(body))

	templateFromDB, err := models.Templates(models.TemplateWhere.TemplateName.EQ(template.TemplateName)).One(s.ctx, s.pdb.DBS().Reader.DB)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), template.Version, templateFromDB.Version)
}

func (s *DeviceConfigControllerTestSuite) TestGetConfigURLs_EmptyDBC() {

	vin := "TMBEK6NW1N3088739"

	mockedUserDevice := &pb.UserDevice{
		Id:                  ksuid.New().String(),
		UserId:              ksuid.New().String(),
		Vin:                 &vin,
		DeviceDefinitionId:  ksuid.New().String(),
		VinConfirmed:        true,
		CountryCode:         "USA",
		PowerTrainType:      "HEV",
		CANProtocol:         "7",
		PostalCode:          "48025",
		GeoDecodedCountry:   "USA",
		GeoDecodedStateProv: "MI",
	}
	s.mockUserDeviceSvc.EXPECT().GetUserDeviceByVIN(gomock.Any(), vin).Return(mockedUserDevice, nil)

	mockedDeviceDefinition := &p_grpc.GetDeviceDefinitionItemResponse{
		DeviceDefinitionId: ksuid.New().String(),
		Type: &p_grpc.DeviceType{
			Year: 2020,
		},
	}
	s.mockDeviceDefSvc.EXPECT().GetDeviceDefinitionByID(gomock.Any(), gomock.Any()).Return(mockedDeviceDefinition, nil)

	template := &models.Template{
		TemplateName: "some-template-emptydbc",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN29_500,
		Powertrain:   "HEV",
	}
	err := template.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	templateVehicle := &models.TemplateVehicle{
		TemplateName: template.TemplateName,
		YearStart:    2010,
		YearEnd:      2025,
	}
	err = templateVehicle.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	ds := &models.DeviceSetting{
		Name:         "default-hev-emptydbc",
		Powertrain:   "HEV",
		TemplateName: null.NewString(template.TemplateName, true),
	}
	err = ds.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	s.app.Get("/config-urls/:vin", s.controller.GetConfigURLsFromVIN)

	request := dbtest.BuildRequest("GET", "/config-urls/"+vin, "")
	response, err := s.app.Test(request)
	require.NoError(s.T(), err)

	body, _ := io.ReadAll(response.Body)
	require.Equal(s.T(), fiber.StatusOK, response.StatusCode)

	var receivedResp DeviceConfigResponse
	err = json.Unmarshal(body, &receivedResp)
	require.NoError(s.T(), err)

	assert.Equal(s.T(), fmt.Sprintf("http://localhost:3000/v1/device-config/%s/pids", template.TemplateName), receivedResp.PidURL)
	assert.Equal(s.T(), fmt.Sprintf("http://localhost:3000/v1/device-config/%s/device-settings", ds.Name), receivedResp.DeviceSettingURL)
	assert.Empty(s.T(), receivedResp.DbcURL)
	assert.Equal(s.T(), template.Version, receivedResp.Version)
}

func (s *DeviceConfigControllerTestSuite) TestGetConfigURLs_DecodeVIN() {

	vin := "TMBEK6NW1N3088739"

	template := &models.Template{
		TemplateName: "some-template",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN11_500,
		Powertrain:   "HEV",
	}
	err := template.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	mockedDeviceDefinition := &p_grpc.GetDeviceDefinitionItemResponse{
		DeviceDefinitionId: ksuid.New().String(),
		Type: &p_grpc.DeviceType{
			Year:      2020,
			MakeSlug:  "Ford",
			ModelSlug: "Mustang",
		},
		DeviceAttributes: []*p_grpc.DeviceTypeAttribute{{
			Name:  "powertrain_type",
			Value: "HEV",
		}},
	}

	ds := &models.DeviceSetting{
		Name:         "default-hev",
		Powertrain:   "HEV",
		TemplateName: null.NewString(template.TemplateName, true),
	}
	err = ds.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	s.mockUserDeviceSvc.EXPECT().GetUserDeviceByVIN(gomock.Any(), vin).Return(nil, errors.New("user device not found"))

	s.mockDeviceDefSvc.EXPECT().DecodeVIN(gomock.Any(), vin).Return(&p_grpc.DecodeVinResponse{
		DeviceDefinitionId: mockedDeviceDefinition.DeviceDefinitionId,
	}, nil)
	s.mockDeviceDefSvc.EXPECT().GetDeviceDefinitionByID(gomock.Any(), mockedDeviceDefinition.DeviceDefinitionId).Return(mockedDeviceDefinition, nil)

	s.app.Get("/config-urls/:vin", s.controller.GetConfigURLsFromVIN)

	request := dbtest.BuildRequest("GET", "/config-urls/"+vin, "")
	response, err := s.app.Test(request)
	require.NoError(s.T(), err)

	body, _ := io.ReadAll(response.Body)
	require.Equal(s.T(), fiber.StatusOK, response.StatusCode)

	var receivedResp DeviceConfigResponse
	err = json.Unmarshal(body, &receivedResp)
	require.NoError(s.T(), err)

	assert.Equal(s.T(), fmt.Sprintf("http://localhost:3000/v1/device-config/%s/pids", template.TemplateName), receivedResp.PidURL)
	assert.Equal(s.T(), fmt.Sprintf("http://localhost:3000/v1/device-config/%s/device-settings", ds.Name), receivedResp.DeviceSettingURL)
	assert.Empty(s.T(), receivedResp.DbcURL)
	assert.Equal(s.T(), template.Version, receivedResp.Version)
}

func (s *DeviceConfigControllerTestSuite) TestGetConfigURLs_ProtocolOverrideQS() {

	vin := "TMBEK6NW1N3088739"

	decoy := &models.Template{
		TemplateName: "not-wanted-template",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN11_500,
		Powertrain:   "HEV",
	}
	err := decoy.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	template := &models.Template{
		TemplateName: "some-template-protocol-override",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN29_500,
		Powertrain:   "HEV",
	}
	err = template.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	ds := &models.DeviceSetting{
		Name:         "default-hev-protocol-override",
		Powertrain:   "HEV",
		TemplateName: null.NewString(template.TemplateName, template.TemplateName != ""),
	}
	err = ds.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	mockedDeviceDefinition := &p_grpc.GetDeviceDefinitionItemResponse{
		DeviceDefinitionId: ksuid.New().String(),
		Type: &p_grpc.DeviceType{
			Year:      2020,
			MakeSlug:  "Ford",
			ModelSlug: "Mustang",
		},
		DeviceAttributes: []*p_grpc.DeviceTypeAttribute{{
			Name:  "powertrain_type",
			Value: "HEV",
		}},
	}
	s.mockDeviceDefSvc.EXPECT().GetDeviceDefinitionByID(gomock.Any(), gomock.Any()).Return(mockedDeviceDefinition, nil)
	s.mockUserDeviceSvc.EXPECT().GetUserDeviceByVIN(gomock.Any(), vin).Return(nil, errors.New("user device not found"))
	s.mockDeviceDefSvc.EXPECT().DecodeVIN(gomock.Any(), vin).Return(&p_grpc.DecodeVinResponse{DeviceDefinitionId: mockedDeviceDefinition.DeviceDefinitionId}, nil)

	s.app.Get("/config-urls/:vin", s.controller.GetConfigURLsFromVIN)

	request := dbtest.BuildRequest("GET", "/config-urls/"+vin+"?protocol=7", "")
	response, err := s.app.Test(request, -1)
	require.NoError(s.T(), err)

	body, _ := io.ReadAll(response.Body)

	assert.Equal(s.T(), fiber.StatusOK, response.StatusCode, "response body: "+string(body))

	var receivedResp DeviceConfigResponse
	err = json.Unmarshal(body, &receivedResp)
	assert.NoError(s.T(), err)

	assert.Equal(s.T(), fmt.Sprintf("http://localhost:3000/v1/device-config/%s/pids", template.TemplateName), receivedResp.PidURL)
	assert.Equal(s.T(), fmt.Sprintf("http://localhost:3000/v1/device-config/%s/device-settings", ds.Name), receivedResp.DeviceSettingURL)
	assert.Equal(s.T(), template.Version, receivedResp.Version)

}

func (s *DeviceConfigControllerTestSuite) TestGetConfigURLs_FallbackLogic() {

	vin := "TMBEK6NW1N3088739"

	parentTemplate := &models.Template{
		TemplateName: "parent-template",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN11_500,
		Powertrain:   "BEV",
	}
	err := parentTemplate.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	parentDS := &models.DeviceSetting{
		Name:         "parent-settings-fallback",
		Powertrain:   "BEV",
		TemplateName: null.NewString(parentTemplate.TemplateName, true),
	}
	err = parentDS.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	decoyDS := &models.DeviceSetting{
		Name:       "decoy-device-settings",
		Powertrain: "BEV",
	}
	err = decoyDS.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	// matched template without device settings but has a parent template
	matchedTemplate := &models.Template{
		TemplateName:       "matched-template",
		Version:            "1.0",
		Protocol:           models.CanProtocolTypeCAN29_500,
		Powertrain:         "BEV",
		ParentTemplateName: null.NewString(parentTemplate.TemplateName, true),
	}
	err = matchedTemplate.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	mockedDeviceDefinition := &p_grpc.GetDeviceDefinitionItemResponse{
		DeviceDefinitionId: ksuid.New().String(),
		Type: &p_grpc.DeviceType{
			Year:      2020,
			MakeSlug:  "Ford",
			ModelSlug: "Mustang",
		},
		DeviceAttributes: []*p_grpc.DeviceTypeAttribute{{
			Name:  "powertrain_type",
			Value: "BEV",
		}},
	}
	s.mockDeviceDefSvc.EXPECT().GetDeviceDefinitionByID(gomock.Any(), gomock.Any()).Return(mockedDeviceDefinition, nil)
	s.mockUserDeviceSvc.EXPECT().GetUserDeviceByVIN(gomock.Any(), vin).Return(nil, errors.New("user device not found"))
	s.mockDeviceDefSvc.EXPECT().DecodeVIN(gomock.Any(), vin).Return(&p_grpc.DecodeVinResponse{DeviceDefinitionId: mockedDeviceDefinition.DeviceDefinitionId}, nil)

	s.app.Get("/config-urls/:vin", s.controller.GetConfigURLsFromVIN)

	request := dbtest.BuildRequest("GET", "/config-urls/"+vin+"?protocol=7", "")
	response, err := s.app.Test(request)
	require.NoError(s.T(), err)

	body, _ := io.ReadAll(response.Body)

	var receivedResp DeviceConfigResponse
	err = json.Unmarshal(body, &receivedResp)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), fmt.Sprintf("http://localhost:3000/v1/device-config/%s/device-settings", parentDS.Name), receivedResp.DeviceSettingURL)
	assert.Equal(s.T(), matchedTemplate.Version, receivedResp.Version)
}

func (s *DeviceConfigControllerTestSuite) TestRetrieveAndSetVehicleInfo() {

	ud := &pb.UserDevice{
		DeviceDefinitionId: "some-definition-id",
	}
	expectedDDResponse := &p_grpc.GetDeviceDefinitionItemResponse{
		Type: &p_grpc.DeviceType{
			Year:      2021,
			MakeSlug:  "Ford",
			ModelSlug: "Mustang",
		},
	}

	s.mockDeviceDefSvc.EXPECT().
		GetDeviceDefinitionByID(gomock.Any(), ud.DeviceDefinitionId).
		Return(expectedDDResponse, nil)

	vehicleMake, vehicleModel, vehicleYear, err := s.controller.retrieveAndSetVehicleInfo(s.ctx, ud)

	require.NoError(s.T(), err)
	assert.Equal(s.T(), "Ford", vehicleMake)
	assert.Equal(s.T(), "Mustang", vehicleModel)
	assert.Equal(s.T(), 2021, vehicleYear)
}

func (s *DeviceConfigControllerTestSuite) TestSetPowerTrainType() {

	testCases := []struct {
		name          string
		deviceAttrs   []*p_grpc.DeviceTypeAttribute
		expectedPower string
	}{
		{
			name: "With Specified Powertrain",
			deviceAttrs: []*p_grpc.DeviceTypeAttribute{
				{Name: "powertrain_type", Value: "Electric"},
			},
			expectedPower: "Electric",
		},
		{
			name:          "Without Specified Powertrain",
			deviceAttrs:   []*p_grpc.DeviceTypeAttribute{},
			expectedPower: "ICE", // Default value
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ddResponse := &p_grpc.GetDeviceDefinitionItemResponse{
				DeviceAttributes: tc.deviceAttrs,
			}
			ud := &pb.UserDevice{}
			s.controller.setPowerTrainType(ddResponse, ud)
			assert.Equal(s.T(), tc.expectedPower, ud.PowerTrainType)
		})
	}
}

func (s *DeviceConfigControllerTestSuite) TestSetCANProtocol() {

	testCases := []struct {
		name        string
		initialCAN  string
		expectedCAN string
	}{
		{
			name:        "CAN Protocol 6",
			initialCAN:  "6",
			expectedCAN: models.CanProtocolTypeCAN11_500,
		},
		{
			name:        "CAN Protocol 7",
			initialCAN:  "7",
			expectedCAN: models.CanProtocolTypeCAN29_500,
		},
		{
			name:        "Empty CAN Protocol",
			initialCAN:  "",
			expectedCAN: models.CanProtocolTypeCAN11_500,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ud := &pb.UserDevice{CANProtocol: tc.initialCAN}
			s.controller.setCANProtocol(ud)
			assert.Equal(s.T(), tc.expectedCAN, ud.CANProtocol)
		})
	}
}

func (s *DeviceConfigControllerTestSuite) TestSelectAndFetchTemplate_DeviceDefinitions() {

	template := &models.Template{
		TemplateName: "some-template",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN29_500,
		Powertrain:   "HEV",
	}
	err := template.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	deviceDef := &models.TemplateDeviceDefinition{
		DeviceDefinitionID: "device-def-id",
		TemplateName:       template.TemplateName,
	}
	err = deviceDef.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	mockedUserDevice := &pb.UserDevice{
		Id:                 ksuid.New().String(),
		UserId:             ksuid.New().String(),
		DeviceDefinitionId: ksuid.New().String(),
		CANProtocol:        models.CanProtocolTypeCAN29_500,
		PowerTrainType:     "HEV",
	}

	vehicleMake := "Ford"
	vehicleModel := "Mustang"
	vehicleYear := 2021

	fetchedTemplate, err := s.controller.selectAndFetchTemplate(s.ctx, mockedUserDevice, vehicleMake, vehicleModel, vehicleYear)

	require.NoError(s.T(), err)
	assert.NotNil(s.T(), fetchedTemplate)
	assert.Equal(s.T(), template.TemplateName, fetchedTemplate.TemplateName)
}

func (s *DeviceConfigControllerTestSuite) TestSelectAndFetchTemplate_MMY() {

	decoy := &models.Template{
		TemplateName: "mmy-template-decoy",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN29_500,
		Powertrain:   "HEV",
	}
	err := decoy.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	template := &models.Template{
		TemplateName: "mmy-template",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN29_500,
		Powertrain:   "HEV",
	}
	err = template.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	templateVehicle := &models.TemplateVehicle{
		TemplateName:   template.TemplateName,
		MakeSlug:       null.StringFrom("Ford"),
		ModelWhitelist: types.StringArray{"Mustang"},
		YearStart:      2010,
		YearEnd:        2025,
	}
	err = templateVehicle.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	mockedUserDevice := &pb.UserDevice{
		Id:                 ksuid.New().String(),
		UserId:             ksuid.New().String(),
		DeviceDefinitionId: "non-existing-def-id",
		CANProtocol:        models.CanProtocolTypeCAN29_500,
		PowerTrainType:     "HEV",
	}

	vehicleMake := "Ford"
	vehicleModel := "Mustang"
	vehicleYear := 2021

	fetchedTemplate, err := s.controller.selectAndFetchTemplate(s.ctx, mockedUserDevice, vehicleMake, vehicleModel, vehicleYear)

	require.NoError(s.T(), err)
	assert.NotNil(s.T(), fetchedTemplate)
	assert.Equal(s.T(), template.TemplateName, fetchedTemplate.TemplateName)
}

func (s *DeviceConfigControllerTestSuite) TestSelectAndFetchTemplate_ModelWhitelistMatch() {

	decoyTemplate := &models.Template{
		TemplateName: "decoy-template",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN11_500,
		Powertrain:   "ICE",
	}
	err := decoyTemplate.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	template := &models.Template{
		TemplateName: "template",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN11_500,
		Powertrain:   "ICE",
	}
	err = template.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	// template vehicles with different model whitelists
	decoyTemplateVehicle := &models.TemplateVehicle{
		TemplateName:   decoyTemplate.TemplateName,
		MakeSlug:       null.StringFrom("Ford"),
		ModelWhitelist: types.StringArray{"Fiesta"},
		YearStart:      2010,
		YearEnd:        2025,
	}
	err = decoyTemplateVehicle.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	templateVehicle := &models.TemplateVehicle{
		TemplateName:   template.TemplateName,
		MakeSlug:       null.StringFrom("Ford"),
		ModelWhitelist: types.StringArray{"Mustang"},
		YearStart:      2010,
		YearEnd:        2025,
	}
	err = templateVehicle.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	mockedUserDevice := &pb.UserDevice{
		Id:                 ksuid.New().String(),
		UserId:             ksuid.New().String(),
		DeviceDefinitionId: "non-existing-def-id",
		CANProtocol:        models.CanProtocolTypeCAN11_500,
		PowerTrainType:     "ICE",
	}

	vehicleMake := "Ford"
	vehicleModel := "Mustang"
	vehicleYear := 2021

	fetchedTemplate, err := s.controller.selectAndFetchTemplate(s.ctx, mockedUserDevice, vehicleMake, vehicleModel, vehicleYear)

	require.NoError(s.T(), err)
	assert.NotNil(s.T(), fetchedTemplate)
	assert.Equal(s.T(), template.TemplateName, fetchedTemplate.TemplateName)
}

func (s *DeviceConfigControllerTestSuite) TestSelectAndFetchTemplate_YearRange() {

	template2 := &models.Template{
		TemplateName: "default-template",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN11_500,
		Powertrain:   "ICE",
	}
	err := template2.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	template := &models.Template{
		TemplateName: "2019plus-template",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN11_500,
		Powertrain:   "ICE",
	}
	err = template.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	templateVehicle := &models.TemplateVehicle{
		TemplateName: template.TemplateName,
		YearStart:    2019,
		YearEnd:      2025,
	}
	err = templateVehicle.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	mockedUserDevice := &pb.UserDevice{
		Id:                 ksuid.New().String(),
		UserId:             ksuid.New().String(),
		DeviceDefinitionId: "some-2019-vehicle",
		PowerTrainType:     template.Powertrain,
		CANProtocol:        template.Protocol,
	}

	vehicleMake := "Ford"
	vehicleModel := "Mustang"
	vehicleYear := 2019

	fetchedTemplate, err := s.controller.selectAndFetchTemplate(s.ctx, mockedUserDevice, vehicleMake, vehicleModel, vehicleYear)

	require.NoError(s.T(), err)
	assert.NotNil(s.T(), fetchedTemplate)
	assert.Equal(s.T(), template.TemplateName, fetchedTemplate.TemplateName)
}

func (s *DeviceConfigControllerTestSuite) TestSelectAndFetchTemplate_PowertrainProtocol() {

	decoy := &models.Template{
		TemplateName: "protocol-powertrain-template-decoy",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN29_500,
		Powertrain:   "PHEV",
	}
	err := decoy.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	template := &models.Template{
		TemplateName: "protocol-powertrain-template",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN29_500,
		Powertrain:   "HEV",
	}
	err = template.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	mockedUserDevice := &pb.UserDevice{
		Id:                 ksuid.New().String(),
		UserId:             ksuid.New().String(),
		DeviceDefinitionId: "non-existing-def-id",
		CANProtocol:        models.CanProtocolTypeCAN29_500,
		PowerTrainType:     "HEV",
	}

	vehicleMake := "Ford"
	vehicleModel := "Mustang"
	vehicleYear := 2021

	fetchedTemplate, err := s.controller.selectAndFetchTemplate(s.ctx, mockedUserDevice, vehicleMake, vehicleModel, vehicleYear)

	require.NoError(s.T(), err)
	assert.NotNil(s.T(), fetchedTemplate)
	assert.Equal(s.T(), template.TemplateName, fetchedTemplate.TemplateName)
}

func (s *DeviceConfigControllerTestSuite) TestSelectAndFetchTemplate_Default() {

	nonDefaultTmpl := &models.Template{
		TemplateName: "some-template-special",
		Version:      "1.0",
		Protocol:     "CAN11_500",
		Powertrain:   "ICE",
	}
	err := nonDefaultTmpl.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	defaultTemplate := &models.Template{
		TemplateName: "default-some-template",
		Version:      "1.0",
		Protocol:     "CAN11_500",
		Powertrain:   "ICE",
	}
	err = defaultTemplate.Insert(s.ctx, s.pdb.DBS().Writer, boil.Infer())
	require.NoError(s.T(), err)

	mockedUserDevice := &pb.UserDevice{
		Id:                 ksuid.New().String(),
		UserId:             ksuid.New().String(),
		DeviceDefinitionId: "non-existing-def-id",
		CANProtocol:        models.CanProtocolTypeCAN29_500,
		PowerTrainType:     "HEV",
	}

	vehicleMake := "NonExistingMake"
	vehicleModel := "NonExistingModel"
	vehicleYear := 2010

	fetchedTemplate, err := s.controller.selectAndFetchTemplate(s.ctx, mockedUserDevice, vehicleMake, vehicleModel, vehicleYear)

	require.NoError(s.T(), err)
	assert.NotNil(s.T(), fetchedTemplate)
	assert.Equal(s.T(), defaultTemplate.TemplateName, fetchedTemplate.TemplateName)
}

func Test_modelMatch(t *testing.T) {
	tests := []struct {
		name      string
		modelList types.StringArray
		modelName string
		want      bool
	}{
		{
			name:      "match found",
			modelList: types.StringArray{"falcon", "model-x"},
			modelName: "model-x",
			want:      true,
		},
		{
			name:      "no match found",
			modelList: types.StringArray{"falcon", "model-x"},
			modelName: "model-y",
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, modelMatch(tt.modelList, tt.modelName), "modelMatch(%v, %v)", tt.modelList, tt.modelName)
		})
	}
}

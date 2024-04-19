package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	mock_gateways "github.com/DIMO-Network/vehicle-signal-decoding/internal/gateways/mocks"

	common2 "github.com/ethereum/go-ethereum/common"

	"google.golang.org/protobuf/types/known/timestamppb"

	_ "github.com/lib/pq"

	gdata "github.com/DIMO-Network/device-data-api/pkg/grpc"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/appmodels"

	"github.com/DIMO-Network/shared/db"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"

	"github.com/volatiletech/null/v8"

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
	ctx                   context.Context
	pdb                   db.Store
	container             testcontainers.Container
	mockCtrl              *gomock.Controller
	logger                *zerolog.Logger
	mockUserDeviceSvc     *mock_services.MockUserDeviceService
	mockDeviceDefSvc      *mock_services.MockDeviceDefinitionsService
	mockDeviceTemplateSvc *mock_services.MockDeviceTemplateService
	controller            *DeviceConfigController
	app                   *fiber.App
	mockIdentityAPI       *mock_gateways.MockIdentityAPI
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
	s.mockDeviceTemplateSvc = mock_services.NewMockDeviceTemplateService(s.mockCtrl)
	s.mockIdentityAPI = mock_gateways.NewMockIdentityAPI(s.mockCtrl)
	ctrl := NewDeviceConfigController(&config.Settings{Port: "3000", DeploymentURL: "http://localhost:3000"}, s.logger,
		s.pdb.DBS().Reader.DB, s.mockUserDeviceSvc, s.mockDeviceDefSvc, s.mockDeviceTemplateSvc, s.mockIdentityAPI)
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

func (s *DeviceConfigControllerTestSuite) TestGetConfigURLsFromVIN_EmptyDBC() {
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

	s.mockDeviceTemplateSvc.EXPECT().ResolveDeviceConfiguration(gomock.Any(), mockedUserDevice, nil, common2.HexToAddress("")).Return(&appmodels.DeviceConfigResponse{
		PidURL:           "http://localhost:3000/v1/device-config/pids/some-template-emptydbc@v1.0.0",
		DeviceSettingURL: "http://localhost:3000/v1/device-config/settings/default-hev-emptydbc@v1.0.0",
	}, nil)

	s.app.Get("/config-urls/:vin", s.controller.GetConfigURLsFromVIN)

	request := dbtest.BuildRequest("GET", "/config-urls/"+vin, "")
	response, err := s.app.Test(request)
	require.NoError(s.T(), err)

	body, _ := io.ReadAll(response.Body)
	require.Equal(s.T(), fiber.StatusOK, response.StatusCode)

	var receivedResp appmodels.DeviceConfigResponse
	err = json.Unmarshal(body, &receivedResp)
	require.NoError(s.T(), err)

	assert.Equal(s.T(), fmt.Sprintf("http://localhost:3000/v1/device-config/pids/%s@v1.0.0", template.TemplateName), receivedResp.PidURL)
	assert.Equal(s.T(), fmt.Sprintf("http://localhost:3000/v1/device-config/settings/%s@v1.0.0", ds.Name), receivedResp.DeviceSettingURL)
	assert.Empty(s.T(), receivedResp.DbcURL)
}

func (s *DeviceConfigControllerTestSuite) TestGetConfigURLsFromVIN_DecodeVIN() {
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

	s.mockDeviceTemplateSvc.EXPECT().ResolveDeviceConfiguration(gomock.Any(), &pb.UserDevice{
		Vin:                &vin,
		DeviceDefinitionId: mockedDeviceDefinition.DeviceDefinitionId,
		//PowerTrainType:     "HEV",
	}, nil, common2.HexToAddress("")).Return(&appmodels.DeviceConfigResponse{
		PidURL:           "http://localhost:3000/v1/device-config/pids/some-template@v1.0.0",
		DeviceSettingURL: "http://localhost:3000/v1/device-config/settings/default-hev@v1.0.0",
	}, nil)
	s.app.Get("/config-urls/:vin", s.controller.GetConfigURLsFromVIN)

	request := dbtest.BuildRequest("GET", "/config-urls/"+vin, "")
	response, err := s.app.Test(request)
	require.NoError(s.T(), err)

	body, _ := io.ReadAll(response.Body)
	require.Equal(s.T(), fiber.StatusOK, response.StatusCode)

	var receivedResp appmodels.DeviceConfigResponse
	err = json.Unmarshal(body, &receivedResp)
	require.NoError(s.T(), err)

	assert.Equal(s.T(), fmt.Sprintf("http://localhost:3000/v1/device-config/pids/%s@v1.0.0", template.TemplateName), receivedResp.PidURL)
	assert.Equal(s.T(), fmt.Sprintf("http://localhost:3000/v1/device-config/settings/%s@v1.0.0", ds.Name), receivedResp.DeviceSettingURL)
	assert.Empty(s.T(), receivedResp.DbcURL)
}

func (s *DeviceConfigControllerTestSuite) TestGetConfigURLsFromVIN_ProtocolOverrideQS() {

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
	s.mockUserDeviceSvc.EXPECT().GetUserDeviceByVIN(gomock.Any(), vin).Return(nil, errors.New("user device not found"))
	s.mockDeviceDefSvc.EXPECT().DecodeVIN(gomock.Any(), vin).Return(&p_grpc.DecodeVinResponse{DeviceDefinitionId: mockedDeviceDefinition.DeviceDefinitionId}, nil)

	s.mockDeviceTemplateSvc.EXPECT().ResolveDeviceConfiguration(gomock.Any(), &pb.UserDevice{
		Vin:                &vin,
		DeviceDefinitionId: mockedDeviceDefinition.DeviceDefinitionId,
		CANProtocol:        "7",
	}, nil, common2.HexToAddress("")).Return(&appmodels.DeviceConfigResponse{
		PidURL:           "http://localhost:3000/v1/device-config/pids/some-template-protocol-override@v1.0.0",
		DeviceSettingURL: "http://localhost:3000/v1/device-config/settings/default-hev-protocol-override@v1.0.0",
	}, nil)

	s.app.Get("/config-urls/:vin", s.controller.GetConfigURLsFromVIN)

	request := dbtest.BuildRequest("GET", "/config-urls/"+vin+"?protocol=7", "")
	response, err := s.app.Test(request, -1)
	require.NoError(s.T(), err)

	body, _ := io.ReadAll(response.Body)

	assert.Equal(s.T(), fiber.StatusOK, response.StatusCode, "response body: "+string(body))

	var receivedResp appmodels.DeviceConfigResponse
	err = json.Unmarshal(body, &receivedResp)
	assert.NoError(s.T(), err)

	assert.Equal(s.T(), fmt.Sprintf("http://localhost:3000/v1/device-config/pids/%s@v1.0.0", template.TemplateName), receivedResp.PidURL)
	assert.Equal(s.T(), fmt.Sprintf("http://localhost:3000/v1/device-config/settings/%s@v1.0.0", ds.Name), receivedResp.DeviceSettingURL)

}

func (s *DeviceConfigControllerTestSuite) TestGetConfigURLsFromVIN_FallbackLogic() {
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
	s.mockUserDeviceSvc.EXPECT().GetUserDeviceByVIN(gomock.Any(), vin).Return(nil, errors.New("user device not found"))
	s.mockDeviceDefSvc.EXPECT().DecodeVIN(gomock.Any(), vin).Return(&p_grpc.DecodeVinResponse{DeviceDefinitionId: mockedDeviceDefinition.DeviceDefinitionId}, nil)

	s.mockDeviceTemplateSvc.EXPECT().ResolveDeviceConfiguration(gomock.Any(), &pb.UserDevice{
		Vin:                &vin,
		DeviceDefinitionId: mockedDeviceDefinition.DeviceDefinitionId,
		CANProtocol:        "7",
	}, nil, common2.HexToAddress("")).Return(&appmodels.DeviceConfigResponse{
		PidURL:           "http://localhost:3000/v1/device-config/pids/parent-template@v1.0.0",
		DeviceSettingURL: "http://localhost:3000/v1/device-config/settings/parent-settings-fallback@v1.0.0",
	}, nil)

	s.app.Get("/config-urls/:vin", s.controller.GetConfigURLsFromVIN)

	request := dbtest.BuildRequest("GET", "/config-urls/"+vin+"?protocol=7", "")
	response, err := s.app.Test(request)
	require.NoError(s.T(), err)

	body, _ := io.ReadAll(response.Body)

	var receivedResp appmodels.DeviceConfigResponse
	err = json.Unmarshal(body, &receivedResp)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), fmt.Sprintf("http://localhost:3000/v1/device-config/settings/%s@v1.0.0", parentDS.Name), receivedResp.DeviceSettingURL)
}

func (s *DeviceConfigControllerTestSuite) TestGetConfigStatusByEthAddr_DeviceDataOnly() {
	ethAddr := "0x29e8Ec52A3d2c9b72aA9F0e3e2576F3A28480299"
	s.app.Get("/device-config/eth-addr/:ethAddr/status", s.controller.GetConfigStatusByEthAddr)
	vin := "TMBEK6NW1N3088739"
	s.controller.fwVersionAPI = mockHTTPClientFwVersion{}

	testUD := &pb.UserDevice{
		Id:                 ksuid.New().String(),
		UserId:             ksuid.New().String(),
		Vin:                &vin,
		DeviceDefinitionId: ksuid.New().String(),
		VinConfirmed:       true,
		CountryCode:        "USA",
		PowerTrainType:     "ICE",
		CANProtocol:        "6",
	}
	s.mockUserDeviceSvc.EXPECT().GetUserDeviceByEthAddr(gomock.Any(), common2.HexToAddress(ethAddr)).Return(testUD, nil)

	s.mockUserDeviceSvc.EXPECT().GetRawDeviceData(gomock.Any(), testUD.Id).Return(&gdata.RawDeviceDataResponse{Items: []*gdata.RawDeviceDataResponseItem{
		{
			IntegrationId:   ksuid.New().String(),
			UserDeviceId:    testUD.Id,
			SignalsJsonData: []byte(`{"fwVersion": { "value": "v0.8.5"} }`),
			RecordUpdatedAt: timestamppb.New(time.Now()),
			RecordCreatedAt: timestamppb.New(time.Now()),
		},
	}}, nil)

	request := dbtest.BuildRequest("GET", "/device-config/eth-addr/"+ethAddr+"/status", "")
	response, err := s.app.Test(request)
	require.NoError(s.T(), err)

	body, _ := io.ReadAll(response.Body)
	var receivedResp DeviceTemplateStatusResponse
	err = json.Unmarshal(body, &receivedResp)
	assert.NoError(s.T(), err)

	assert.Equal(s.T(), false, receivedResp.IsTemplateUpToDate)
	assert.Equal(s.T(), true, receivedResp.IsFirmwareUpToDate)
	assert.Equal(s.T(), "v0.8.5", receivedResp.FirmwareVersion)
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

func Test_parseOutFWVersion(t *testing.T) {
	type args struct {
		data *gdata.RawDeviceDataResponse
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "get version",
			args: args{data: &gdata.RawDeviceDataResponse{Items: []*gdata.RawDeviceDataResponseItem{{
				SignalsJsonData: []byte(`{
"fwVersion": {
    "value": "0.8.5",
    "source": "dimo/integration/2ULfuC8U9dOqRshZBAi0lMM1Rrx",
    "timestamp": "2024-01-02T11:17:20Z"
  }			
}`),
			}}}},
			want: "v0.8.5",
		},
		{
			name: "empty version",
			args: args{data: &gdata.RawDeviceDataResponse{Items: []*gdata.RawDeviceDataResponseItem{{
				SignalsJsonData: []byte(`{}`),
			},
			}}},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, parseOutFWVersion(tt.args.data), "parseOutFWVersion(%v)", tt.args.data)
		})
	}
}

func Test_parseOutTemplateAndVersion(t *testing.T) {
	type args struct {
		templateNameWithVersion string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			name:  "name with version",
			args:  args{templateNameWithVersion: "default-ice@v1.0.0"},
			want:  "default-ice",
			want1: "v1.0.0",
		},
		{
			name:  "name without version",
			args:  args{templateNameWithVersion: "default-ice"},
			want:  "default-ice",
			want1: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseOutTemplateAndVersion(tt.args.templateNameWithVersion)
			assert.Equalf(t, tt.want, got, "parseOutTemplateAndVersion(%v)", tt.args.templateNameWithVersion)
			assert.Equalf(t, tt.want1, got1, "parseOutTemplateAndVersion(%v)", tt.args.templateNameWithVersion)
		})
	}
}

func Test_isFwUpToDate(t *testing.T) {
	type args struct {
		latest  string
		current string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "blank",
			args: args{
				latest:  "v0.8.5",
				current: "v",
			},
			want: false,
		},
		{
			name: "up to date",
			args: args{
				latest:  "v0.8.5",
				current: "v0.8.5",
			},
			want: true,
		},
		{
			name: "not up to date",
			args: args{
				latest:  "v0.8.5",
				current: "v0.8.4",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isFwUpToDate(tt.args.latest, tt.args.current), "isFwUpToDate(%v, %v)", tt.args.latest, tt.args.current)
		})
	}
}

// used for test
type mockHTTPClientFwVersion struct {
}

func (m mockHTTPClientFwVersion) ExecuteRequest(_, _ string, _ []byte) (*http.Response, error) {
	buf := bytes.NewBufferString(`{"name": "v0.8.5"}`)
	requestBody := io.NopCloser(buf)

	mockResponse := http.Response{
		Status:           "OK",
		StatusCode:       200,
		Body:             requestBody,
		Header:           make(http.Header),
		ContentLength:    0,
		TransferEncoding: nil,
		Close:            false,
	}
	mockResponse.Header.Set("Content-Type", "application/json")

	return &mockResponse, nil
}

package services

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/gateways"

	p_grpc "github.com/DIMO-Network/device-definitions-api/pkg/grpc"
	pb "github.com/DIMO-Network/devices-api/pkg/grpc"
	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	mock_services "github.com/DIMO-Network/vehicle-signal-decoding/internal/core/services/mocks"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/dbtest"
	"github.com/rs/zerolog"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/types"
	"go.uber.org/mock/gomock"
)

const migrationsDirRelPath = "../../infrastructure/db/migrations"

type DeviceTemplateServiceTestSuite struct {
	suite.Suite
	ctx              context.Context
	pdb              db.Store
	container        testcontainers.Container
	mockCtrl         *gomock.Controller
	logger           *zerolog.Logger
	mockDeviceDefSvc *mock_services.MockDeviceDefinitionsService
	// subject under test
	sut *deviceTemplateService
}

const dbSchemaName = "vehicle_signal_decoding"

func (s *DeviceTemplateServiceTestSuite) SetupSuite() {
	s.ctx = context.Background()
	s.pdb, s.container = dbtest.StartContainerDatabase(s.ctx, dbSchemaName, s.T(), migrationsDirRelPath)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	s.logger = &logger
	s.mockCtrl = gomock.NewController(s.T())
	s.mockDeviceDefSvc = mock_services.NewMockDeviceDefinitionsService(s.mockCtrl)

	s.sut = &deviceTemplateService{
		db:           s.pdb.DBS().Writer.DB,
		log:          *s.logger,
		settings:     &config.Settings{},
		deviceDefSvc: s.mockDeviceDefSvc,
	}
}

func (s *DeviceTemplateServiceTestSuite) SetupTest() {
	dbtest.TruncateTables(s.pdb.DBS().Writer.DB, dbSchemaName, s.T())
}

func (s *DeviceTemplateServiceTestSuite) TearDownSuite() {
	if err := s.container.Terminate(s.ctx); err != nil {
		s.T().Fatal(err)
	}
	s.mockCtrl.Finish()
}

func (s *DeviceTemplateServiceTestSuite) TearDownTest() {
	dbtest.TruncateTables(s.pdb.DBS().Writer.DB, dbSchemaName, s.T())
	temps, err := models.Templates().All(context.Background(), s.pdb.DBS().Writer)
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, temp := range temps {
		fmt.Printf("template: %+v\n", temp)
	}
}

func TestDeviceTemplateServiceTestSuite(t *testing.T) {
	suite.Run(t, new(DeviceTemplateServiceTestSuite))
}

func (s *DeviceTemplateServiceTestSuite) TestRetrievePowertrain() {
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

	dts := &deviceTemplateService{
		db:           s.pdb.DBS().Writer.DB,
		log:          *s.logger,
		settings:     &config.Settings{},
		deviceDefSvc: s.mockDeviceDefSvc,
	}

	powertrain, err := dts.retrievePowertrain(s.ctx, ud.DeviceDefinitionId)

	require.NoError(s.T(), err)
	assert.Equal(s.T(), "ICE", powertrain)
}

func (s *DeviceTemplateServiceTestSuite) TestConvertCANProtocol() {
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
			cp := convertCANProtocol(s.sut.log, ud.CANProtocol)
			assert.Equal(s.T(), tc.expectedCAN, cp)
		})
	}
}

func (s *DeviceTemplateServiceTestSuite) Test_selectAndFetchTemplate_DeviceDefinitions() {

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
	vehicle := &gateways.VehicleInfo{
		TokenID: 123,
		Definition: gateways.VehicleDefinition{
			Make:  "Ford",
			Model: "Mustang",
			Year:  2021,
		},
	}

	fetchedTemplate, err := s.sut.selectAndFetchTemplate(s.ctx, mockedUserDevice.CANProtocol, mockedUserDevice.PowerTrainType,
		mockedUserDevice.DeviceDefinitionId, vehicle)

	require.NoError(s.T(), err)
	assert.NotNil(s.T(), fetchedTemplate)
	assert.Equal(s.T(), template.TemplateName, fetchedTemplate.TemplateName)
}

func (s *DeviceTemplateServiceTestSuite) Test_selectAndFetchTemplate_nilVehicle() {
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
	s.mockDeviceDefSvc.EXPECT().GetDeviceDefinitionByID(gomock.Any(), mockedUserDevice.DeviceDefinitionId).Return(&p_grpc.GetDeviceDefinitionItemResponse{
		DeviceDefinitionId: mockedUserDevice.DeviceDefinitionId,
		Type: &p_grpc.DeviceType{
			Type:  "vehicle",
			Make:  "Ford",
			Model: "Mustang",
			Year:  2021,
		},
	}, nil)

	fetchedTemplate, err := s.sut.selectAndFetchTemplate(s.ctx, mockedUserDevice.CANProtocol, mockedUserDevice.PowerTrainType,
		mockedUserDevice.DeviceDefinitionId, nil)

	require.NoError(s.T(), err)
	assert.NotNil(s.T(), fetchedTemplate)
	assert.Equal(s.T(), template.TemplateName, fetchedTemplate.TemplateName)
}

func (s *DeviceTemplateServiceTestSuite) Test_selectAndFetchTemplate_MMY() {

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
	vehicle := &gateways.VehicleInfo{
		TokenID: 123,
		Definition: gateways.VehicleDefinition{
			Make:  "Ford",
			Model: "Mustang",
			Year:  2021,
		},
	}

	fetchedTemplate, err := s.sut.selectAndFetchTemplate(s.ctx, mockedUserDevice.CANProtocol, mockedUserDevice.PowerTrainType,
		mockedUserDevice.DeviceDefinitionId, vehicle)

	require.NoError(s.T(), err)
	assert.NotNil(s.T(), fetchedTemplate)
	assert.Equal(s.T(), template.TemplateName, fetchedTemplate.TemplateName)
}

func (s *DeviceTemplateServiceTestSuite) Test_selectAndFetchTemplate_ModelWhitelistMatch() {

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

	vehicle := &gateways.VehicleInfo{
		TokenID: 123,
		Definition: gateways.VehicleDefinition{
			Make:  "Ford",
			Model: "Mustang",
			Year:  2021,
		},
	}

	fetchedTemplate, err := s.sut.selectAndFetchTemplate(s.ctx, mockedUserDevice.CANProtocol, mockedUserDevice.PowerTrainType,
		mockedUserDevice.DeviceDefinitionId, vehicle)

	require.NoError(s.T(), err)
	assert.NotNil(s.T(), fetchedTemplate)
	assert.Equal(s.T(), template.TemplateName, fetchedTemplate.TemplateName)
}

func (s *DeviceTemplateServiceTestSuite) Test_selectAndFetchTemplate_YearRange() {

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
	vehicle := &gateways.VehicleInfo{
		TokenID: 123,
		Definition: gateways.VehicleDefinition{
			Make:  "Ford",
			Model: "Mustang",
			Year:  2020,
		},
	}

	fetchedTemplate, err := s.sut.selectAndFetchTemplate(s.ctx, mockedUserDevice.CANProtocol, mockedUserDevice.PowerTrainType,
		mockedUserDevice.DeviceDefinitionId, vehicle)

	require.NoError(s.T(), err)
	assert.NotNil(s.T(), fetchedTemplate)
	assert.Equal(s.T(), template.TemplateName, fetchedTemplate.TemplateName)
}

func (s *DeviceTemplateServiceTestSuite) Test_selectAndFetchTemplate_PowertrainProtocol() {

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

	vehicle := &gateways.VehicleInfo{
		TokenID: 123,
		Definition: gateways.VehicleDefinition{
			Make:  "Ford",
			Model: "Mustang",
			Year:  2021,
		},
	}

	fetchedTemplate, err := s.sut.selectAndFetchTemplate(s.ctx, mockedUserDevice.CANProtocol, mockedUserDevice.PowerTrainType,
		mockedUserDevice.DeviceDefinitionId, vehicle)

	require.NoError(s.T(), err)
	assert.NotNil(s.T(), fetchedTemplate)
	assert.Equal(s.T(), template.TemplateName, fetchedTemplate.TemplateName)
}

func (s *DeviceTemplateServiceTestSuite) Test_selectAndFetchTemplate_Default() {

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
	vehicle := &gateways.VehicleInfo{
		TokenID: 123,
		Definition: gateways.VehicleDefinition{
			Make:  "NonExistingMake",
			Model: "NonExistingModel",
			Year:  2010,
		},
	}

	fetchedTemplate, err := s.sut.selectAndFetchTemplate(s.ctx, mockedUserDevice.CANProtocol, mockedUserDevice.PowerTrainType,
		mockedUserDevice.DeviceDefinitionId, vehicle)

	require.NoError(s.T(), err)
	assert.NotNil(s.T(), fetchedTemplate)
	assert.Equal(s.T(), defaultTemplate.TemplateName, fetchedTemplate.TemplateName)
}

func Test_deviceTemplateService_buildConfigRoute(t *testing.T) {
	type fields struct {
		settings *config.Settings
	}
	type args struct {
		ct      configType
		name    string
		version string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "pids url",
			fields: fields{settings: &config.Settings{DeploymentURL: "https://vehicle-signal-decoding.dimo.zone"}},
			args:   args{name: "default-ice-can11", ct: PIDs, version: "v1.0.0"},
			want:   "https://vehicle-signal-decoding.dimo.zone/v1/device-config/pids/default-ice-can11@v1.0.0",
		},
		{
			name:   "dbc url",
			fields: fields{settings: &config.Settings{DeploymentURL: "https://vehicle-signal-decoding.dimo.zone"}},
			args:   args{name: "default-ice-can11", ct: DBC, version: "v2.0.0"},
			want:   "https://vehicle-signal-decoding.dimo.zone/v1/device-config/dbc/default-ice-can11@v2.0.0",
		},
		{
			name:   "settings url",
			fields: fields{settings: &config.Settings{DeploymentURL: "https://vehicle-signal-decoding.dimo.zone"}},
			args:   args{name: "default-ice", ct: Setting, version: "v1.0.0"},
			want:   "https://vehicle-signal-decoding.dimo.zone/v1/device-config/settings/default-ice@v1.0.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dts := &deviceTemplateService{
				db:           nil,
				log:          zerolog.Logger{},
				settings:     tt.fields.settings,
				deviceDefSvc: nil,
			}
			assert.Equalf(t, tt.want, dts.buildConfigRoute(tt.args.ct, tt.args.name, tt.args.version), "buildConfigRoute(%v, %v, %v)", tt.args.ct, tt.args.name, tt.args.version)
		})
	}
}

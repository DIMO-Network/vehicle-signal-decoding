package services

import (
	"context"
	"fmt"
	"os"
	"testing"

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

const migrationsDirRelPath = "../infrastructure/db/migrations"

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

func (s *DeviceTemplateServiceTestSuite) TestRetrieveAndSetVehicleInfo() {
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

	vehicleMake, vehicleModel, vehicleYear, err := dts.retrieveAndSetVehicleInfo(s.ctx, ud)

	require.NoError(s.T(), err)
	assert.Equal(s.T(), "Ford", vehicleMake)
	assert.Equal(s.T(), "Mustang", vehicleModel)
	assert.Equal(s.T(), 2021, vehicleYear)
}

func (s *DeviceTemplateServiceTestSuite) TestSetPowerTrainType() {

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
			setPowerTrainType(ddResponse, ud)
			assert.Equal(s.T(), tc.expectedPower, ud.PowerTrainType)
		})
	}
}

func (s *DeviceTemplateServiceTestSuite) TestSetCANProtocol() {

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
			s.sut.setCANProtocol(ud)
			assert.Equal(s.T(), tc.expectedCAN, ud.CANProtocol)
		})
	}
}

func (s *DeviceTemplateServiceTestSuite) TestSelectAndFetchTemplate_DeviceDefinitions() {

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

	fetchedTemplate, err := s.sut.selectAndFetchTemplate(s.ctx, mockedUserDevice, vehicleMake, vehicleModel, vehicleYear)

	require.NoError(s.T(), err)
	assert.NotNil(s.T(), fetchedTemplate)
	assert.Equal(s.T(), template.TemplateName, fetchedTemplate.TemplateName)
}

func (s *DeviceTemplateServiceTestSuite) TestSelectAndFetchTemplate_MMY() {

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

	fetchedTemplate, err := s.sut.selectAndFetchTemplate(s.ctx, mockedUserDevice, vehicleMake, vehicleModel, vehicleYear)

	require.NoError(s.T(), err)
	assert.NotNil(s.T(), fetchedTemplate)
	assert.Equal(s.T(), template.TemplateName, fetchedTemplate.TemplateName)
}

func (s *DeviceTemplateServiceTestSuite) TestSelectAndFetchTemplate_ModelWhitelistMatch() {

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

	fetchedTemplate, err := s.sut.selectAndFetchTemplate(s.ctx, mockedUserDevice, vehicleMake, vehicleModel, vehicleYear)

	require.NoError(s.T(), err)
	assert.NotNil(s.T(), fetchedTemplate)
	assert.Equal(s.T(), template.TemplateName, fetchedTemplate.TemplateName)
}

func (s *DeviceTemplateServiceTestSuite) TestSelectAndFetchTemplate_YearRange() {

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

	fetchedTemplate, err := s.sut.selectAndFetchTemplate(s.ctx, mockedUserDevice, vehicleMake, vehicleModel, vehicleYear)

	require.NoError(s.T(), err)
	assert.NotNil(s.T(), fetchedTemplate)
	assert.Equal(s.T(), template.TemplateName, fetchedTemplate.TemplateName)
}

func (s *DeviceTemplateServiceTestSuite) TestSelectAndFetchTemplate_PowertrainProtocol() {

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

	fetchedTemplate, err := s.sut.selectAndFetchTemplate(s.ctx, mockedUserDevice, vehicleMake, vehicleModel, vehicleYear)

	require.NoError(s.T(), err)
	assert.NotNil(s.T(), fetchedTemplate)
	assert.Equal(s.T(), template.TemplateName, fetchedTemplate.TemplateName)
}

func (s *DeviceTemplateServiceTestSuite) TestSelectAndFetchTemplate_Default() {

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

	fetchedTemplate, err := s.sut.selectAndFetchTemplate(s.ctx, mockedUserDevice, vehicleMake, vehicleModel, vehicleYear)

	require.NoError(s.T(), err)
	assert.NotNil(s.T(), fetchedTemplate)
	assert.Equal(s.T(), defaultTemplate.TemplateName, fetchedTemplate.TemplateName)
}

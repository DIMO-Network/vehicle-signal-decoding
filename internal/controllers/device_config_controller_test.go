package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"

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
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/test"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq" //nolint
)

const migrationsDirRelPath = "../infrastructure/db/migrations"

func TestGetPIDsByTemplate(t *testing.T) {

	// arrange global db and route setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Arrange: db and route setup
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("app", "vehicle-signal-decoding").
		Logger()

	ctx := context.Background()

	//Spin up test database in a Docker container
	pdb, container := test.StartContainerDatabase(ctx, t, migrationsDirRelPath)
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatal(err)
		}
	}()

	template := models.Template{
		TemplateName: "exampleTemplate",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN11_500,
		Powertrain:   "ICE",
	}
	err := template.Insert(context.Background(), pdb.DBS().Writer, boil.Infer())
	require.NoError(t, err)

	pc := models.PidConfig{
		ID:              1,
		SignalName:      "odometer",
		TemplateName:    "exampleTemplate",
		Header:          []byte{0x07, 0xdf}, // short notation without padding
		Mode:            []byte{0x01},
		Pid:             []byte{0xa6},
		Formula:         "A*5",
		IntervalSeconds: 60,
		Protocol:        models.CanProtocolTypeCAN11_500,
	}

	err = pc.Insert(context.Background(), pdb.DBS().Writer, boil.Infer())
	require.NoError(t, err)

	mockUserDeviceSvc := mock_services.NewMockUserDeviceService(mockCtrl)
	mockDeviceDefSvc := mock_services.NewMockDeviceDefinitionsService(mockCtrl)
	c := NewDeviceConfigController(&config.Settings{Port: "3000"}, &logger, pdb.DBS().Reader.DB, mockUserDeviceSvc, mockDeviceDefSvc)

	app := fiber.New()
	app.Get("/device-config/:templateName/pids", c.GetPIDsByTemplate)

	t.Run("GET - PIDs by Template", func(t *testing.T) {

		// Act: make the request
		request := test.BuildRequest("GET", "/device-config/"+template.TemplateName+"/pids", "")
		response, _ := app.Test(request)
		body, _ := io.ReadAll(response.Body)

		// Assert: check the results
		if assert.Equal(t, fiber.StatusOK, response.StatusCode) == false {
			fmt.Println("response body: " + string(body))
		}
		fmt.Printf("Received response: %s", body)

		pids := grpc.PIDRequests{}
		err = json.Unmarshal(body, &pids)
		assert.NoError(t, err)

		require.Equal(t, 1, len(pids.Requests))
		assert.Equal(t, pc.SignalName, pids.Requests[0].Name)
		// use known values uint32 that we expect above bytes to convert to as uint32 decimal
		assert.Equal(t, uint32(2015), pids.Requests[0].Header)
		assert.Equal(t, uint32(1), pids.Requests[0].Mode)
		assert.Equal(t, uint32(166), pids.Requests[0].Pid)

		assert.Equal(t, pc.Formula, pids.Requests[0].Formula)
		assert.Equal(t, pc.IntervalSeconds, int(pids.Requests[0].IntervalSeconds))
		assert.Equal(t, pc.Protocol, pids.Requests[0].Protocol)
		assert.Equal(t, template.Version, pids.Version)

		// Teardown: cleanup after test
		test.TruncateTables(pdb.DBS().Writer.DB, t)

	})
}

func TestGetDeviceSettingsByTemplate(t *testing.T) {

	// arrange global db and route setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Arrange: db and route setup
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("app", "vehicle-signal-decoding").
		Logger()

	ctx := context.Background()

	// Spin up test database in a Docker container
	pdb, container := test.StartContainerDatabase(ctx, t, migrationsDirRelPath)
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatal(err)
		}
	}()

	template := models.Template{
		TemplateName: "testTemplate",
		Version:      "2.0",
		Protocol:     models.CanProtocolTypeCAN11_500,
		Powertrain:   "ICE",
	}
	err := template.Insert(ctx, pdb.DBS().Writer, boil.Infer())
	assert.NoError(t, err)

	ds := models.DeviceSetting{
		ID:                            1,
		TemplateName:                  "testTemplate",
		BatteryCriticalLevelVoltage:   3.2,
		SafetyCutOutVoltage:           2.8,
		SleepTimerEventDrivenInterval: 5,
		//etc
	}

	err = ds.Insert(ctx, pdb.DBS().Writer, boil.Infer())
	assert.NoError(t, err)

	mockUserDeviceSvc := mock_services.NewMockUserDeviceService(mockCtrl)
	mockDeviceDefSvc := mock_services.NewMockDeviceDefinitionsService(mockCtrl)
	c := NewDeviceConfigController(&config.Settings{Port: "3000"}, &logger, pdb.DBS().Reader.DB, mockUserDeviceSvc, mockDeviceDefSvc)

	app := fiber.New()
	app.Get("/device-config/:templateName", c.GetDeviceSettingsByTemplate)

	t.Run("GET - Device Settings by Template", func(t *testing.T) {

		// Act: make the request
		request := test.BuildRequest("GET", "/device-config/"+template.TemplateName, "")
		response, _ := app.Test(request)
		body, _ := io.ReadAll(response.Body)

		// Assert: check the results
		if assert.Equal(t, fiber.StatusOK, response.StatusCode) == false {
			fmt.Println("response body: " + string(body))
		}

		var receivedDS grpc.DeviceSetting
		err = json.Unmarshal(body, &receivedDS)
		assert.NoError(t, err)

		assert.Equal(t, float32(ds.BatteryCriticalLevelVoltage), receivedDS.BatteryCriticalLevelVoltage)
		assert.Equal(t, float32(ds.SafetyCutOutVoltage), receivedDS.SafetyCutOutVoltage)
		assert.Equal(t, float32(ds.SleepTimerEventDrivenInterval), receivedDS.SleepTimerEventDrivenIntervalSecs)

		// Testing Version
		templateFromDB, err := models.Templates(models.TemplateWhere.TemplateName.EQ(template.TemplateName)).One(context.Background(), pdb.DBS().Reader.DB)
		assert.NoError(t, err)
		assert.Equal(t, template.Version, templateFromDB.Version)

		// Teardown: cleanup after test
		test.TruncateTables(pdb.DBS().Writer.DB, t)

	})
}
func TestGetDBCFileByTemplateName(t *testing.T) {

	// arrange global db and route setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Arrange: db and route setup
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("app", "vehicle-signal-decoding").
		Logger()

	ctx := context.Background()

	// Spin up test database in a Docker container
	pdb, container := test.StartContainerDatabase(ctx, t, migrationsDirRelPath)
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatal(err)
		}
	}()

	template := models.Template{
		TemplateName: "exampleDBCFileTemplate",
		Version:      "3.0",
		Protocol:     models.CanProtocolTypeCAN11_500,
		Powertrain:   "ICE",
		// etc
	}
	err := template.Insert(context.Background(), pdb.DBS().Writer, boil.Infer())
	assert.NoError(t, err)

	dbcf := models.DBCFile{
		TemplateName: "exampleDBCFileTemplate",
		DBCFile:      "ThisIsTheDBCFileContent",
	}

	err = dbcf.Insert(context.Background(), pdb.DBS().Writer, boil.Infer())
	assert.NoError(t, err)

	mockUserDeviceSvc := mock_services.NewMockUserDeviceService(mockCtrl)

	mockDeviceDefSvc := mock_services.NewMockDeviceDefinitionsService(mockCtrl)
	c := NewDeviceConfigController(&config.Settings{Port: "3000"}, &logger, pdb.DBS().Reader.DB, mockUserDeviceSvc, mockDeviceDefSvc)

	app := fiber.New()
	app.Get("/device-config/:templateName/dbc-file", c.GetDBCFileByTemplateName)

	t.Run("GET - DBCFile by TemplateName", func(t *testing.T) {
		// Act: make the request
		request := test.BuildRequest("GET", "/device-config/"+template.TemplateName+"/dbc-file", "")
		response, _ := app.Test(request)
		body, _ := io.ReadAll(response.Body)

		// Assert: check the results
		if assert.Equal(t, fiber.StatusOK, response.StatusCode) == false {
			fmt.Println("response body: " + string(body))
		}

		assert.Equal(t, dbcf.DBCFile, string(body))

		// Testing Version
		templateFromDB, err := models.Templates(models.TemplateWhere.TemplateName.EQ(template.TemplateName)).One(context.Background(), pdb.DBS().Reader.DB)
		assert.NoError(t, err)
		assert.Equal(t, template.Version, templateFromDB.Version)

		// Teardown: cleanup after test
		test.TruncateTables(pdb.DBS().Writer.DB, t)
	})
}

func TestGetConfigURLsEmptyDBC(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("app", "vehicle-signal-decoding").
		Logger()

	ctx := context.Background()

	// Spin up test database in a Docker container
	pdb, container := test.StartContainerDatabase(ctx, t, migrationsDirRelPath)
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatal(err)
		}
	}()
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

	mockUserDeviceSvc := mock_services.NewMockUserDeviceService(mockCtrl)
	mockUserDeviceSvc.EXPECT().GetUserDeviceByVIN(gomock.Any(), vin).Return(mockedUserDevice, nil)

	// Mock the device definition service
	mockedDeviceDefinition := &p_grpc.GetDeviceDefinitionItemResponse{
		DeviceDefinitionId: ksuid.New().String(),
		Type: &p_grpc.DeviceType{
			Year: 2020,
		},
	}

	mockDeviceDefSvc := mock_services.NewMockDeviceDefinitionsService(mockCtrl)
	mockDeviceDefSvc.EXPECT().GetDeviceDefinitionByID(gomock.Any(), gomock.Any()).Return(mockedDeviceDefinition, nil)
	c := NewDeviceConfigController(&config.Settings{Port: "3000", DeploymentURL: "http://localhost:3000"}, &logger, pdb.DBS().Reader.DB, mockUserDeviceSvc, mockDeviceDefSvc)

	// insert template in DB
	template := &models.Template{
		TemplateName: "some-template",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN29_500,
		Powertrain:   "HEV",
	}
	err := template.Insert(ctx, pdb.DBS().Writer, boil.Infer())
	require.NoError(t, err)
	// insert device settings
	ds := &models.DeviceSetting{
		TemplateName: "some-template",
	}
	err = ds.Insert(ctx, pdb.DBS().Writer, boil.Infer())
	require.NoError(t, err)

	app := fiber.New()
	app.Get("/config-urls/:vin", c.GetConfigURLsFromVIN)

	t.Run("GET - Config URLs by VIN with Empty DbcURL", func(t *testing.T) {
		// Add vehicle year range to the template
		templateVehicle := &models.TemplateVehicle{
			TemplateName: template.TemplateName,
			YearStart:    2010,
			YearEnd:      2025,
		}
		err2 := templateVehicle.Insert(context.Background(), pdb.DBS().Writer, boil.Infer())
		require.NoError(t, err2)
		// Act: make the request
		request := test.BuildRequest("GET", "/config-urls/"+vin, "")
		response, err := app.Test(request)
		require.NoError(t, err)
		body, _ := io.ReadAll(response.Body)

		// Assert: check the results
		if !assert.Equal(t, fiber.StatusOK, response.StatusCode) {
			fmt.Println("response body: " + string(body))
		}

		var receivedResp DeviceConfigResponse
		err = json.Unmarshal(body, &receivedResp)
		assert.NoError(t, err)

		assert.Equal(t, fmt.Sprintf("http://localhost:3000/v1/device-config/%s/pids", template.TemplateName), receivedResp.PidURL)
		assert.Equal(t, fmt.Sprintf("http://localhost:3000/v1/device-config/%s/device-settings", template.TemplateName), receivedResp.DeviceSettingURL)
		assert.Equal(t, "", receivedResp.DbcURL)

		assert.Equal(t, template.Version, receivedResp.Version)

		// Teardown: cleanup after test
		test.TruncateTables(pdb.DBS().Writer.DB, t)
	})

}

func TestGetConfigURLsEmptyDeviceSettings(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("app", "vehicle-signal-decoding").
		Logger()

	ctx := context.Background()

	// Spin up test database in a Docker container
	pdb, container := test.StartContainerDatabase(ctx, t, migrationsDirRelPath)
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatal(err)
		}
	}()
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

	mockUserDeviceSvc := mock_services.NewMockUserDeviceService(mockCtrl)
	mockUserDeviceSvc.EXPECT().GetUserDeviceByVIN(gomock.Any(), vin).Return(mockedUserDevice, nil)
	// Mock the device definition service
	mockedDeviceDefinition := &p_grpc.GetDeviceDefinitionItemResponse{
		DeviceDefinitionId: ksuid.New().String(),
		Type: &p_grpc.DeviceType{
			Year: 2020,
		},
	}
	mockDeviceDefSvc := mock_services.NewMockDeviceDefinitionsService(mockCtrl)
	mockDeviceDefSvc.EXPECT().GetDeviceDefinitionByID(gomock.Any(), gomock.Any()).Return(mockedDeviceDefinition, nil)
	c := NewDeviceConfigController(&config.Settings{Port: "3000", DeploymentURL: "http://localhost:3000"}, &logger, pdb.DBS().Reader.DB, mockUserDeviceSvc, mockDeviceDefSvc)

	// insert template in DB
	template := &models.Template{
		TemplateName: "some-template",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN29_500,
		Powertrain:   "HEV",
	}
	err := template.Insert(context.Background(), pdb.DBS().Writer, boil.Infer())
	require.NoError(t, err)

	app := fiber.New()
	app.Get("/config-urls/:vin", c.GetConfigURLsFromVIN)

	t.Run("GET - Config URLs by VIN with Matching Year Range", func(t *testing.T) {
		// Insert DBCFile in the database
		dbcFile := &models.DBCFile{
			TemplateName: "some-template",
			DBCFile:      "sample-dbc-file-name",
		}
		err := dbcFile.Insert(context.Background(), pdb.DBS().Writer, boil.Infer())
		require.NoError(t, err)
		// Add vehicle year range to the template
		templateVehicle := &models.TemplateVehicle{
			TemplateName: template.TemplateName,
			YearStart:    2019,
			YearEnd:      2100,
		}
		err2 := templateVehicle.Insert(context.Background(), pdb.DBS().Writer, boil.Infer())
		require.NoError(t, err2)

		// Act: make the request
		request := test.BuildRequest("GET", "/config-urls/"+vin, "")
		response, err := app.Test(request)
		require.NoError(t, err)
		body, _ := io.ReadAll(response.Body)

		// Assert: check the results
		if !assert.Equal(t, fiber.StatusOK, response.StatusCode) {
			fmt.Println("response body: " + string(body))
		}

		var receivedResp DeviceConfigResponse
		err = json.Unmarshal(body, &receivedResp)
		assert.NoError(t, err)

		assert.Equal(t, fmt.Sprintf("http://localhost:3000/v1/device-config/%s/pids", template.TemplateName), receivedResp.PidURL)
		assert.Equal(t, "", receivedResp.DeviceSettingURL)
		assert.Equal(t, fmt.Sprintf("http://localhost:3000/v1/device-config/%s/dbc", template.TemplateName), receivedResp.DbcURL)

		assert.Equal(t, template.Version, receivedResp.Version)

		test.TruncateTables(pdb.DBS().Writer.DB, t)
	})

}
func TestGetConfigURLsDecodeVIN(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("app", "vehicle-signal-decoding").
		Logger()

	ctx := context.Background()

	pdb, container := test.StartContainerDatabase(ctx, t, migrationsDirRelPath)
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatal(err)
		}
	}()
	vin := "TMBEK6NW1N3088739"

	// Insert template "some-template" into the database
	template := &models.Template{
		TemplateName: "some-template",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN11_500,
		Powertrain:   "HEV",
	}
	err := template.Insert(ctx, pdb.DBS().Writer, boil.Infer())
	require.NoError(t, err)

	// Insert device settings for "some-template"
	ds := &models.DeviceSetting{
		TemplateName: template.TemplateName,
	}
	err = ds.Insert(ctx, pdb.DBS().Writer, boil.Infer())
	require.NoError(t, err)

	// Mock the device definition service
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

	mockUserDeviceSvc := mock_services.NewMockUserDeviceService(mockCtrl)
	mockUserDeviceSvc.EXPECT().GetUserDeviceByVIN(gomock.Any(), vin).Return(nil, errors.New("user device not found"))

	mockDeviceDefSvc := mock_services.NewMockDeviceDefinitionsService(mockCtrl)
	mockDeviceDefSvc.EXPECT().DecodeVIN(gomock.Any(), vin).Return(&p_grpc.DecodeVinResponse{
		DeviceDefinitionId: mockedDeviceDefinition.DeviceDefinitionId,
	}, nil)
	mockDeviceDefSvc.EXPECT().GetDeviceDefinitionByID(gomock.Any(), mockedDeviceDefinition.DeviceDefinitionId).Return(mockedDeviceDefinition, nil)

	c := NewDeviceConfigController(&config.Settings{Port: "3000", DeploymentURL: "http://localhost:3000"}, &logger, pdb.DBS().Reader.DB, mockUserDeviceSvc, mockDeviceDefSvc)

	app := fiber.New()
	app.Get("/config-urls/:vin", c.GetConfigURLsFromVIN)

	// Act
	request := test.BuildRequest("GET", "/config-urls/"+vin, "")
	response, err := app.Test(request)
	require.NoError(t, err)

	body, _ := io.ReadAll(response.Body)

	// Assert
	assert.Equal(t, fiber.StatusOK, response.StatusCode, "response body: "+string(body))

	var receivedResp DeviceConfigResponse
	err = json.Unmarshal(body, &receivedResp)
	assert.NoError(t, err)

	//"some-template"
	assert.Equal(t, fmt.Sprintf("http://localhost:3000/v1/device-config/%s/pids", template.TemplateName), receivedResp.PidURL)
	assert.Equal(t, fmt.Sprintf("http://localhost:3000/v1/device-config/%s/device-settings", template.TemplateName), receivedResp.DeviceSettingURL)
	assert.Equal(t, "", receivedResp.DbcURL)
	assert.Equal(t, template.Version, receivedResp.Version)

	test.TruncateTables(pdb.DBS().Writer.DB, t)
}

func TestRetrieveAndSetVehicleInfo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDeviceDefSvc := mock_services.NewMockDeviceDefinitionsService(mockCtrl)
	mockUserDeviceSvc := mock_services.NewMockUserDeviceService(mockCtrl)

	logger := zerolog.New(os.Stdout)

	pdb, container := test.StartContainerDatabase(context.Background(), t, migrationsDirRelPath)
	defer func() {
		if err := container.Terminate(context.Background()); err != nil {
			t.Fatal(err)
		}
	}()

	c := NewDeviceConfigController(
		&config.Settings{Port: "3000"},
		&logger,
		pdb.DBS().Reader.DB,
		mockUserDeviceSvc,
		mockDeviceDefSvc,
	)

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

	mockDeviceDefSvc.EXPECT().
		GetDeviceDefinitionByID(gomock.Any(), ud.DeviceDefinitionId).
		Return(expectedDDResponse, nil)

	vehicleMake, vehicleModel, vehicleYear, err := c.retrieveAndSetVehicleInfo(context.Background(), ud)

	// Assert:
	require.NoError(t, err)
	assert.Equal(t, "Ford", vehicleMake)
	assert.Equal(t, "Mustang", vehicleModel)
	assert.Equal(t, 2021, vehicleYear)

}
func TestSetPowerTrainType(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDeviceDefSvc := mock_services.NewMockDeviceDefinitionsService(mockCtrl)
	mockUserDeviceSvc := mock_services.NewMockUserDeviceService(mockCtrl)

	logger := zerolog.New(os.Stdout)

	pdb, container := test.StartContainerDatabase(context.Background(), t, migrationsDirRelPath)
	defer func() {
		if err := container.Terminate(context.Background()); err != nil {
			t.Fatal(err)
		}
	}()

	c := NewDeviceConfigController(
		&config.Settings{Port: "3000"},
		&logger,
		pdb.DBS().Reader.DB,
		mockUserDeviceSvc,
		mockDeviceDefSvc,
	)

	// Define test cases
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
		t.Run(tc.name, func(t *testing.T) {
			// Create a mock GetDeviceDefinitionItemResponse
			ddResponse := &p_grpc.GetDeviceDefinitionItemResponse{
				DeviceAttributes: tc.deviceAttrs,
			}

			ud := &pb.UserDevice{}

			// Act:
			c.setPowerTrainType(ddResponse, ud)

			// Assert:
			assert.Equal(t, tc.expectedPower, ud.PowerTrainType)
		})
	}
}

func TestSetCANProtocol(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDeviceDefSvc := mock_services.NewMockDeviceDefinitionsService(mockCtrl)
	mockUserDeviceSvc := mock_services.NewMockUserDeviceService(mockCtrl)

	logger := zerolog.New(os.Stdout)

	pdb, container := test.StartContainerDatabase(context.Background(), t, migrationsDirRelPath)
	defer func() {
		if err := container.Terminate(context.Background()); err != nil {
			t.Fatal(err)
		}
	}()

	c := NewDeviceConfigController(
		&config.Settings{Port: "3000"},
		&logger,
		pdb.DBS().Reader.DB,
		mockUserDeviceSvc,
		mockDeviceDefSvc,
	)

	// Define test cases
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
		t.Run(tc.name, func(t *testing.T) {
			ud := &pb.UserDevice{CANProtocol: tc.initialCAN}

			// Act
			c.setCANProtocol(ud)

			// Assert
			assert.Equal(t, tc.expectedCAN, ud.CANProtocol)
		})
	}
}

func TestSelectAndFetchTemplate_DeviceDefinitions(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("app", "vehicle-signal-decoding").
		Logger()

	ctx := context.Background()

	pdb, container := test.StartContainerDatabase(ctx, t, migrationsDirRelPath)
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatal(err)
		}
	}()

	mockUserDeviceSvc := mock_services.NewMockUserDeviceService(mockCtrl)
	mockDeviceDefSvc := mock_services.NewMockDeviceDefinitionsService(mockCtrl)
	c := NewDeviceConfigController(&config.Settings{Port: "3000"}, &logger, pdb.DBS().Reader.DB, mockUserDeviceSvc, mockDeviceDefSvc)

	// Insert template into the database
	template := &models.Template{
		TemplateName: "some-template",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN29_500,
		Powertrain:   "HEV",
	}
	err := template.Insert(ctx, pdb.DBS().Writer, boil.Infer())
	require.NoError(t, err)

	// Insert a matching template device definition
	deviceDef := &models.TemplateDeviceDefinition{
		DeviceDefinitionID: "device-def-id",
		TemplateName:       template.TemplateName,
	}
	err = deviceDef.Insert(ctx, pdb.DBS().Writer, boil.Infer())
	require.NoError(t, err)

	// Create a mocked user device
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

	// Act
	fetchedTemplate, err := c.selectAndFetchTemplate(ctx, mockedUserDevice, vehicleMake, vehicleModel, vehicleYear)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, fetchedTemplate)
	assert.Equal(t, template.TemplateName, fetchedTemplate.TemplateName)

	test.TruncateTables(pdb.DBS().Writer.DB, t)
}

func TestSelectAndFetchTemplate_MMY(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("app", "vehicle-signal-decoding").
		Logger()

	ctx := context.Background()

	pdb, container := test.StartContainerDatabase(ctx, t, migrationsDirRelPath)
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatal(err)
		}
	}()

	mockUserDeviceSvc := mock_services.NewMockUserDeviceService(mockCtrl)
	mockDeviceDefSvc := mock_services.NewMockDeviceDefinitionsService(mockCtrl)
	c := NewDeviceConfigController(&config.Settings{Port: "3000"}, &logger, pdb.DBS().Reader.DB, mockUserDeviceSvc, mockDeviceDefSvc)

	// Insert template into the database
	template := &models.Template{
		TemplateName: "mmy-template",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN29_500,
		Powertrain:   "HEV",
	}
	err := template.Insert(ctx, pdb.DBS().Writer, boil.Infer())
	require.NoError(t, err)

	// Insert a template vehicle that matches the MMY
	templateVehicle := &models.TemplateVehicle{
		TemplateName:   template.TemplateName,
		MakeSlug:       "Ford",
		ModelWhitelist: types.StringArray{"Mustang"},
		YearStart:      2010,
		YearEnd:        2025,
	}
	err = templateVehicle.Insert(ctx, pdb.DBS().Writer, boil.Infer())
	require.NoError(t, err)

	// Create a mocked user device without a matching device definition
	mockedUserDevice := &pb.UserDevice{
		Id:                 ksuid.New().String(),
		UserId:             ksuid.New().String(),
		DeviceDefinitionId: "non-existing-def-id",
	}

	vehicleMake := "Ford"
	vehicleModel := "Mustang"
	vehicleYear := 2021

	// Act
	fetchedTemplate, err := c.selectAndFetchTemplate(ctx, mockedUserDevice, vehicleMake, vehicleModel, vehicleYear)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, fetchedTemplate)
	assert.Equal(t, template.TemplateName, fetchedTemplate.TemplateName)

	test.TruncateTables(pdb.DBS().Writer.DB, t)
}

func TestSelectAndFetchTemplate_YearRange(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("app", "vehicle-signal-decoding").
		Logger()

	ctx := context.Background()

	pdb, container := test.StartContainerDatabase(ctx, t, migrationsDirRelPath)
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatal(err)
		}
	}()

	mockUserDeviceSvc := mock_services.NewMockUserDeviceService(mockCtrl)
	mockDeviceDefSvc := mock_services.NewMockDeviceDefinitionsService(mockCtrl)
	c := NewDeviceConfigController(&config.Settings{Port: "3000"}, &logger, pdb.DBS().Reader.DB, mockUserDeviceSvc, mockDeviceDefSvc)
	// insert another template to have more test data - we should not get this one
	template2 := &models.Template{
		TemplateName: "default-template",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN11_500,
		Powertrain:   "ICE",
	}
	err := template2.Insert(ctx, pdb.DBS().Writer, boil.Infer())
	require.NoError(t, err)

	// Insert template we expect to get
	template := &models.Template{
		TemplateName: "2019plus-template",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN11_500,
		Powertrain:   "ICE",
	}
	err = template.Insert(ctx, pdb.DBS().Writer, boil.Infer())
	require.NoError(t, err)

	// Insert a template vehicle that matches the year range
	templateVehicle := &models.TemplateVehicle{
		TemplateName: template.TemplateName,
		YearStart:    2019,
		YearEnd:      2025,
	}
	err = templateVehicle.Insert(ctx, pdb.DBS().Writer, boil.Infer())
	require.NoError(t, err)

	// Create a mocked user device without a matching device definition
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

	// Act
	fetchedTemplate, err := c.selectAndFetchTemplate(ctx, mockedUserDevice, vehicleMake, vehicleModel, vehicleYear)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, fetchedTemplate)
	assert.Equal(t, template.TemplateName, fetchedTemplate.TemplateName)

	test.TruncateTables(pdb.DBS().Writer.DB, t)
}

func TestSelectAndFetchTemplate_PowertrainProtocol(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("app", "vehicle-signal-decoding").
		Logger()

	ctx := context.Background()

	pdb, container := test.StartContainerDatabase(ctx, t, migrationsDirRelPath)
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatal(err)
		}
	}()

	mockUserDeviceSvc := mock_services.NewMockUserDeviceService(mockCtrl)
	mockDeviceDefSvc := mock_services.NewMockDeviceDefinitionsService(mockCtrl)
	c := NewDeviceConfigController(&config.Settings{Port: "3000"}, &logger, pdb.DBS().Reader.DB, mockUserDeviceSvc, mockDeviceDefSvc)

	// Insert template into the database
	template := &models.Template{
		TemplateName: "protocol-powertrain-template",
		Version:      "1.0",
		Protocol:     "CAN29_500",
		Powertrain:   "HEV",
	}
	err := template.Insert(ctx, pdb.DBS().Writer, boil.Infer())
	require.NoError(t, err)

	// Create a mocked user device
	mockedUserDevice := &pb.UserDevice{
		Id:                 ksuid.New().String(),
		UserId:             ksuid.New().String(),
		DeviceDefinitionId: "non-existing-def-id",
		CANProtocol:        "CAN29_500",
		PowerTrainType:     "HEV",
	}

	vehicleMake := "Ford"
	vehicleModel := "Mustang"
	vehicleYear := 2021

	// Act
	fetchedTemplate, err := c.selectAndFetchTemplate(ctx, mockedUserDevice, vehicleMake, vehicleModel, vehicleYear)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, fetchedTemplate)
	assert.Equal(t, template.TemplateName, fetchedTemplate.TemplateName)

	// Teardown
	test.TruncateTables(pdb.DBS().Writer.DB, t)
}

func TestSelectAndFetchTemplate_Default(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("app", "vehicle-signal-decoding").
		Logger()

	ctx := context.Background()

	pdb, container := test.StartContainerDatabase(ctx, t, migrationsDirRelPath)
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatal(err)
		}
	}()

	mockUserDeviceSvc := mock_services.NewMockUserDeviceService(mockCtrl)
	mockDeviceDefSvc := mock_services.NewMockDeviceDefinitionsService(mockCtrl)
	c := NewDeviceConfigController(&config.Settings{Port: "3000"}, &logger, pdb.DBS().Reader.DB, mockUserDeviceSvc, mockDeviceDefSvc)

	// Insert a default template into the database
	defaultTemplate := &models.Template{
		TemplateName: "default-some-template",
		Version:      "1.0",
		Protocol:     "CAN11_500",
		Powertrain:   "ICE",
	}
	err := defaultTemplate.Insert(ctx, pdb.DBS().Writer, boil.Infer())
	require.NoError(t, err)

	// Create a mocked user device that does not match any existing definitions, vehicles, or powertrain/protocol
	mockedUserDevice := &pb.UserDevice{
		Id:                 ksuid.New().String(),
		UserId:             ksuid.New().String(),
		DeviceDefinitionId: "non-existing-def-id",
		CANProtocol:        "CAN29_500",
		PowerTrainType:     "HEV",
	}

	vehicleMake := "NonExistingMake"
	vehicleModel := "NonExistingModel"
	vehicleYear := 1999

	// Act
	fetchedTemplate, err := c.selectAndFetchTemplate(ctx, mockedUserDevice, vehicleMake, vehicleModel, vehicleYear)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, fetchedTemplate)
	assert.Equal(t, defaultTemplate.TemplateName, fetchedTemplate.TemplateName)

	test.TruncateTables(pdb.DBS().Writer.DB, t)
}

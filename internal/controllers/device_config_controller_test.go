package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"

	p_grpc "github.com/DIMO-Network/device-definitions-api/pkg/grpc"

	pb "github.com/DIMO-Network/devices-api/pkg/grpc"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/golang/mock/gomock"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/require"

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
		Header:          []byte("07E8"),
		Mode:            []byte("0001"),
		Pid:             []byte("0005"),
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
		// convert uint32 back to bytes to compare
		hdr, _ := bytesToUint32(pc.Header)
		assert.Equal(t, hdr, pids.Requests[0].Header)
		mde, _ := bytesToUint32(pc.Mode)
		assert.Equal(t, mde, pids.Requests[0].Mode)
		pid, _ := bytesToUint32(pc.Pid)
		assert.Equal(t, pid, pids.Requests[0].Pid)
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
		BatteryCriticalLevelVoltage:   "3.2V",
		SafetyCutOutVoltage:           "2.8V",
		SleepTimerEventDrivenInterval: "5s",
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

		var receivedDS DeviceSetting
		err = json.Unmarshal(body, &receivedDS)
		assert.NoError(t, err)

		assert.Equal(t, ds.BatteryCriticalLevelVoltage, receivedDS.BatteryCriticalLevelVoltage)
		assert.Equal(t, ds.SafetyCutOutVoltage, receivedDS.SafetyCutOutVoltage)
		assert.Equal(t, ds.SleepTimerEventDrivenInterval, receivedDS.SleepTimerEventDrivenInterval)

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
	mockedDeviceDefinition := &p_grpc.GetDeviceDefinitionResponse{
		DeviceDefinitions: []*p_grpc.GetDeviceDefinitionItemResponse{
			{
				DeviceDefinitionId: ksuid.New().String(),
				Type: &p_grpc.DeviceType{
					Year: 2020,
				},
			},
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

		assert.Equal(t, fmt.Sprintf("http://localhost:3000/device-config/%s/pids", template.TemplateName), receivedResp.PidURL)
		assert.Equal(t, fmt.Sprintf("http://localhost:3000/device-config/%s/deviceSettings", template.TemplateName), receivedResp.DeviceSettingURL)
		assert.Equal(t, "", receivedResp.DbcURL)

		assert.Equal(t, template.Version, receivedResp.Version)

		// Teardown: cleanup after test
		test.TruncateTables(pdb.DBS().Writer.DB, t)
	})

}

func TestGetConfigURLsMatchingYearRange(t *testing.T) {
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
	mockedDeviceDefinition := &p_grpc.GetDeviceDefinitionResponse{
		DeviceDefinitions: []*p_grpc.GetDeviceDefinitionItemResponse{
			{
				DeviceDefinitionId: ksuid.New().String(),
				Type: &p_grpc.DeviceType{
					Year: 2020,
				},
			},
		},
	}
	mockDeviceDefSvc := mock_services.NewMockDeviceDefinitionsService(mockCtrl)
	mockDeviceDefSvc.EXPECT().GetDeviceDefinitionByID(gomock.Any(), gomock.Any()).Return(mockedDeviceDefinition, nil)
	decodeVinResponse := []*p_grpc.DecodeVinResponse{
		{
			DeviceMakeId:       "some_device_make_id",
			DeviceDefinitionId: "some_device_definition_id",
			DeviceStyleId:      "some_device_style_id",
			Year:               2023,
			Source:             "some_source",
		},
	}

	mockDeviceDefSvc.EXPECT().DecodeVIN(gomock.Any(), vin).Return(decodeVinResponse, nil)
	c := NewDeviceConfigController(&config.Settings{Port: "3000", DeploymentURL: "http://localhost:3000"}, &logger, pdb.DBS().Reader.DB, mockUserDeviceSvc, mockDeviceDefSvc)

	// Mocking to simulate an error for GetUserDeviceByVIN
	mockUserDeviceSvc.EXPECT().GetUserDeviceByVIN(gomock.Any(), vin).Return(nil, errors.New("simulated error"))

	expectedPowerTrainType := "some_value"
	deviceDefinitionResp := &p_grpc.GetDeviceDefinitionResponse{
		DeviceDefinitions: []*p_grpc.GetDeviceDefinitionItemResponse{
			{
				DeviceDefinitionId: ksuid.New().String(),
				DeviceAttributes: []*p_grpc.DeviceTypeAttribute{
					{
						Name:  "powertrain_type",
						Value: expectedPowerTrainType,
					},
				},
			},
		},
	}
	// Mocking to return a valid definition response
	mockDeviceDefSvc.EXPECT().GetDeviceDefinitionByID(gomock.Any(), gomock.Any()).Return(deviceDefinitionResp, nil)
	mockDeviceDefSvc.EXPECT().DecodeVIN(gomock.Any(), vin).Return(deviceDefinitionResp, nil)

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

		assert.Equal(t, fmt.Sprintf("http://localhost:3000/device-config/%s/pids", template.TemplateName), receivedResp.PidURL)
		assert.Equal(t, fmt.Sprintf("http://localhost:3000/device-config/%s/deviceSettings", template.TemplateName), receivedResp.DeviceSettingURL)
		assert.Equal(t, fmt.Sprintf("http://localhost:3000/device-config/%s/dbc", template.TemplateName), receivedResp.DbcURL)

		assert.Equal(t, template.Version, receivedResp.Version)

		// Teardown: cleanup after test
		test.TruncateTables(pdb.DBS().Writer.DB, t)
	})

}

func TestGetConfigURLsNonMatchingYearRange(t *testing.T) {
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
	mockedDeviceDefinition := &p_grpc.GetDeviceDefinitionResponse{
		DeviceDefinitions: []*p_grpc.GetDeviceDefinitionItemResponse{
			{
				DeviceDefinitionId: ksuid.New().String(),
				Type: &p_grpc.DeviceType{
					Year: 2019,
				},
			},
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

	t.Run("GET - Config URLs by VIN with Non-Matching Year Range", func(t *testing.T) {
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
			YearStart:    2010,
			YearEnd:      2015,
		}
		err2 := templateVehicle.Insert(context.Background(), pdb.DBS().Writer, boil.Infer())
		require.NoError(t, err2)

		// Act: make the request
		request := test.BuildRequest("GET", "/config-urls/"+vin, "")
		response, err := app.Test(request)
		require.NoError(t, err)
		body, _ := io.ReadAll(response.Body)

		// Assert: check the results
		assert.Equal(t, fiber.StatusOK, response.StatusCode, "response body: "+string(body))

		var receivedResp DeviceConfigResponse
		err = json.Unmarshal(body, &receivedResp)
		assert.NoError(t, err)

		assert.Equal(t, fmt.Sprintf("http://localhost:3000/device-config/%s/pids", template.TemplateName), receivedResp.PidURL)
		assert.Equal(t, fmt.Sprintf("http://localhost:3000/device-config/%s/deviceSettings", template.TemplateName), receivedResp.DeviceSettingURL)
		assert.Equal(t, fmt.Sprintf("http://localhost:3000/device-config/%s/dbc", template.TemplateName), receivedResp.DbcURL)

		assert.Equal(t, template.Version, receivedResp.Version)

		// Teardown: cleanup after test
		test.TruncateTables(pdb.DBS().Writer.DB, t)
	})

}

package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	mock_services "github.com/DIMO-Network/vehicle-signal-decoding/internal/core/services/mocks"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/test"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
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

	c := NewDeviceConfigController(&config.Settings{Port: "3000"}, &logger, pdb.DBS().Reader.DB, mockUserDeviceSvc)
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

	c := NewDeviceConfigController(&config.Settings{Port: "3000"}, &logger, pdb.DBS().Reader.DB, mockUserDeviceSvc)
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

		assert.Equal(t, ds.ID, receivedDS.ID)

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

	c := NewDeviceConfigController(&config.Settings{Port: "3000"}, &logger, pdb.DBS().Reader.DB, mockUserDeviceSvc)
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

package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/test"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

const migrationsDirRelPath = "../infrastructure/db/migrations"

func TestGetPIDsByTemplate(t *testing.T) {

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
		// etc
	}
	err := template.Insert(context.Background(), pdb.DBS().Writer, boil.Infer())
	assert.NoError(t, err)

	pc := models.PidConfig{
		ID:              1,
		TemplateName:    "exampleTemplate",
		Header:          []byte("7E8"),
		Mode:            []byte("01"),
		Pid:             []byte("05"),
		Formula:         "A*5",
		IntervalSeconds: 60,
		Version:         "1.0",
	}

	errr := pc.Insert(context.Background(), pdb.DBS().Writer, boil.Infer())
	assert.NoError(t, errr)

	c := NewDeviceConfigController(&config.Settings{Port: "3000"}, &logger, pdb.DBS().Reader.DB)
	app := fiber.New()
	app.Get("/device-config/:template_name/pids", c.GetPIDsByTemplate)

	t.Run("GET - PID by Template", func(t *testing.T) {

		// Act: make the request
		request := test.BuildRequest("GET", "/device-config/"+template.TemplateName+"/pids", "")
		response, _ := app.Test(request)
		body, _ := io.ReadAll(response.Body)

		// Assert: check the results
		if assert.Equal(t, fiber.StatusOK, response.StatusCode) == false {
			fmt.Println("response body: " + string(body))
		}

		pids := make([]PIDConfig, 0)
		err = json.Unmarshal(body, &pids)
		assert.NoError(t, err)

		fmt.Printf("Received PIDs: %v\n", pids)

		require.Equal(t, 1, len(pids))
		assert.Equal(t, pc.ID, pids[0].ID)
		assert.Equal(t, pc.Header, pids[0].Header)
		assert.Equal(t, pc.Mode, pids[0].Mode)
		assert.Equal(t, pc.Pid, pids[0].Pid)
		assert.Equal(t, pc.Formula, pids[0].Formula)
		assert.Equal(t, pc.IntervalSeconds, pids[0].IntervalSeconds)
		assert.Equal(t, pc.Version, pids[0].Version)

		// Teardown: cleanup after test
		test.TruncateTables(pdb.DBS().Writer.DB, t)

	})
}

func TestGetDeviceSettingsByTemplate(t *testing.T) {
	// Arrange: db and route setup
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("app", "vehicle-signal-decoding").
		Logger()

	ctx := context.Background()

	// Start test database in a Docker container
	pdb, container := test.StartContainerDatabase(ctx, t, migrationsDirRelPath)
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatal(err)
		}
	}()

	template := models.Template{
		TemplateName: "exampleDeviceSettingsTemplate",
		// etc
	}
	err := template.Insert(context.Background(), pdb.DBS().Writer, boil.Infer())
	assert.NoError(t, err)

	pc := models.PowerConfig{
		ID:           1,
		TemplateName: "exampleDeviceSettings",
		Version:      "1.0",
		//etc
	}

	err = pc.Insert(context.Background(), pdb.DBS().Writer, boil.Infer())
	assert.NoError(t, err)

	c := NewDeviceConfigController(&config.Settings{Port: "3000"}, &logger, pdb.DBS().Reader.DB)
	app := fiber.New()
	app.Get("/device-config/:template_name/deviceSettings", c.GetDeviceSettingsByTemplate)

	t.Run("GET - Device Settings by Template", func(t *testing.T) {
		// Act: make the request
		request := test.BuildRequest("GET", "/device-config/"+template.TemplateName+"/deviceSettings", "")
		response, _ := app.Test(request)
		body, _ := io.ReadAll(response.Body)

		// Assert: check the results
		if assert.Equal(t, fiber.StatusOK, response.StatusCode) == false {
			fmt.Println("response body: " + string(body))
		}

		var deviceSettings DeviceSetting
		err = json.Unmarshal(body, &deviceSettings)
		assert.NoError(t, err)

		fmt.Printf("Received PowerConfig: %v\n", deviceSettings)

		assert.Equal(t, pc.ID, deviceSettings.ID)
		assert.Equal(t, pc.Version, deviceSettings.Version)
		// assert other fields here

		// Teardown: cleanup after test
		test.TruncateTables(pdb.DBS().Writer.DB, t)
	})
}
func TestGetDBCFileByTemplateName(t *testing.T) {
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

	c := NewDeviceConfigController(&config.Settings{Port: "3000"}, &logger, pdb.DBS().Reader.DB)
	app := fiber.New()
	app.Get("/device-config/:template_name/dbc-file", c.GetDBCFileByTemplateName)

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

		// Teardown: cleanup after test
		test.TruncateTables(pdb.DBS().Writer.DB, t)
	})
}

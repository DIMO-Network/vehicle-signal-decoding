package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/test"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

const migrationsDirRelPath = "../../migrations"

func TestGetPIDsByTemplate(t *testing.T) {

	// Arrange: global db and route setup
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

	pidService := mock_services.NewMockPIDService(mockCtrl)

	templateName := "exampleTemplate"
	c := NewDeviceConfigController(&config.Settings{Port: "3000"}, &logger, pidService, pdb.DBS)
	app := fiber.New()
	app.Get("/device-config/pid/:template_name", c.GetPIDsByTemplate)

	t.Run("GET - PID by Template", func(t *testing.T) {
		// Arrange: db, insert some PIDConfig data
		samplePID := PIDConfig{
			Name:            "EngineTemperature",
			Header:          "7E8",
			Mode:            "01",
			PID:             "05",
			Formula:         "A*5",
			IntervalSeconds: 60,
			TemplateName:    templateName,
		}
		err := samplePID.Insert(ctx, pdb.DBS().Writer, boil.Infer())
		assert.NoError(t, err)

		pidService.EXPECT().GetPIDsByTemplate(gomock.Any(), templateName).Times(1).
			Return([]models.PIDConfig{samplePID}, nil)

		// Act: make the request
		request := test.BuildRequest("GET", "/device-config/pid/"+templateName, "")
		response, _ := app.Test(request)
		body, _ := io.ReadAll(response.Body)

		// Assert: check the results
		if assert.Equal(t, fiber.StatusOK, response.StatusCode) == false {
			fmt.Println("response body: " + string(body))
		}

		pids := make([]PIDConfig, 0)
		err = json.Unmarshal(body, &pids)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(pids))
		assert.Equal(t, samplePID.Name, pids[0].Name)

		// Teardown: cleanup after test
		test.TruncateTables(pdb.DBS().Writer.DB, t)
	})
}

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
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

const migrationsDirRelPath = "../../migrations"

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

	// Insert test data into database
	templateName := "exampleTemplate"
	samplePID := PIDConfig{
		ID:              1,
		Header:          "7E8",
		Mode:            "01",
		Pid:             "05",
		Formula:         "A*5",
		IntervalSeconds: 60,
		Version:         "1.0",
	}

	_, err := pdb.DBS().Writer.ExecContext(ctx,
		`INSERT INTO pid_config (id, header, mode, pid, formula, interval_seconds, version, template_name) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		samplePID.ID, samplePID.Header, samplePID.Mode, samplePID.Pid, samplePID.Formula, samplePID.IntervalSeconds, samplePID.Version, templateName)

	c := NewDeviceConfigController(&config.Settings{Port: "3000"}, &logger, pdb.DBS().Reader.DB)
	app := fiber.New()
	app.Get("/device-config/pid/:template_name", c.GetPIDsByTemplate)

	t.Run("GET - PID by Template", func(t *testing.T) {

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
		assert.Equal(t, samplePID.ID, pids[0].ID)
	})
}

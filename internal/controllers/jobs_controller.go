package controllers

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/ethutil"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/services"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq" //nolint
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type JobsController struct {
	settings      *config.Settings
	log           *zerolog.Logger
	db            *sql.DB
	userDeviceSvc services.UserDevicesService
	deviceDefSvc  services.DeviceDefinitionsService
}

// NewJobsController constructor
func NewJobsController(settings *config.Settings, logger *zerolog.Logger, database *sql.DB, userDeviceSvc services.UserDevicesService, deviceDefSvc services.DeviceDefinitionsService) JobsController {
	return JobsController{
		settings:      settings,
		log:           logger,
		db:            database,
		userDeviceSvc: userDeviceSvc,
		deviceDefSvc:  deviceDefSvc,
	}

}

type JobResponse struct {
	ID      string `json:"id"`
	Command string `json:"command"`
	Status  string `json:"status"`
}

// GetJobsFromEthAddr godoc
// @Description  Retrieve the jobs based on device's Ethereum Address.
// @Tags         vehicle-signal-decoding
// @Produce      json
// @Success      200 {object} JobResponse "Successfully retrieved jobs"
// @Failure 400  "incorrect eth addr format"
// @Param        ethAddr  path   string  false  "Ethereum Address"
// @Router       /device-config/eth-addr/{ethAddr}/jobs [get]
func (d *JobsController) GetJobsFromEthAddr(c *fiber.Ctx) error {
	ethAddr := c.Params("ethAddr")

	ethAddrBytes, err := ethutil.ResolveEtherumAddressFromString(ethAddr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("invalid ethereum address: %s", ethAddr)})
	}

	jobs, err := models.Jobs(models.JobWhere.DeviceEthereumAddress.EQ(ethAddrBytes)).All(c.Context(), d.db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get jobs"})
	}

	jobResponse := make([]JobResponse, 0)

	for _, item := range jobs {
		jobResponse = append(jobResponse, JobResponse{
			ID:      item.ID,
			Command: item.Command,
			Status:  item.Status,
		})
	}

	return c.Status(fiber.StatusOK).JSON(jobResponse)
}

// GetJobsPendingFromEthAddr godoc
// @Description  Retrieve the jobs based on device's Ethereum Address.
// @Tags         vehicle-signal-decoding
// @Produce      json
// @Success      200 {object} JobResponse "Successfully retrieved jobs"
// @Failure 400  "incorrect eth addr format"
// @Param        ethAddr  path   string  false  "Ethereum Address"
// @Router       /device-config/eth-addr/{ethAddr}/jobs/pending [get]
func (d *JobsController) GetJobsPendingFromEthAddr(c *fiber.Ctx) error {
	ethAddr := c.Params("ethAddr")

	ethAddrBytes, err := ethutil.ResolveEtherumAddressFromString(ethAddr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("invalid ethereum address: %s", ethAddr)})
	}

	jobs, err := models.Jobs(models.JobWhere.DeviceEthereumAddress.EQ(ethAddrBytes),
		models.JobWhere.Status.EQ("PENDING")).All(c.Context(), d.db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get jobs"})
	}

	jobResponse := make([]JobResponse, 0)

	for _, item := range jobs {
		jobResponse = append(jobResponse, JobResponse{
			ID:      item.ID,
			Command: item.Command,
			Status:  item.Status,
		})
	}

	return c.Status(fiber.StatusOK).JSON(jobResponse)
}

// PatchJobsFromEthAddr godoc
// @Description  Path job status based on device's Ethereum Address.
// @Tags         vehicle-signal-decoding
// @Produce      json
// @Success      200
// @Failure 404  "Not Found - No job found with id"
// @Failure 400  "incorrect eth addr format"
// @Param        ethAddr  path   string  false  "Ethereum Address"
// @Param        jobId    path   string  false  "Job ID"
// @Param        status    path   string  false  "Status to set"
// @Router       /device-config/eth-addr/{ethAddr}/jobs/{jobId}/{status} [patch]
func (d *JobsController) PatchJobsFromEthAddr(c *fiber.Ctx) error {
	id := c.Params("jobId")
	status := c.Params("status")

	job, err := models.Jobs(models.JobWhere.ID.EQ(id)).One(c.Context(), d.db)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get job"})
		}

		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": fmt.Sprintf("could not find job id: %s", id)})
		}
	}

	job.Status = status
	job.LastExecution = null.NewTime(time.Now(), true)

	if _, err := job.Update(c.Context(), d.db, boil.Infer()); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update the job"})
	}

	var jobResponse JobResponse
	return c.Status(fiber.StatusOK).JSON(jobResponse)
}

package commands

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/services"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type RunTestSignalCommandHandler struct {
	DBS               func() *db.ReaderWriter
	userDeviceService services.UserDeviceService
	logger            zerolog.Logger
}

func NewRunTestSignalCommandHandler(dbs func() *db.ReaderWriter, logger zerolog.Logger, userDeviceService services.UserDeviceService) RunTestSignalCommandHandler {
	return RunTestSignalCommandHandler{DBS: dbs, logger: logger, userDeviceService: userDeviceService}
}

type RunTestSignalCommandRequest struct {
	AutoPIUnitID string
	Time         time.Time
	Signals      map[string]RunTestSignalItemCommandRequest
}

type RunTestSignalItemCommandRequest struct {
	Value any    `json:"value"`
	Time  string `json:"_stamp"`
}

type RunTestSignalCommandResponse struct {
}

func (h RunTestSignalCommandHandler) Execute(ctx context.Context, command *RunTestSignalCommandRequest) error {

	// Get user device
	userDevice, err := h.userDeviceService.GetUserDeviceServiceByAutoPIUnitID(ctx, command.AutoPIUnitID)
	if err != nil {
		return err
	}
	if userDevice == nil {
		return fmt.Errorf("User Device not found associate to autopi_unit_id %s", command.AutoPIUnitID)
	}

	// Validate Signals exists
	for k := range command.Signals {

		value := ""
		switch v := command.Signals[k].Value.(type) {
		case string:
			value = v
		case int32, int64:
			value = strconv.Itoa(v.(int))
		case float32, float64:
			value = fmt.Sprintf("%v", v)
		}

		h.logger.Info().Str("signal_name", k).
			Str("signal_value", value)
		dbcCode, err := models.DBCCodes(models.DBCCodeWhere.Name.EQ(k)).One(ctx, h.DBS().Reader)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				dbc := &models.DBCCode{}
				dbc.ID = ksuid.New().String()
				dbc.Name = k
				dbc.RecordingEnabled = true
				dbc.MaxSampleSize = 5

				err = dbc.Insert(ctx, h.DBS().Writer, boil.Infer())
				if err != nil {
					return fmt.Errorf("error inserting dbc_code")
				}

				dbcCode = dbc
			}
		}

		if !dbcCode.RecordingEnabled {
			return fmt.Errorf("recording is not enabled")
		}

		testSignalsCount, err := models.TestSignals(models.TestSignalWhere.AutopiUnitID.EQ(command.AutoPIUnitID)).
			Count(ctx, h.DBS().Reader)

		if err != nil {
			return fmt.Errorf("error getting count() test signals for autopi_unit_id %s", command.AutoPIUnitID)
		}

		if int(testSignalsCount) >= dbcCode.MaxSampleSize {
			return fmt.Errorf("reached signal limit. Signal tests %d", dbcCode.MaxSampleSize)
		}

		// Insert test_signals
		test := models.TestSignal{}
		test.ID = ksuid.New().String()
		test.DeviceDefinitionID = userDevice.DeviceDefinitionID
		test.DBCCodesID = dbcCode.ID
		test.AutopiUnitID = command.AutoPIUnitID
		test.Value = value
		test.VehicleTimestamp = time.Now() //todo: validate with james
		test.Approved = false

		err = test.Insert(ctx, h.DBS().Writer, boil.Infer())
		if err != nil {
			return fmt.Errorf("error inserting test signal: %s", k)
		}
	}

	return nil
}

package commands

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/volatiletech/null/v8"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/services"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

//go:generate mockgen -source run_test_signal.go -destination mocks/run_test_signal_mock.go

type RunTestSignalCommandHandler interface {
	Execute(ctx context.Context, command *RunTestSignalCommandRequest) error
}

type runTestSignalCommandHandler struct {
	DBS               func() *db.ReaderWriter
	userDeviceService services.UserDevicesService
	logger            zerolog.Logger
}

func NewRunTestSignalCommandHandler(dbs func() *db.ReaderWriter, logger zerolog.Logger, userDeviceService services.UserDevicesService) RunTestSignalCommandHandler {
	return runTestSignalCommandHandler{DBS: dbs, logger: logger, userDeviceService: userDeviceService}
}

type RunTestSignalCommandRequest struct {
	AutoPIUnitID string
	Time         time.Time
	Signals      map[string]RunTestSignalItemCommandRequest
}

type RunTestSignalItemCommandRequest struct {
	Value any    `json:"value"`
	Time  string `json:"_stamp"` //nolint
}

type RunTestSignalCommandResponse struct {
}

func (h runTestSignalCommandHandler) Execute(ctx context.Context, command *RunTestSignalCommandRequest) error {

	// Get user device
	userDevice, err := h.userDeviceService.GetUserDeviceByAutoPIUnitID(ctx, command.AutoPIUnitID)
	if err != nil {
		return err
	}
	if userDevice == nil {
		return fmt.Errorf("failed to find user device associated to autopi_unit_id %s", command.AutoPIUnitID)
	}

	signals, err := json.Marshal(command.Signals)
	if err != nil {
		return err
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

		localLog := h.logger.With().Str("signal_name", k).
			Str("signal_value", value).Str("autopi_unit_id", command.AutoPIUnitID).Logger()

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
					localLog.Err(err).Msg("error inserting dbc_code")
					continue
				}

				dbcCode = dbc
			}
		}

		if !dbcCode.RecordingEnabled {
			localLog.Info().Msg("recording is not enabled for dbc_code")
			continue
		}

		testSignalsCount, err := models.TestSignals(models.TestSignalWhere.AutopiUnitID.EQ(command.AutoPIUnitID)).
			Count(ctx, h.DBS().Reader)

		if err != nil {
			localLog.Err(err).Msgf("error getting count() test signals for autopi_unit_id %s", command.AutoPIUnitID)
			continue
		}

		if int(testSignalsCount) >= dbcCode.MaxSampleSize {
			// don't want to log this b/c will happen a lot
			//"reached signal sample size limit. Signal tests %d", dbcCode.MaxSampleSize)
			continue
		}

		// Insert test_signals
		test := models.TestSignal{
			ID:                 ksuid.New().String(),
			DeviceDefinitionID: userDevice.DeviceDefinitionID,
			DBCCodesID:         dbcCode.ID,
			AutopiUnitID:       command.AutoPIUnitID,
			Value:              value,
			Signals:            null.JSONFrom(signals),
			VehicleTimestamp:   command.Time,
			Approved:           false,
			UserDeviceID:       strings.TrimSpace(userDevice.UserDeviceID),
		}

		err = test.Insert(ctx, h.DBS().Writer, boil.Infer())
		if err != nil {
			localLog.Err(err).Msg("error inserting test signal")
		}
	}

	return nil
}

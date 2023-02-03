package commands

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/services"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/pkg/errors"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"strconv"
	"time"

	"github.com/DIMO-Network/shared/db"
)

type RunTestSignalCommandHandler struct {
	DBS               func() *db.ReaderWriter
	userDeviceService services.UserDeviceService
}

func NewRunTestSignalCommandHandler(dbs func() *db.ReaderWriter, userDeviceService services.UserDeviceService) RunTestSignalCommandHandler {
	return RunTestSignalCommandHandler{DBS: dbs, userDeviceService: userDeviceService}
}

type RunTestSignalCommandRequest struct {
	AutoPIUnitID string
	Time         time.Time
	Signals      map[string]RunTestSignalItemCommandRequest
}

type RunTestSignalItemCommandRequest struct {
	Value any `json:"value"`
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
		fmt.Printf("key[%s] value[%s]\n", k, command.Signals[k])
		item, err := models.DBCCodes(models.DBCCodeWhere.Name.EQ(k)).One(ctx, h.DBS().Reader)
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

				item = dbc
			}
		}

		if !item.RecordingEnabled {
			return fmt.Errorf("recording is not enabled")
		}

		// Insert test_signals
		test := models.TestSignal{}
		test.ID = ksuid.New().String()
		test.DeviceDefinitionID = userDevice.DeviceDefinitionID
		test.DBCCodesID = item.ID
		test.AutopiUnitID = command.AutoPIUnitID
		switch v := command.Signals[k].Value.(type) {
		case string:
			test.Value = v
		case int32, int64:
			test.Value = strconv.Itoa(v.(int))
		case float32, float64:
			test.Value = fmt.Sprintf("%v", v)
		}
		test.Approved = false

		err = test.Insert(ctx, h.DBS().Writer, boil.Infer())
		if err != nil {
			return fmt.Errorf("error inserting test signal: %s", k)
		}
	}

	return nil
}

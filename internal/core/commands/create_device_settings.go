package commands

import (
	"context"
	"encoding/json"

	"github.com/volatiletech/null/v8"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/DIMO-Network/shared/db"
)

type CreateDeviceSettingsCommandHandler struct {
	DBS func() *db.ReaderWriter
}

func NewCreateDeviceSettingsCommandHandler(dbs func() *db.ReaderWriter) CreateDeviceSettingsCommandHandler {
	return CreateDeviceSettingsCommandHandler{DBS: dbs}
}

type SettingsData struct {
	SafetyCutOutVoltage             float64 `json:"safety_cut_out_voltage"`
	SleepTimerEventDrivenPeriodSecs float64 `json:"sleep_timer_event_driven_period_secs"`
	WakeTriggerVoltageLevel         float64 `json:"wake_trigger_voltage_level"`
}

type CreateDeviceSettingsCommandRequest struct {
	Name     string
	Settings SettingsData `json:"settings"`
}

type CreateDeviceSettingsCommandResponse struct {
	Name string
}

func (h CreateDeviceSettingsCommandHandler) Execute(ctx context.Context, req *CreateDeviceSettingsCommandRequest) (*CreateDeviceSettingsCommandResponse, error) {

	exists, err := models.DeviceSettings(models.DeviceSettingWhere.Name.EQ(req.Name)).Exists(ctx, h.DBS().Reader)
	if err != nil {
		return nil, &exceptions.InternalError{
			Err: errors.Wrapf(err, "error checking if device setting exists: %s", req.Name),
		}
	}
	if exists {
		return nil, &exceptions.ConflictError{
			Err: errors.Errorf("device setting already exists: %s", req.Name),
		}
	}

	settingsBytes, err := json.Marshal(req.Settings)
	if err != nil {
		return nil, &exceptions.InternalError{
			Err: errors.Wrapf(err, "error serializing settings for device setting: %s", req.Name),
		}
	}

	settingsJSON := null.JSONFrom(settingsBytes)

	deviceSetting := &models.DeviceSetting{
		Name:     req.Name,
		Settings: settingsJSON,
	}

	err = deviceSetting.Insert(ctx, h.DBS().Writer, boil.Infer())
	if err != nil {
		return nil, &exceptions.InternalError{
			Err: errors.Wrapf(err, "error inserting device setting with template name: %s", req.Name),
		}
	}

	return &CreateDeviceSettingsCommandResponse{Name: deviceSetting.Name}, nil
}

package commands

import (
	"context"

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

type CreateDeviceSettingsCommandRequest struct {
	ID                                     int64
	TemplateName                           string
	BatteryCriticalLevelVoltage            float64
	SafetyCutOutVoltage                    float64
	SleepTimerEventDrivenInterval          float64
	SleepTimerEventDrivenPeriod            float64
	SleepTimerInactivityAfterSleepInterval float64
	SleepTimerInactivityFallbackInterval   float64
	WakeTriggerVoltageLevel                float64
}

type CreateDeviceSettingsCommandResponse struct {
	Name string
}

func (h CreateDeviceSettingsCommandHandler) Execute(ctx context.Context, req *CreateDeviceSettingsCommandRequest) (*CreateDeviceSettingsCommandResponse, error) {

	exists, err := models.DeviceSettings(models.DeviceSettingWhere.TemplateName.EQ(req.TemplateName)).Exists(ctx, h.DBS().Reader)
	if err != nil {
		return nil, &exceptions.InternalError{
			Err: errors.Wrapf(err, "error checking if device setting exists: %s", req.TemplateName),
		}
	}
	if exists {
		return nil, &exceptions.ConflictError{
			Err: errors.Errorf("device setting already exists: %s", req.TemplateName),
		}
	}

	deviceSetting := &models.DeviceSetting{
		ID:                                     req.ID,
		TemplateName:                           req.TemplateName,
		BatteryCriticalLevelVoltage:            req.BatteryCriticalLevelVoltage,
		SafetyCutOutVoltage:                    req.SafetyCutOutVoltage,
		SleepTimerEventDrivenInterval:          req.SleepTimerEventDrivenInterval,
		SleepTimerEventDrivenPeriod:            req.SleepTimerEventDrivenPeriod,
		SleepTimerInactivityAfterSleepInterval: req.SleepTimerInactivityAfterSleepInterval,
		SleepTimerInactivityFallbackInterval:   req.SleepTimerInactivityFallbackInterval,
		WakeTriggerVoltageLevel:                req.WakeTriggerVoltageLevel,
	}

	err = deviceSetting.Insert(ctx, h.DBS().Writer, boil.Infer())
	if err != nil {
		return nil, &exceptions.InternalError{
			Err: errors.Wrapf(err, "error inserting device setting with template name: %s", req.TemplateName),
		}
	}

	return &CreateDeviceSettingsCommandResponse{Name: deviceSetting.TemplateName}, nil
}

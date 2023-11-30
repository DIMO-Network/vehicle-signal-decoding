package commands

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/volatiletech/null/v8"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UpdateDeviceSettingsCommandHandler struct {
	DBS func() *db.ReaderWriter
}

func NewUpdateDeviceSettingsCommandHandler(dbs func() *db.ReaderWriter) UpdateDeviceSettingsCommandHandler {
	return UpdateDeviceSettingsCommandHandler{DBS: dbs}
}

type UpdateDeviceSettingsCommandRequest struct {
	TemplateName string
	Settings     SettingsData `json:"settings"`
}

type UpdateDeviceSettingsCommandResponse struct {
	TemplateName string
}

func (h UpdateDeviceSettingsCommandHandler) Execute(ctx context.Context, req *UpdateDeviceSettingsCommandRequest) (*UpdateDeviceSettingsCommandResponse, error) {

	deviceSettings, err := models.DeviceSettings(models.DeviceSettingWhere.TemplateName.EQ(req.TemplateName)).One(ctx, h.DBS().Reader)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &exceptions.NotFoundError{
				Err: fmt.Errorf("device settings not found with template name: %s", req.TemplateName),
			}
		}
		return nil, &exceptions.InternalError{
			Err: err,
		}
	}
	settingsJSON, err := json.Marshal(req.Settings)
	if err != nil {
		return nil, &exceptions.InternalError{
			Err: fmt.Errorf("failed to marshal settings to JSON: %w", err),
		}
	}

	deviceSettings.TemplateName = req.TemplateName
	deviceSettings.Settings = null.NewJSON(settingsJSON, true)

	if _, err := deviceSettings.Update(ctx, h.DBS().Writer.DB, boil.Infer()); err != nil {
		return nil, &exceptions.InternalError{
			Err: err,
		}
	}

	return &UpdateDeviceSettingsCommandResponse{TemplateName: deviceSettings.TemplateName}, nil
}

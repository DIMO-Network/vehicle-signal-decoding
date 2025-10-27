package commands

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/appmodels"

	"github.com/aarondl/null/v8"

	"github.com/DIMO-Network/shared/pkg/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	"github.com/aarondl/sqlboiler/v4/boil"
)

type UpdateDeviceSettingsCommandHandler struct {
	DBS func() *db.ReaderWriter
}

func NewUpdateDeviceSettingsCommandHandler(dbs func() *db.ReaderWriter) UpdateDeviceSettingsCommandHandler {
	return UpdateDeviceSettingsCommandHandler{DBS: dbs}
}

type UpdateDeviceSettingsCommandRequest struct {
	Name         string                 `json:"name"`
	TemplateName *string                `json:"templateName"`
	PowerTrain   string                 `json:"powerTrain"`
	Version      string                 `json:"version"`
	Settings     appmodels.SettingsData `json:"settings"`
}

type UpdateDeviceSettingsCommandResponse struct {
	Name string
}

func (h UpdateDeviceSettingsCommandHandler) Execute(ctx context.Context, req *UpdateDeviceSettingsCommandRequest) (*UpdateDeviceSettingsCommandResponse, error) {

	deviceSettings, err := models.DeviceSettings(models.DeviceSettingWhere.Name.EQ(req.Name)).One(ctx, h.DBS().Reader)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &exceptions.NotFoundError{
				Err: fmt.Errorf("device settings not found with template name: %s", req.Name),
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

	deviceSettings.Powertrain = req.PowerTrain
	deviceSettings.Settings = null.NewJSON(settingsJSON, true)
	deviceSettings.TemplateName = null.StringFromPtr(req.TemplateName)
	deviceSettings.Version = req.Version

	deviceSettings.UpdatedAt = time.Now().UTC()

	if _, err := deviceSettings.Update(ctx, h.DBS().Writer.DB, boil.Infer()); err != nil {
		return nil, &exceptions.InternalError{
			Err: err,
		}
	}

	return &UpdateDeviceSettingsCommandResponse{Name: deviceSettings.Name}, nil
}

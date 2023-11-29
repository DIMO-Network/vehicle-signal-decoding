package queries

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	grpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
)

type GetDeviceSettingsByTemplateNameQueryHandler struct {
	DBS    func() *db.ReaderWriter
	logger *zerolog.Logger
}

func NewGetDeviceSettingsByTemplateNameQueryHandler(dbs func() *db.ReaderWriter, logger *zerolog.Logger) GetDeviceSettingsByTemplateNameQueryHandler {
	return GetDeviceSettingsByTemplateNameQueryHandler{
		DBS:    dbs,
		logger: logger,
	}
}

type GetDeviceSettingsByTemplateNameQueryRequest struct {
	TemplateName string
}

func (h GetDeviceSettingsByTemplateNameQueryHandler) Handle(ctx context.Context, query *GetDeviceSettingsByTemplateNameQueryRequest) (*grpc.GetDeviceSettingByTemplateNameResponse, error) {
	item, err := models.DeviceSettings(models.DeviceSettingWhere.TemplateName.EQ(query.TemplateName)).One(ctx, h.DBS().Reader)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("device setting not found for template name: %s", query.TemplateName)
		}
		return nil, fmt.Errorf("failed to get device setting by template name: %w", err)
	}

	jsonBytes, err := item.Settings.MarshalJSON()
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to marshal settings JSON")
		return nil, fmt.Errorf("failed to marshal settings JSON: %w", err)
	}

	settingsString := string(jsonBytes)

	result := &grpc.GetDeviceSettingByTemplateNameResponse{
		DeviceSettings: &grpc.DeviceSettings{
			TemplateName: item.TemplateName,
			Settings:     settingsString,
		},
	}

	return result, nil
}

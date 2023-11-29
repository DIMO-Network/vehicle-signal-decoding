package queries

import (
	"context"
	"fmt"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	grpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
)

type GetDeviceSettingsAllQueryHandler struct {
	DBS    func() *db.ReaderWriter
	logger *zerolog.Logger
}

func NewGetDeviceSettingsAllQueryHandler(dbs func() *db.ReaderWriter, logger *zerolog.Logger) GetDeviceSettingsAllQueryHandler {
	return GetDeviceSettingsAllQueryHandler{
		DBS:    dbs,
		logger: logger,
	}
}

type GetDeviceSettingsAllQueryRequest struct {
	TemplateName string
}

func (h GetDeviceSettingsAllQueryHandler) Handle(ctx context.Context, _ *GetDeviceSettingsAllQueryRequest) (*grpc.GetDeviceSettingListResponse, error) {

	all, err := models.DeviceSettings().All(ctx, h.DBS().Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to get DeviceSettings: %w", err)
	}

	deviceSettingsList := make([]*grpc.DeviceSettings, 0, len(all))

	for _, item := range all {
		// Convert null.JSON to []byte
		jsonBytes, err := item.Settings.MarshalJSON()
		if err != nil {
			h.logger.Error().Err(err).Msgf("Failed to marshal settings for template: %s", item.TemplateName)
			continue
		}

		settingsString := string(jsonBytes)

		deviceSettings := &grpc.DeviceSettings{
			TemplateName: item.TemplateName,
			Settings:     settingsString,
		}
		deviceSettingsList = append(deviceSettingsList, deviceSettings)
	}

	result := &grpc.GetDeviceSettingListResponse{
		DeviceSettings: deviceSettingsList,
	}

	return result, nil
}

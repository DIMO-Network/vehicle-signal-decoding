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

type GetDeviceSettingsByNameQueryHandler struct {
	DBS    func() *db.ReaderWriter
	logger *zerolog.Logger
}

func NewGetDeviceSettingsByNameQueryHandler(dbs func() *db.ReaderWriter, logger *zerolog.Logger) GetDeviceSettingsByNameQueryHandler {
	return GetDeviceSettingsByNameQueryHandler{
		DBS:    dbs,
		logger: logger,
	}
}

type GetDeviceSettingsByNameQueryRequest struct {
	Name string
}

func (h GetDeviceSettingsByNameQueryHandler) Handle(ctx context.Context, query *GetDeviceSettingsByNameQueryRequest) (*grpc.GetDeviceSettingByNameResponse, error) {
	item, err := models.DeviceSettings(models.DeviceSettingWhere.Name.EQ(query.Name)).One(ctx, h.DBS().Reader)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("device setting not found for name: %s", query.Name)
		}
		return nil, fmt.Errorf("failed to get device setting by name: %w", err)
	}

	jsonBytes, err := item.Settings.MarshalJSON()
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to marshal settings JSON")
		return nil, fmt.Errorf("failed to marshal settings JSON: %w", err)
	}

	settingsString := string(jsonBytes)

	result := &grpc.GetDeviceSettingByNameResponse{
		DeviceSettings: &grpc.DeviceSettings{
			Name:       item.Name,
			Settings:   settingsString,
			Powertrain: item.Powertrain,
		},
	}

	return result, nil
}

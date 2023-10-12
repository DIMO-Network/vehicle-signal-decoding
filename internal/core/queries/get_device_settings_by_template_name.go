package queries

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"

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
		return nil, fmt.Errorf("failed to get device setting by template name: %s", err)
	}

	result := &grpc.GetDeviceSettingByTemplateNameResponse{
		DeviceSettings: &grpc.DeviceSettings{
			TemplateName:                           item.TemplateName,
			BatteryCriticalLevelVoltage:            item.BatteryCriticalLevelVoltage,
			SafetyCutOutVoltage:                    item.SafetyCutOutVoltage,
			SleepTimerEventDrivenInterval:          item.SleepTimerEventDrivenInterval,
			SleepTimerEventDrivenPeriod:            item.SleepTimerEventDrivenPeriod,
			SleepTimerInactivityAfterSleepInterval: item.SleepTimerInactivityAfterSleepInterval,
			SleepTimerInactivityFallbackInterval:   item.SleepTimerInactivityFallbackInterval,
			WakeTriggerVoltageLevel:                item.WakeTriggerVoltageLevel,
			CreatedAt:                              timestamppb.New(item.CreatedAt),
			UpdatedAt:                              timestamppb.New(item.UpdatedAt),
		},
	}

	return result, nil
}

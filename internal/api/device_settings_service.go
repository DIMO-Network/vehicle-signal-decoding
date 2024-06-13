package api

import (
	"context"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/appmodels"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/commands"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/queries"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DeviceSettingsConfigService struct {
	grpc.DeviceSettingsServiceServer
	logger *zerolog.Logger
	dbs    func() *db.ReaderWriter
}

func NewDeviceSettingsConfigService(logger *zerolog.Logger, dbs func() *db.ReaderWriter) grpc.DeviceSettingsServiceServer {
	return &DeviceSettingsConfigService{logger: logger, dbs: dbs}
}

func (s *DeviceSettingsConfigService) CreateDeviceSettings(ctx context.Context, in *grpc.UpdateDeviceSettingsRequest) (*emptypb.Empty, error) {
	service := commands.NewCreateDeviceSettingsCommandHandler(s.dbs)

	inboundDeviceSettings := in.DeviceSettings.Settings

	_, err := service.Execute(ctx, &commands.CreateDeviceSettingsCommandRequest{
		Name: in.DeviceSettings.Name,
		Settings: appmodels.SettingsData{
			SafetyCutOutVoltage:                      float64(inboundDeviceSettings.SafetyCutOutVoltage),
			SleepTimerEventDrivenPeriodSecs:          float64(inboundDeviceSettings.SleepTimerEventDrivenPeriodSecs),
			WakeTriggerVoltageLevel:                  float64(inboundDeviceSettings.WakeTriggerVoltageLevel),
			MinVoltageOBDLoggers:                     float64(inboundDeviceSettings.MinVoltageObdLoggers),
			LocationFrequencySecs:                    float64(inboundDeviceSettings.LocationFrequencySecs),
			SleepTimerEventDrivenIntervalSecs:        float64(inboundDeviceSettings.SleepTimerEventDrivenIntervalSecs),
			SleepTimerInactivityFallbackIntervalSecs: float64(inboundDeviceSettings.SleepTimerInactivityFallbackIntervalSecs),
			SleepTimerInactivityAfterSleepSecs:       float64(inboundDeviceSettings.SleepTimerInactivityAfterSleepSecs),
			BatteryCriticalLevelVoltage:              float64(inboundDeviceSettings.BatteryCriticalLevelVoltage),
		},
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error executing CreateDeviceSettingsCommand: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *DeviceSettingsConfigService) UpdateDeviceSettings(ctx context.Context, in *grpc.UpdateDeviceSettingsRequest) (*emptypb.Empty, error) {
	service := commands.NewUpdateDeviceSettingsCommandHandler(s.dbs)
	inboundDeviceSettings := in.DeviceSettings.Settings
	_, err := service.Execute(ctx, &commands.UpdateDeviceSettingsCommandRequest{
		Name: in.DeviceSettings.Name,
		Settings: appmodels.SettingsData{
			SafetyCutOutVoltage:                      float64(inboundDeviceSettings.SafetyCutOutVoltage),
			SleepTimerEventDrivenPeriodSecs:          float64(inboundDeviceSettings.SleepTimerEventDrivenPeriodSecs),
			WakeTriggerVoltageLevel:                  float64(inboundDeviceSettings.WakeTriggerVoltageLevel),
			MinVoltageOBDLoggers:                     float64(inboundDeviceSettings.MinVoltageObdLoggers),
			LocationFrequencySecs:                    float64(inboundDeviceSettings.LocationFrequencySecs),
			SleepTimerEventDrivenIntervalSecs:        float64(inboundDeviceSettings.SleepTimerEventDrivenIntervalSecs),
			SleepTimerInactivityFallbackIntervalSecs: float64(inboundDeviceSettings.SleepTimerInactivityFallbackIntervalSecs),
			SleepTimerInactivityAfterSleepSecs:       float64(inboundDeviceSettings.SleepTimerInactivityAfterSleepSecs),
			BatteryCriticalLevelVoltage:              float64(inboundDeviceSettings.BatteryCriticalLevelVoltage),
		},
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error executing UpdateDeviceSettingsCommand: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *DeviceSettingsConfigService) GetDeviceSettingList(ctx context.Context, _ *emptypb.Empty) (*grpc.GetDeviceSettingListResponse, error) {
	service := queries.NewGetDeviceSettingsAllQueryHandler(s.dbs, s.logger)
	response, err := service.Handle(ctx, &queries.GetDeviceSettingsAllQueryRequest{})

	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *DeviceSettingsConfigService) GetDeviceSettingByName(ctx context.Context, in *grpc.GetDeviceSettingByNameRequest) (*grpc.GetDeviceSettingByNameResponse, error) {
	service := queries.NewGetDeviceSettingsByNameQueryHandler(s.dbs, s.logger)
	response, err := service.Handle(ctx, &queries.GetDeviceSettingsByNameQueryRequest{
		Name: in.Name,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

package api

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/commands"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/queries"

	db "github.com/DIMO-Network/shared/db"
	grpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
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

	// Deserialize the JSON settings into SettingsData
	var settingsData commands.SettingsData
	err := json.Unmarshal([]byte(in.DeviceSettings.Settings), &settingsData)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Failed to parse settings JSON: %v", err)
	}

	_, err = service.Execute(ctx, &commands.CreateDeviceSettingsCommandRequest{
		Name:     in.DeviceSettings.Name,
		Settings: settingsData,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error executing CreateDeviceSettingsCommand: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *DeviceSettingsConfigService) UpdateDeviceSettings(ctx context.Context, in *grpc.UpdateDeviceSettingsRequest) (*emptypb.Empty, error) {
	service := commands.NewUpdateDeviceSettingsCommandHandler(s.dbs)

	// Deserialize the JSON settings into SettingsData
	var settingsData commands.SettingsData
	err := json.Unmarshal([]byte(in.DeviceSettings.Settings), &settingsData)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Failed to parse settings JSON: %v", err)
	}

	_, err = service.Execute(ctx, &commands.UpdateDeviceSettingsCommandRequest{
		Name: in.DeviceSettings.Name,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error executing UpdateDeviceSettingsCommand: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *DeviceSettingsConfigService) GetDeviceSettingList(ctx context.Context, in *grpc.GetDeviceSettingListRequest) (*grpc.GetDeviceSettingListResponse, error) {
	service := queries.NewGetDeviceSettingsAllQueryHandler(s.dbs, s.logger)
	response, err := service.Handle(ctx, &queries.GetDeviceSettingsAllQueryRequest{
		Name: *in.Name,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *DeviceSettingsConfigService) GetDeviceSettingByTemplateName(ctx context.Context, in *grpc.GetDeviceSettingByTemplateNameRequest) (*grpc.GetDeviceSettingByTemplateNameResponse, error) {
	service := queries.NewGetDeviceSettingsByTemplateNameQueryHandler(s.dbs, s.logger)
	response, err := service.Handle(ctx, &queries.GetDeviceSettingsByTemplateNameQueryRequest{
		Name: in.Name,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

package api

import (
	"context"

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
	_, err := service.Execute(ctx, &commands.CreateDeviceSettingsCommandRequest{
		TemplateName: in.DeviceSettings.TemplateName,
	})

	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *DeviceSettingsConfigService) UpdateDeviceSettings(ctx context.Context, in *grpc.UpdateDeviceSettingsRequest) (*emptypb.Empty, error) {
	service := commands.NewUpdateDeviceSettingsCommandHandler(s.dbs)
	_, err := service.Execute(ctx, &commands.UpdateDeviceSettingsCommandRequest{
		TemplateName: in.DeviceSettings.TemplateName,
	})

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DeviceSettingsConfigService) GetDeviceSettingList(ctx context.Context, in *grpc.GetDeviceSettingListRequest) (*grpc.GetDeviceSettingListResponse, error) {
	service := queries.NewGetDeviceSettingsAllQueryHandler(s.dbs, s.logger)
	response, err := service.Handle(ctx, &queries.GetDeviceSettingsAllQueryRequest{
		TemplateName: *in.TemplateName,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *DeviceSettingsConfigService) GetDeviceSettingByTemplateName(ctx context.Context, in *grpc.GetDeviceSettingByTemplateNameRequest) (*grpc.GetDeviceSettingByTemplateNameResponse, error) {
	service := queries.NewGetDeviceSettingsByTemplateNameQueryHandler(s.dbs, s.logger)
	response, err := service.Handle(ctx, &queries.GetDeviceSettingsByTemplateNameQueryRequest{
		TemplateName: in.TemplateName,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

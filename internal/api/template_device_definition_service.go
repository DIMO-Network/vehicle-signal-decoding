package api

import (
	"context"
	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TemplateDeviceDefinitionService struct {
	grpc.TemplateDeviceDefinitionServiceServer
	logger *zerolog.Logger
	dbs    func() *db.ReaderWriter
}

func NewTemplateDeviceDefinitionService(logger *zerolog.Logger, dbs func() *db.ReaderWriter) *TemplateDeviceDefinitionService {
	return &TemplateDeviceDefinitionService{
		logger: logger,
		dbs:    dbs,
	}
}

func (s *TemplateDeviceDefinitionService) CreateTemplateDeviceDefinition(ctx context.Context, request *grpc.TemplateDeviceDefinition) (*emptypb.Empty, error) {
	service := NewTemplateDeviceDefinitionService(s.logger, s.dbs)

	return service.CreateTemplateDeviceDefinition(ctx, request)
}

func (s *TemplateDeviceDefinitionService) UpdateTemplateDeviceDefinition(ctx context.Context, request *grpc.TemplateDeviceDefinition) (*emptypb.Empty, error) {
	service := NewTemplateDeviceDefinitionService(s.logger, s.dbs)

	return service.UpdateTemplateDeviceDefinition(ctx, request)
}

func (s *TemplateDeviceDefinitionService) GetTemplateDeviceDefinitionById(ctx context.Context, request *grpc.GetTemplateDeviceDefinitionByIdRequest) (*grpc.TemplateDeviceDefinition, error) {
	service := NewTemplateDeviceDefinitionService(s.logger, s.dbs)

	return service.GetTemplateDeviceDefinitionById(ctx, request)
}

func (s *TemplateDeviceDefinitionService) GetTemplateDeviceDefinitions(ctx context.Context, empty *emptypb.Empty) (*grpc.GetTemplateDeviceDefinitionResponse, error) {
	service := NewTemplateDeviceDefinitionService(s.logger, s.dbs)

	return service.GetTemplateDeviceDefinitions(ctx, empty)
}

func (s *TemplateDeviceDefinitionService) DeleteTemplateDeviceDefinition(ctx context.Context, request *grpc.DeleteTemplateDeviceDefinitionRequest) (*emptypb.Empty, error) {
	service := NewTemplateDeviceDefinitionService(s.logger, s.dbs)

	return service.DeleteTemplateDeviceDefinition(ctx, request)
}

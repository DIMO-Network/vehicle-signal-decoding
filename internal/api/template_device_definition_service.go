package api

import (
	"context"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/commands"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/queries"

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
	service := commands.NewCreateTemplateDeviceDefinitionCommandHandler(s.dbs)

	rq := commands.CreateTemplateDeviceDefinitionCommand{
		DeviceDefinitionID: request.DeviceDefinitionId,
		TemplateName:       request.TemplateName,
		DeviceStyleID:      request.DeviceStyleId,
	}

	response, err := service.Execute(ctx, rq)

	if err != nil {
		s.logger.Error().Err(err).Msg("Error while creating template device definition")
		return nil, err
	}

	return response, nil
}

func (s *TemplateDeviceDefinitionService) UpdateTemplateDeviceDefinition(ctx context.Context, request *grpc.TemplateDeviceDefinition) (*emptypb.Empty, error) {
	service := commands.NewUpdateTemplateDeviceDefinitionCommandHandler(s.dbs)

	rq := commands.UpdateTemplateDeviceDefinitionCommand{
		ID:                 request.Id,
		DeviceDefinitionID: request.DeviceDefinitionId,
		TemplateName:       request.TemplateName,
		DeviceStyleID:      request.DeviceStyleId,
	}

	response, err := service.Execute(ctx, rq)

	if err != nil {
		s.logger.Error().Err(err).Msg("Error while updating template device definition")
		return nil, err
	}

	return response, nil

}

func (s *TemplateDeviceDefinitionService) GetTemplateDeviceDefinition(ctx context.Context, request *grpc.GetTemplateDeviceDefinitionByIdRequest) (*grpc.TemplateDeviceDefinition, error) {
	service := queries.NewGetTemplateDeviceDefinitionByIDQueryHandler(s.dbs)

	rq := queries.GetTemplateDeviceDefinitionByIDQuery{
		ID: request.Id,
	}

	response, err := service.Handle(ctx, rq)

	if err != nil {
		s.logger.Error().Err(err).Msg("Error while getting template device definition")
	}

	return response, nil
}

func (s *TemplateDeviceDefinitionService) GetTemplateDeviceDefinitions(ctx context.Context, _ *emptypb.Empty) (*grpc.GetTemplateDeviceDefinitionResponse, error) {
	service := queries.NewGetTemplateDeviceDefinitionAllQueryHandler(s.dbs)

	rq := queries.GetTemplateDeviceDefinitionAllQuery{}

	response, err := service.Handle(ctx, rq)

	if err != nil {
		s.logger.Error().Err(err).Msg("Error while getting template device definitions")
		return nil, err
	}

	return response, nil
}

func (s *TemplateDeviceDefinitionService) DeleteTemplateDeviceDefinition(ctx context.Context, request *grpc.DeleteTemplateDeviceDefinitionRequest) (*emptypb.Empty, error) {
	service := commands.NewDeleteTemplateDeviceDefinitionCommandHandler(s.dbs)

	rq := commands.DeleteTemplateDeviceDefinitionCommand{
		ID: request.Id,
	}

	response, err := service.Execute(ctx, rq)

	if err != nil {
		s.logger.Error().Err(err).Msg("Error while deleting template device definition")
		return nil, err
	}

	return response, nil
}

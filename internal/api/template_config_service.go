package api

import (
	"context"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/commands"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/queries"

	"github.com/DIMO-Network/shared/pkg/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TemplateConfigService struct {
	grpc.TemplateConfigServiceServer
	logger *zerolog.Logger
	dbs    func() *db.ReaderWriter
}

func NewTemplateConfigService(logger *zerolog.Logger, dbs func() *db.ReaderWriter) grpc.TemplateConfigServiceServer {
	return &TemplateConfigService{logger: logger, dbs: dbs}
}

func (s *TemplateConfigService) CreateTemplate(ctx context.Context, in *grpc.UpdateTemplateRequest) (*emptypb.Empty, error) {
	service := commands.NewCreateTemplateCommandHandler(s.dbs)
	_, err := service.Execute(ctx, &commands.CreateTemplateCommandRequest{
		Name:               in.Template.Name,
		ParentTemplateName: in.Template.ParentTemplateName,
		Version:            in.Template.Version,
		Protocol:           in.Template.Protocol,
		Powertrain:         in.Template.Powertrain,
		DBC:                in.Template.Dbc,
		TemplateVehicles:   in.Template.TemplateVehicles,
		Comments:           in.Template.Comments,
	})

	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *TemplateConfigService) UpdateTemplate(ctx context.Context, in *grpc.UpdateTemplateRequest) (*emptypb.Empty, error) {
	service := commands.NewUpdateTemplateCommandHandler(s.dbs)
	_, err := service.Execute(ctx, &commands.UpdateTemplateCommandRequest{
		Name:               in.Template.Name,
		ParentTemplateName: in.Template.ParentTemplateName,
		Version:            in.Template.Version,
		Protocol:           in.Template.Protocol,
		Powertrain:         in.Template.Powertrain,
		DBC:                in.Template.Dbc,
		TemplateVehicles:   in.Template.TemplateVehicles,
		Comments:           in.Template.Comments,
	})

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *TemplateConfigService) GetTemplateList(ctx context.Context, in *grpc.GetTemplateListRequest) (*grpc.GetTemplateListResponse, error) {
	service := queries.NewGetTemplatesAllQueryHandler(s.dbs, s.logger)
	response, err := service.Handle(ctx, &queries.GetTemplatesAllQueryRequest{
		Protocol:   in.Protocol,
		Powertrain: in.Powertrain,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *TemplateConfigService) GetTemplateByName(ctx context.Context, in *grpc.GetTemplateByNameRequest) (*grpc.GetTemplateByNameResponse, error) {
	service := queries.NewGetTemplateByNameQueryHandler(s.dbs, s.logger)
	response, err := service.Handle(ctx, &queries.GetTemplateByNameQueryRequest{
		Name: in.Name,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

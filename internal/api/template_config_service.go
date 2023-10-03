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

type TemplateConfigService struct {
	grpc.TemplateConfigServiceServer
	logger *zerolog.Logger
	DBS    func() *db.ReaderWriter
}

func NewTemplateConfigService(logger *zerolog.Logger, dbs func() *db.ReaderWriter) grpc.TemplateConfigServiceServer {
	return &TemplateConfigService{logger: logger, DBS: dbs}
}

func (s *TemplateConfigService) CreateTemplate(ctx context.Context, in *grpc.UpdateTemplateRequest) (*emptypb.Empty, error) {
	service := commands.NewCreateTemplateCommandHandler(s.DBS)
	_, err := service.Execute(ctx, &commands.CreateTemplateCommandRequest{
		Name:               in.GetTemplate().GetName(),
		ParentTemplateName: in.GetTemplate().GetParentTemplateName(),
		Version:            in.GetTemplate().GetVersion(),
		Protocol:           in.GetTemplate().GetProtocol(),
		Powertrain:         in.GetTemplate().GetPowertrain(),
		HasDBC:             in.GetTemplate().GetHasDbc(),
		//PidsCount:          in.GetTemplate().GetPidsCount(),
		DBC:              in.GetTemplate().GetDbc(),
		TemplateVehicles: in.GetTemplate().GetTemplateVehicles(),
	})

	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *TemplateConfigService) UpdateTemplate(ctx context.Context, in *grpc.UpdateTemplateRequest) (*emptypb.Empty, error) {
	service := commands.NewUpdateTemplateCommandHandler(s.DBS)
	_, err := service.Execute(ctx, &commands.UpdateTemplateCommandRequest{
		Name:               in.GetTemplate().GetName(),
		ParentTemplateName: in.GetTemplate().GetParentTemplateName(),
		Version:            in.GetTemplate().GetVersion(),
		Protocol:           in.GetTemplate().GetProtocol(),
		Powertrain:         in.GetTemplate().GetPowertrain(),
		HasDBC:             in.GetTemplate().GetHasDbc(),
		//PidsCount:          in.GetTemplate().GetPidsCount(),
		DBC:              in.GetTemplate().GetDbc(),
		TemplateVehicles: in.GetTemplate().GetTemplateVehicles(),
	})

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *TemplateConfigService) GetTemplateList(ctx context.Context, in *grpc.GetTemplateListRequest) (*grpc.GetTemplateListResponse, error) {
	service := queries.NewGetTemplatesAllQueryHandler(s.DBS, s.logger)
	response, err := service.Handle(ctx, &queries.GetTemplatesAllQueryRequest{
		Protocol:   in.Protocol,
		Powertrain: in.Powertrain,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *TemplateConfigService) GetTemplateByID(ctx context.Context, in *grpc.GetTemplateByIDRequest) (*grpc.GetTemplateByIDResponse, error) {
	service := queries.NewGetTemplateByIDQueryHandler(s.DBS, s.logger)
	response, err := service.Handle(ctx, &queries.GetTemplateByIDQueryRequest{
		ID: in.Id,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

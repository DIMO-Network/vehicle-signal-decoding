package api

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/DIMO-Network/shared/pkg/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/commands"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/queries"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/emptypb"
)

type VehicleTemplateService struct {
	grpc.VehicleTemplateServiceServer
	logger *zerolog.Logger
	dbs    func() *db.ReaderWriter
}

func NewVehicleTemplateService(logger *zerolog.Logger, dbs func() *db.ReaderWriter) grpc.VehicleTemplateServiceServer {
	return &VehicleTemplateService{
		logger: logger,
		dbs:    dbs,
	}
}

func (s *VehicleTemplateService) GetVehicleTemplates(ctx context.Context, request *grpc.GetVehicleTemplatesRequest) (*grpc.GetVehicleTemplatesResponse, error) {
	service := queries.NewGetVehicleTemplatesFilteredQueryHandler(s.dbs, s.logger)

	query := queries.GetVehicleTemplatesFilteredQueryRequest{}

	if request.Make != "" {
		query.MakeSlug = &request.Make
	}

	if request.YearStart > 0 {
		query.YearStart = &request.YearStart
	}

	if request.YearEnd > 0 {
		query.YearEnd = &request.YearEnd
	}

	if request.Template != "" {
		query.Template = &request.Template
	}

	vehicleTemplates, err := service.Handle(ctx, &query)

	if err != nil {
		return nil, err
	}

	return vehicleTemplates, nil

}

func (s *VehicleTemplateService) GetVehicleTemplate(ctx context.Context, request *grpc.GetVehicleTemplateRequest) (*grpc.VehicleTemplate, error) {
	service := queries.NewGetVehicleTemplateByIDQueryHandler(s.dbs, s.logger)

	vehicleTemplate, err := service.Handle(ctx, &queries.GetVehicleTemplateByIDQueryRequest{
		ID: request.Id,
	})

	if err != nil {
		return nil, err
	}

	return vehicleTemplate, nil
}

func (s *VehicleTemplateService) CreateVehicleTemplate(ctx context.Context, request *grpc.VehicleTemplate) (*grpc.CreateVehicleTemplateResponse, error) {
	service := commands.NewCreateVehicleTemplateCommandHandler(s.dbs)

	command := commands.CreateVehicleTemplateCommandRequest{
		TemplateName:   request.Template,
		YearStart:      request.YearStart,
		YearEnd:        request.YearEnd,
		ModelWhitelist: request.Models,
	}

	if request.Make != "" {
		command.MakeSlug = &request.Make
	}
	if command.MakeSlug == nil {
		return nil, status.Error(codes.InvalidArgument, "make slug required")
	}

	resp, err := service.Execute(ctx, &command)

	if err != nil {
		return nil, err
	}

	return &grpc.CreateVehicleTemplateResponse{
		Id: int64(*resp),
	}, nil
}

func (s *VehicleTemplateService) UpdateVehicleTemplate(ctx context.Context, request *grpc.VehicleTemplate) (*emptypb.Empty, error) {
	service := commands.NewUpdateVehicleTemplateCommandHandler(s.dbs)

	command := commands.UpdateVehicleTemplateCommandRequest{
		ID:             request.Id,
		TemplateName:   request.Template,
		YearStart:      request.YearStart,
		YearEnd:        request.YearEnd,
		ModelWhitelist: request.Models,
	}

	if request.Make != "" {
		command.MakeSlug = &request.Make
	}

	resp, err := service.Execute(ctx, &command)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *VehicleTemplateService) DeleteVehicleTemplate(ctx context.Context, request *grpc.DeleteVehicleTemplateRequest) (*emptypb.Empty, error) {
	service := commands.NewDeleteVehicleTemplateCommandHandler(s.dbs)

	resp, err := service.Execute(ctx, commands.DeleteVehicleTemplateCommand{
		ID: request.Id,
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}

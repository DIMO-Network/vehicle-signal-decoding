package api

import (
	"context"

	"github.com/volatiletech/null/v8"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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
	mods := make([]qm.QueryMod, 0)

	if request.Make != "" {
		mods = append(mods, models.TemplateVehicleWhere.MakeSlug.EQ(null.StringFrom(request.Make)))
	}

	if request.YearStart > 0 {
		mods = append(mods, models.TemplateVehicleWhere.YearStart.GTE(int(request.YearStart)))
	}

	if request.YearEnd > 0 {
		mods = append(mods, models.TemplateVehicleWhere.YearEnd.LTE(int(request.YearEnd)))
	}

	vehicleTemplates, err := models.TemplateVehicles(mods...).All(ctx, s.dbs().Reader)

	if err != nil {
		return nil, err
	}

	grpcVehicleTemplates := make([]*grpc.VehicleTemplate, 0)

	for _, vehicleTemplate := range vehicleTemplates {

		vt := &grpc.VehicleTemplate{
			Make:      vehicleTemplate.MakeSlug.String,
			YearStart: int32(vehicleTemplate.YearStart),
			YearEnd:   int32(vehicleTemplate.YearEnd),
			Template:  vehicleTemplate.TemplateName,
			Id:        int64(vehicleTemplate.ID),
		}

		for _, model := range vehicleTemplate.ModelWhitelist {
			vt.Models = append(vt.Models, model)
		}

		grpcVehicleTemplates = append(grpcVehicleTemplates, vt)
	}

	return &grpc.GetVehicleTemplatesResponse{
		Templates: grpcVehicleTemplates,
	}, nil
}

func (s *VehicleTemplateService) GetVehicleTemplate(ctx context.Context, request *grpc.GetVehicleTemplateRequest) (*grpc.VehicleTemplate, error) {
	vehicleTemplate, err := models.TemplateVehicles(
		models.TemplateVehicleWhere.ID.EQ(int(request.Id)),
	).One(ctx, s.dbs().Reader)

	if err != nil {
		return nil, err
	}

	vt := &grpc.VehicleTemplate{
		Make:      vehicleTemplate.MakeSlug.String,
		Id:        int64(vehicleTemplate.ID),
		YearStart: int32(vehicleTemplate.YearStart),
		YearEnd:   int32(vehicleTemplate.YearEnd),
		Template:  vehicleTemplate.TemplateName,
		Models:    vehicleTemplate.ModelWhitelist,
	}

	return vt, nil
}

func (s *VehicleTemplateService) CreateVehicleTemplate(ctx context.Context, request *grpc.VehicleTemplate) (*emptypb.Empty, error) {
	vehicleTemplate := &models.TemplateVehicle{
		YearStart:      int(request.YearStart),
		YearEnd:        int(request.YearEnd),
		ModelWhitelist: request.Models,
		TemplateName:   request.Template,
	}
	if request.Make != "" {
		vehicleTemplate.MakeSlug = null.StringFrom(request.Make)
	}

	err := vehicleTemplate.Insert(ctx, s.dbs().Writer, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *VehicleTemplateService) UpdateVehicleTemplate(ctx context.Context, request *grpc.VehicleTemplate) (*emptypb.Empty, error) {
	vehicleTemplate, err := models.TemplateVehicles(
		models.TemplateVehicleWhere.ID.EQ(int(request.Id)),
	).One(ctx, s.dbs().Reader)

	if err != nil {
		return nil, err
	}

	vehicleTemplate.ModelWhitelist = request.Models
	vehicleTemplate.TemplateName = request.Template
	vehicleTemplate.YearStart = int(request.YearStart)
	vehicleTemplate.YearEnd = int(request.YearEnd)

	if request.Make != "" {
		vehicleTemplate.MakeSlug = null.StringFrom(request.Make)
	} else {
		vehicleTemplate.MakeSlug = null.String{}
	}

	_, err = vehicleTemplate.Update(ctx, s.dbs().Writer, boil.Infer())

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

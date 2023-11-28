package api

import (
	"context"
	"fmt"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	grpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
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

	if len(request.Make) > 0 {
		mods = append(mods, models.TemplateVehicleWhere.MakeSlug.EQ(request.Make))
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
			Make:      vehicleTemplate.MakeSlug,
			StartYear: int32(vehicleTemplate.YearStart),
			EndYear:   int32(vehicleTemplate.YearEnd),
			Template:  vehicleTemplate.TemplateName,
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
		models.TemplateVehicleWhere.MakeSlug.EQ(request.Make),
		models.TemplateVehicleWhere.YearStart.EQ(int(request.YearStart)),
		models.TemplateVehicleWhere.YearEnd.EQ(int(request.YearEnd)),
	).One(ctx, s.dbs().Reader)

	if err != nil {
		return nil, err
	}

	vt := &grpc.VehicleTemplate{
		Make:      vehicleTemplate.MakeSlug,
		StartYear: int32(vehicleTemplate.YearStart),
		EndYear:   int32(vehicleTemplate.YearEnd),
		Template:  vehicleTemplate.TemplateName,
		Models:    vehicleTemplate.ModelWhitelist,
	}

	return vt, nil
}

func (s *VehicleTemplateService) CreateVehicleTemplate(ctx context.Context, request *grpc.VehicleTemplate) (*emptypb.Empty, error) {

	fmt.Println(request)

	vehicleTemplate := &models.TemplateVehicle{
		MakeSlug:       request.Make,
		YearStart:      int(request.StartYear),
		YearEnd:        int(request.EndYear),
		ModelWhitelist: request.Models,
		TemplateName:   request.Template,
	}

	err := vehicleTemplate.Insert(ctx, s.dbs().Writer, boil.Infer())

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *VehicleTemplateService) UpdateVehicleTemplate(ctx context.Context, request *grpc.VehicleTemplate) (*emptypb.Empty, error) {

	vehicleTemplate, err := models.TemplateVehicles(
		models.TemplateVehicleWhere.MakeSlug.EQ(request.Make),
		models.TemplateVehicleWhere.YearStart.EQ(int(request.StartYear)),
		models.TemplateVehicleWhere.YearEnd.EQ(int(request.EndYear)),
	).One(ctx, s.dbs().Reader)

	if err != nil {
		return nil, err
	}

	vehicleTemplate.ModelWhitelist = request.Models
	vehicleTemplate.TemplateName = request.Template

	_, err = vehicleTemplate.Update(ctx, s.dbs().Writer, boil.Infer())

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

package queries

import (
	"context"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type GetVehicleTemplatesFilteredQueryHandler struct {
	DBS    func() *db.ReaderWriter
	logger *zerolog.Logger
}

func NewGetVehicleTemplatesFilteredQueryHandler(dbs func() *db.ReaderWriter, logger *zerolog.Logger) GetVehicleTemplatesFilteredQueryHandler {
	return GetVehicleTemplatesFilteredQueryHandler{
		DBS:    dbs,
		logger: logger,
	}
}

type GetVehicleTemplatesFilteredQueryRequest struct {
	MakeSlug  *string
	YearStart *int32
	YearEnd   *int32
	Template  *string
}

func (h GetVehicleTemplatesFilteredQueryHandler) Handle(ctx context.Context, query *GetVehicleTemplatesFilteredQueryRequest) (*grpc.GetVehicleTemplatesResponse, error) {
	mods := make([]qm.QueryMod, 0)

	if query.MakeSlug != nil && *query.MakeSlug != "" {
		mods = append(mods, models.TemplateVehicleWhere.MakeSlug.EQ(null.StringFromPtr(query.MakeSlug)))
	}

	if query.YearStart != nil && *query.YearStart > 0 {
		mods = append(mods, models.TemplateVehicleWhere.YearStart.GTE(int(*query.YearStart)))
	}

	if query.YearEnd != nil && *query.YearEnd > 0 {
		mods = append(mods, models.TemplateVehicleWhere.YearEnd.LTE(int(*query.YearEnd)))
	}

	if query.Template != nil && *query.Template != "" {
		mods = append(mods, models.TemplateVehicleWhere.TemplateName.EQ(*query.Template))
	}

	vehicleTemplates, err := models.TemplateVehicles(mods...).All(ctx, h.DBS().Reader)

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

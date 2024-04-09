package queries

import (
	"context"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
)

type GetVehicleTemplateByIDQueryHandler struct {
	DBS    func() *db.ReaderWriter
	logger *zerolog.Logger
}

func NewGetVehicleTemplateByIDQueryHandler(dbs func() *db.ReaderWriter, logger *zerolog.Logger) GetVehicleTemplateByIDQueryHandler {
	return GetVehicleTemplateByIDQueryHandler{
		DBS:    dbs,
		logger: logger,
	}
}

type GetVehicleTemplateByIDQueryRequest struct {
	ID int64
}

func (h GetVehicleTemplateByIDQueryHandler) Handle(ctx context.Context, query *GetVehicleTemplateByIDQueryRequest) (*grpc.VehicleTemplate, error) {

	vehicleTemplate, err := models.TemplateVehicles(
		models.TemplateVehicleWhere.ID.EQ(int(query.ID)),
	).One(ctx, h.DBS().Reader)

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

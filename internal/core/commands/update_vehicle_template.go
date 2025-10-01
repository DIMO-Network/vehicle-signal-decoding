package commands

import (
	"context"

	"github.com/DIMO-Network/shared/pkg/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UpdateVehicleTemplateCommandHandler struct {
	DBS func() *db.ReaderWriter
}

func NewUpdateVehicleTemplateCommandHandler(dbs func() *db.ReaderWriter) UpdateVehicleTemplateCommandHandler {
	return UpdateVehicleTemplateCommandHandler{DBS: dbs}
}

type UpdateVehicleTemplateCommandRequest struct {
	ID             int64
	MakeSlug       *string
	TemplateName   string
	YearStart      int32
	YearEnd        int32
	ModelWhitelist []string
}

func (h UpdateVehicleTemplateCommandHandler) Execute(ctx context.Context, req *UpdateVehicleTemplateCommandRequest) (*emptypb.Empty, error) {

	vehicleTemplate, err := models.TemplateVehicles(
		models.TemplateVehicleWhere.ID.EQ(int(req.ID)),
	).One(ctx, h.DBS().Reader)

	if err != nil {
		return nil, err
	}

	vehicleTemplate.ModelWhitelist = req.ModelWhitelist
	vehicleTemplate.TemplateName = req.TemplateName
	vehicleTemplate.YearStart = int(req.YearStart)
	vehicleTemplate.YearEnd = int(req.YearEnd)
	vehicleTemplate.MakeSlug = null.StringFromPtr(req.MakeSlug)

	_, err = vehicleTemplate.Update(ctx, h.DBS().Writer, boil.Infer())

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

package commands

import (
	"context"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type CreateVehicleTemplateCommandHandler struct {
	DBS func() *db.ReaderWriter
}

func NewCreateVehicleTemplateCommandHandler(dbs func() *db.ReaderWriter) CreateVehicleTemplateCommandHandler {
	return CreateVehicleTemplateCommandHandler{DBS: dbs}
}

type CreateVehicleTemplateCommandRequest struct {
	MakeSlug       *string
	TemplateName   string
	YearStart      int32
	YearEnd        int32
	ModelWhitelist []string
}

func (h CreateVehicleTemplateCommandHandler) Execute(ctx context.Context, req *CreateVehicleTemplateCommandRequest) (*int, error) {
	vehicleTemplate := &models.TemplateVehicle{
		YearStart:      int(req.YearStart),
		YearEnd:        int(req.YearEnd),
		ModelWhitelist: req.ModelWhitelist,
		TemplateName:   req.TemplateName,
	}

	if req.MakeSlug != nil && *req.MakeSlug != "" {
		vehicleTemplate.MakeSlug = null.StringFromPtr(req.MakeSlug)
	}

	err := vehicleTemplate.Insert(ctx, h.DBS().Writer, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &vehicleTemplate.ID, nil
}

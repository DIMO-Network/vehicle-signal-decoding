package commands

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UpdateTemplateCommandHandler struct {
	DBS func() *db.ReaderWriter
}

func NewUpdateTemplateCommandHandler(dbs func() *db.ReaderWriter) UpdateTemplateCommandHandler {
	return UpdateTemplateCommandHandler{DBS: dbs}
}

type UpdateTemplateCommandRequest struct {
	Name               string
	ParentTemplateName string
	Version            string
	Protocol           string
	Powertrain         string
	DBC                string
	TemplateVehicles   []string
}

type UpdateTemplateCommandResponse struct {
	Name string
}

func (h UpdateTemplateCommandHandler) Execute(ctx context.Context, req *UpdateTemplateCommandRequest) (*UpdateTemplateCommandResponse, error) {

	template, err := models.Templates(models.TemplateWhere.TemplateName.EQ(req.Name)).One(ctx, h.DBS().Reader)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &exceptions.NotFoundError{
				Err: fmt.Errorf("template not found name: %s", req.Name),
			}
		}
		return nil, &exceptions.InternalError{
			Err: err,
		}
	}

	template.TemplateName = req.Name
	template.ParentTemplateName = null.StringFrom(req.ParentTemplateName)
	template.Version = req.Version
	template.Protocol = req.Protocol
	template.Powertrain = req.Powertrain

	if _, err := template.Update(ctx, h.DBS().Writer.DB, boil.Infer()); err != nil {
		return nil, &exceptions.InternalError{
			Err: err,
		}
	}

	return &UpdateTemplateCommandResponse{Name: template.TemplateName}, nil
}

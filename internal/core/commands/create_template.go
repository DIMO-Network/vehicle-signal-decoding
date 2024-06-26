package commands

import (
	"context"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/DIMO-Network/shared/db"
)

type CreateTemplateCommandHandler struct {
	DBS func() *db.ReaderWriter
}

func NewCreateTemplateCommandHandler(dbs func() *db.ReaderWriter) CreateTemplateCommandHandler {
	return CreateTemplateCommandHandler{DBS: dbs}
}

type CreateTemplateCommandRequest struct {
	Name               string
	ParentTemplateName string
	Version            string
	Protocol           string
	Powertrain         string
	DBC                string
	TemplateVehicles   []string
	Comments           *string
}

type CreateTemplateCommandResponse struct {
	Name string
}

func (h CreateTemplateCommandHandler) Execute(ctx context.Context, req *CreateTemplateCommandRequest) (*CreateTemplateCommandResponse, error) {

	exists, err := models.Templates(models.TemplateWhere.TemplateName.EQ(req.Name)).Exists(ctx, h.DBS().Reader)
	if err != nil {
		return nil, &exceptions.InternalError{
			Err: errors.Wrapf(err, "error checking if template exists: %s", req.Name),
		}
	}
	if exists {
		return nil, &exceptions.ConflictError{
			Err: errors.Errorf("template already exists: %s", req.Name),
		}
	}

	template := &models.Template{
		TemplateName:       req.Name,
		ParentTemplateName: null.StringFrom(req.ParentTemplateName),
		Version:            req.Version,
		Protocol:           req.Protocol,
		Powertrain:         req.Powertrain,
		Comments:           null.StringFromPtr(req.Comments),
	}

	err = template.Insert(ctx, h.DBS().Writer, boil.Infer())
	if err != nil {
		return nil, &exceptions.InternalError{
			Err: errors.Wrapf(err, "error inserting template: %s", req.Name),
		}
	}

	return &CreateTemplateCommandResponse{Name: template.TemplateName}, nil
}

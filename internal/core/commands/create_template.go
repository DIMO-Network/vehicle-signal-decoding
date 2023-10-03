package commands

import (
	"context"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	"github.com/pkg/errors"
	"github.com/segmentio/ksuid"
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
	HasDBC             string
	PidsCount          int
	DBC                string
	TemplateVehicles   []string
}

type CreateTemplateCommandResponse struct {
	Name string
}

func (h CreateTemplateCommandHandler) Execute(ctx context.Context, req *CreateTemplateCommandRequest) (*CreateTemplateCommandResponse, error) {
	template := &models.Template{
		TemplateName:       req.Name,
		ParentTemplateName: null.StringFrom(req.ParentTemplateName),
		Version:            req.Version,
		Protocol:           req.Protocol,
		Powertrain:         req.Powertrain,
		//HasDbc:             req.HasDBC,
		//PidsCount:          req.PidsCount,
		//Dbc:                req.DBC,
		//TemplateVehicles:   req.TemplateVehicles,

	}

	template.TemplateName = ksuid.New().String()

	err := template.Insert(ctx, h.DBS().Writer, boil.Infer())
	if err != nil {
		return nil, &exceptions.InternalError{
			Err: errors.Wrapf(err, "error inserting template: %s", req.Name),
		}
	}

	return &CreateTemplateCommandResponse{Name: template.TemplateName}, nil
}

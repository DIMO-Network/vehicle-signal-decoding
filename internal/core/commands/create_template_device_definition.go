package commands

import (
	"context"

	"github.com/DIMO-Network/shared/pkg/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type CreateTemplateDeviceDefinitionCommandHandler struct {
	DBS func() *db.ReaderWriter
}

type CreateTemplateDeviceDefinitionCommand struct {
	DeviceStyleID *string
	TemplateName  string
	DefinitionID  string
}

func NewCreateTemplateDeviceDefinitionCommandHandler(dbs func() *db.ReaderWriter) *CreateTemplateDeviceDefinitionCommandHandler {
	return &CreateTemplateDeviceDefinitionCommandHandler{DBS: dbs}
}

func (h *CreateTemplateDeviceDefinitionCommandHandler) Execute(ctx context.Context, cmd CreateTemplateDeviceDefinitionCommand) (*int64, error) {
	templateDeviceDefinition := &models.TemplateDeviceDefinition{
		TemplateName:  cmd.TemplateName,
		DeviceStyleID: null.StringFromPtr(cmd.DeviceStyleID),
		DefinitionID:  cmd.DefinitionID,
	}

	err := templateDeviceDefinition.Insert(ctx, h.DBS().Writer, boil.Infer())

	if err != nil {
		return nil, err
	}

	return &templateDeviceDefinition.ID, nil
}

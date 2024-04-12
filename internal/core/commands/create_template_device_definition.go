package commands

import (
	"context"
	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CreateTemplateDeviceDefinitionCommandHandler struct {
	DBS func() *db.ReaderWriter
}

type CreateTemplateDeviceDefinitionCommand struct {
	ID                 int64
	DeviceDefinitionID string
	DeviceStyleID      *string
	TemplateName       string
}

func NewCreateTemplateDeviceDefinitionCommandHandler(dbs func() *db.ReaderWriter) *CreateTemplateDeviceDefinitionCommandHandler {
	return &CreateTemplateDeviceDefinitionCommandHandler{DBS: dbs}
}

func (h *CreateTemplateDeviceDefinitionCommandHandler) Execute(ctx context.Context, cmd CreateTemplateDeviceDefinitionCommand) (*emptypb.Empty, error) {
	templateDeviceDefinition := &models.TemplateDeviceDefinition{
		ID:                 cmd.ID,
		DeviceDefinitionID: cmd.DeviceDefinitionID,
		TemplateName:       cmd.TemplateName,
	}

	if cmd.DeviceStyleID != nil && *cmd.DeviceStyleID != "" {
		templateDeviceDefinition.DeviceStyleID = null.StringFromPtr(cmd.DeviceStyleID)
	}

	err := templateDeviceDefinition.Insert(ctx, h.DBS().Writer, boil.Infer())

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
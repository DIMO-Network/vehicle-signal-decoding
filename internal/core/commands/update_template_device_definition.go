package commands

import (
	"context"

	"github.com/DIMO-Network/shared/pkg/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/aarondl/null/v8"
	"github.com/aarondl/sqlboiler/v4/boil"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UpdateTemplateDeviceDefinitionCommandHandler struct {
	DBS func() *db.ReaderWriter
}

type UpdateTemplateDeviceDefinitionCommand struct {
	ID            int64
	DeviceStyleID *string
	TemplateName  string
	DefinitionID  string
}

func NewUpdateTemplateDeviceDefinitionCommandHandler(dbS func() *db.ReaderWriter) *UpdateTemplateDeviceDefinitionCommandHandler {
	return &UpdateTemplateDeviceDefinitionCommandHandler{DBS: dbS}
}

func (h *UpdateTemplateDeviceDefinitionCommandHandler) Execute(ctx context.Context, cmd UpdateTemplateDeviceDefinitionCommand) (*emptypb.Empty, error) {
	templateDeviceDefinition := &models.TemplateDeviceDefinition{
		ID:            cmd.ID,
		TemplateName:  cmd.TemplateName,
		DeviceStyleID: null.StringFromPtr(cmd.DeviceStyleID),
		DefinitionID:  cmd.DefinitionID,
	}

	_, err := templateDeviceDefinition.Update(ctx, h.DBS().Writer, boil.Infer())

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

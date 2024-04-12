package commands

import (
	"context"
	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UpdateTemplateDeviceDefinitionCommandHandler struct {
	DBS func() *db.ReaderWriter
}

type UpdateTemplateDeviceDefinitionCommand struct {
	ID                 int64
	DeviceDefinitionID string
	DeviceStyleID      *string
	TemplateName       string
}

func NewUpdateTemplateDeviceDefinitionCommandHandler(dbS func() *db.ReaderWriter) *UpdateTemplateDeviceDefinitionCommandHandler {
	return &UpdateTemplateDeviceDefinitionCommandHandler{DBS: dbS}
}

func (h *UpdateTemplateDeviceDefinitionCommandHandler) Execute(ctx context.Context, cmd UpdateTemplateDeviceDefinitionCommand) (*emptypb.Empty, error) {
	templateDeviceDefinition := &models.TemplateDeviceDefinition{
		ID:                 cmd.ID,
		DeviceDefinitionID: cmd.DeviceDefinitionID,
		TemplateName:       cmd.TemplateName,
		DeviceStyleID:      null.StringFromPtr(cmd.DeviceStyleID),
	}

	_, err := templateDeviceDefinition.Update(ctx, h.DBS().Writer, boil.Infer())

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

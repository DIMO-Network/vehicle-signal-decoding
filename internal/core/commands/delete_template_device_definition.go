package commands

import (
	"context"
	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DeleteTemplateDeviceDefinitionCommandHandler struct {
	DBS func() *db.ReaderWriter
}

type DeleteTemplateDeviceDefinitionCommand struct {
	ID int64
}

func NewDeleteTemplateDeviceDefinitionCommandHandler(dbs func() *db.ReaderWriter) *DeleteTemplateDeviceDefinitionCommandHandler {
	return &DeleteTemplateDeviceDefinitionCommandHandler{DBS: dbs}
}

func (h *DeleteTemplateDeviceDefinitionCommandHandler) Execute(ctx context.Context, cmd *DeleteTemplateDeviceDefinitionCommand) (*emptypb.Empty, error) {
	templateDeviceDefinition, err := models.FindTemplateDeviceDefinition(ctx, h.DBS().Reader, cmd.ID)

	if err != nil {
		return nil, err
	}

	_, err = templateDeviceDefinition.Delete(ctx, h.DBS().Writer)

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

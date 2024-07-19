package commands

import (
	"context"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DeleteVehicleTemplateCommandHandler struct {
	DBS func() *db.ReaderWriter
}

type DeleteVehicleTemplateCommand struct {
	ID int64
}

func NewDeleteVehicleTemplateCommandHandler(dbs func() *db.ReaderWriter) *DeleteVehicleTemplateCommandHandler {
	return &DeleteVehicleTemplateCommandHandler{DBS: dbs}
}

func (h *DeleteVehicleTemplateCommandHandler) Execute(ctx context.Context, cmd DeleteVehicleTemplateCommand) (*emptypb.Empty, error) {
	vehicleTemplate, err := models.FindTemplateVehicle(ctx, h.DBS().Reader, int(cmd.ID))

	if err != nil {
		return nil, err
	}

	_, err = vehicleTemplate.Delete(ctx, h.DBS().Writer)

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

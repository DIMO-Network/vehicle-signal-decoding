package commands

import (
	"context"

	"github.com/DIMO-Network/shared/pkg/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
)

type DeletePidCommandHandler struct {
	DBS func() *db.ReaderWriter
}

func NewDeletePidCommandHandler(dbs func() *db.ReaderWriter) DeletePidCommandHandler {
	return DeletePidCommandHandler{DBS: dbs}
}

type DeletePidCommandRequest struct {
	ID           int64
	TemplateName string
}

func (h DeletePidCommandHandler) Execute(ctx context.Context, req *DeletePidCommandRequest) error {
	pid, err := models.FindPidConfig(ctx, h.DBS().Reader, req.ID)
	if err != nil {
		return err
	}
	_, err = pid.Delete(ctx, h.DBS().Writer)
	if err != nil {
		return err
	}

	return nil
}

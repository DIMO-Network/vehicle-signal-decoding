package commands

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
)

type UpdateEnableStatusPidCommandHandler struct {
	DBS func() *db.ReaderWriter
}

func NewUpdateEnableStatusPidCommandHandler(dbs func() *db.ReaderWriter) UpdateEnableStatusPidCommandHandler {
	return UpdateEnableStatusPidCommandHandler{DBS: dbs}
}

type UpdateEnableStatusPidCommandRequest struct {
	ID int64
}

func (h UpdateEnableStatusPidCommandHandler) Execute(ctx context.Context, req *UpdateEnableStatusPidCommandRequest) error {
	pid, err := models.FindPidConfig(ctx, h.DBS().Reader, req.ID)
	if err != nil {
		return err
	}

	pid.Enabled = !pid.Enabled

	_, err = pid.Update(ctx, h.DBS().Writer, boil.Infer())

	if err != nil {
		return err
	}

	return nil
}

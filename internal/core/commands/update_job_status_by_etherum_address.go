package commands

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/aarondl/sqlboiler/v4/boil"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	"github.com/pkg/errors"

	"github.com/DIMO-Network/shared/pkg/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
)

type UpdateJobStatusByEtherumAddressCommandHandler struct {
	DBS func() *db.ReaderWriter
}

func NewUpdateJobStatusByEtherumAddressCommandHandler(dbs func() *db.ReaderWriter) UpdateJobStatusByEtherumAddressCommandHandler {
	return UpdateJobStatusByEtherumAddressCommandHandler{DBS: dbs}
}

type UpdateJobStatusCommandRequest struct {
	ID     string
	Status string
}

type UpdateJobStatusCommandResponse struct {
	Status bool
}

func (h UpdateJobStatusByEtherumAddressCommandHandler) Execute(ctx context.Context, command *UpdateJobStatusCommandRequest) (*UpdateJobStatusCommandResponse, error) {

	job, err := models.Jobs(models.JobWhere.ID.EQ(command.ID)).One(ctx, h.DBS().Reader)

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, &exceptions.InternalError{
				Err: err,
			}
		}

		if errors.Is(err, sql.ErrNoRows) {
			return nil, &exceptions.NotFoundError{
				Err: fmt.Errorf("could not find job id: %s", command.ID),
			}
		}
	}

	job.Status = command.Status

	if _, err := job.Update(ctx, h.DBS().Writer.DB, boil.Infer()); err != nil {
		return nil, &exceptions.InternalError{
			Err: err,
		}
	}

	return &UpdateJobStatusCommandResponse{Status: true}, nil
}

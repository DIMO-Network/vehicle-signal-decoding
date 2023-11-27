package commands

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	"github.com/pkg/errors"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
)

type DeleteJobByEtherumAddressCommandHandler struct {
	DBS func() *db.ReaderWriter
}

func NewDeleteJobByEtherumAddressCommandHandler(dbs func() *db.ReaderWriter) DeleteJobByEtherumAddressCommandHandler {
	return DeleteJobByEtherumAddressCommandHandler{DBS: dbs}
}

type DeleteJobCommandRequest struct {
	ID string
}

type DeleteJobCommandResponse struct {
	Status bool
}

func (h DeleteJobByEtherumAddressCommandHandler) Execute(ctx context.Context, command *DeleteJobCommandRequest) (*DeleteJobCommandResponse, error) {

	dbc, err := models.Jobs(models.JobWhere.ID.EQ(command.ID)).One(ctx, h.DBS().Reader)

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

	if _, err := dbc.Delete(ctx, h.DBS().Writer.DB); err != nil {
		return nil, &exceptions.InternalError{
			Err: err,
		}
	}

	return &DeleteJobCommandResponse{Status: true}, nil
}

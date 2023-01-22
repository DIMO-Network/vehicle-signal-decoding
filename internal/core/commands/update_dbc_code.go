package commands

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	"github.com/pkg/errors"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
)

type UpdateDBCCodeCommandHandler struct {
	DBS func() *db.ReaderWriter
}

func NewUpdateDBCCodeCommandHandler(dbs func() *db.ReaderWriter) UpdateDBCCodeCommandHandler {
	return UpdateDBCCodeCommandHandler{DBS: dbs}
}

type UpdateDBCCodeCommandRequest struct {
	ID          string
	Name        string
	DBCContents string
}

type UpdateDBCCodeCommandResponse struct {
	ID string
}

func (h UpdateDBCCodeCommandHandler) Execute(ctx context.Context, command *UpdateDBCCodeCommandRequest) (*UpdateDBCCodeCommandResponse, error) {

	dbc, err := models.DBCCodes(models.DBCCodeWhere.ID.EQ(command.ID)).One(ctx, h.DBS().Reader)

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, &exceptions.InternalError{
				Err: err,
			}
		}

		if errors.Is(err, sql.ErrNoRows) {
			return nil, &exceptions.NotFoundError{
				Err: fmt.Errorf("could not find dbc_code id: %s", command.ID),
			}
		}
	}

	dbc.Name = command.Name
	dbc.DBCContents = command.DBCContents

	if _, err := dbc.Update(ctx, h.DBS().Writer.DB, boil.Infer()); err != nil {
		return nil, &exceptions.InternalError{
			Err: err,
		}
	}

	return &UpdateDBCCodeCommandResponse{ID: dbc.ID}, nil
}

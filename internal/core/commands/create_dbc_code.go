package commands

import (
	"context"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/segmentio/ksuid"
)

type CreateDBCCodeCommandHandler struct {
	DBS func() *db.ReaderWriter
}

func NewCreateDBCCodeCommandHandler(dbs func() *db.ReaderWriter) CreateDBCCodeCommandHandler {
	return CreateDBCCodeCommandHandler{DBS: dbs}
}

type CreateDBCCodeCommandRequest struct {
	Name        string
	DBCContents string
}

type CreateDBCCodeCommandResponse struct {
	ID string
}

func (h CreateDBCCodeCommandHandler) Execute(ctx context.Context, command *CreateDBCCodeCommandRequest) (*CreateDBCCodeCommandResponse, error) {

	dbc := &models.DBCCode{}
	dbc.ID = ksuid.New().String()
	dbc.Name = command.Name
	dbc.DBCContents = command.DBCContents

	err := dbc.Insert(ctx, h.DBS().Writer, boil.Infer())

	if err != nil {
		return nil, &exceptions.InternalError{Err: errors.Wrapf(err, "error inserting dbc_code: %s", command.Name)}
	}

	return &CreateDBCCodeCommandResponse{ID: dbc.ID}, nil
}

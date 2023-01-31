package commands

import (
	"context"

	"github.com/volatiletech/null/v8"

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
	Name             string
	DBCContents      string
	Header           int
	Trigger          string
	RecordingEnabled bool
	MaxSampleSize    int32
}

type CreateDBCCodeCommandResponse struct {
	ID string
}

func (h CreateDBCCodeCommandHandler) Execute(ctx context.Context, command *CreateDBCCodeCommandRequest) (*CreateDBCCodeCommandResponse, error) {

	dbc := &models.DBCCode{}
	dbc.ID = ksuid.New().String()
	dbc.Name = command.Name
	dbc.DBCContents = null.StringFrom(command.DBCContents)
	dbc.Header = null.IntFrom(command.Header)
	dbc.Trigger = command.Trigger
	dbc.RecordingEnabled = command.RecordingEnabled
	dbc.MaxSampleSize = int(command.MaxSampleSize)

	err := dbc.Insert(ctx, h.DBS().Writer, boil.Infer())

	if err != nil {
		return nil, &exceptions.InternalError{Err: errors.Wrapf(err, "error inserting dbc_code: %s", command.Name)}
	}

	return &CreateDBCCodeCommandResponse{ID: dbc.ID}, nil
}

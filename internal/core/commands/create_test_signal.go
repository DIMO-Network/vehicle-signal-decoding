package commands

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/segmentio/ksuid"
)

type CreateTestSignalCommandHandler struct {
	DBS func() *db.ReaderWriter
}

func NewCreateTestSignalCommandHandler(dbs func() *db.ReaderWriter) CreateTestSignalCommandHandler {
	return CreateTestSignalCommandHandler{DBS: dbs}
}

type CreateTestSignalCommandRequest struct {
	DeviceDefinitionID string
	DBCCodesID         string
	UserDeviceID       string
	AutoPIUnitID       string
	Value              string
	Approved           bool
}

type CreateTestSignalCommandResponse struct {
	ID string
}

func (h CreateTestSignalCommandHandler) Execute(ctx context.Context, command *CreateTestSignalCommandRequest) (*CreateTestSignalCommandResponse, error) {

	dbc, err := models.DBCCodes(models.DBCCodeWhere.ID.EQ(command.DBCCodesID)).One(ctx, h.DBS().Reader)

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, &exceptions.InternalError{
				Err: err,
			}
		}

		if errors.Is(err, sql.ErrNoRows) {
			return nil, &exceptions.NotFoundError{
				Err: fmt.Errorf("could not find dbc_code id: %s", command.DBCCodesID),
			}
		}
	}

	test := models.TestSignal{}
	test.ID = ksuid.New().String()
	test.DeviceDefinitionID = command.DeviceDefinitionID
	test.UserDeviceID = command.UserDeviceID
	test.DBCCodesID = dbc.ID
	test.AutopiUnitID = command.AutoPIUnitID
	test.Value = command.Value
	test.Approved = command.Approved

	err = test.Insert(ctx, h.DBS().Writer, boil.Infer())

	if err != nil {
		return nil, &exceptions.InternalError{Err: errors.Wrapf(err, "error inserting test signal: %s", command.Value)}
	}

	return &CreateTestSignalCommandResponse{ID: test.ID}, nil
}

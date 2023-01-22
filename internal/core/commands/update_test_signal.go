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

type UpdateTestSignalCommandHandler struct {
	DBS func() *db.ReaderWriter
}

func NewUpdateTestSignalCommandHandler(dbs func() *db.ReaderWriter) UpdateTestSignalCommandHandler {
	return UpdateTestSignalCommandHandler{DBS: dbs}
}

type UpdateTestSignalCommandRequest struct {
	ID                 string
	DeviceDefinitionID string
	DBCCodesID         string
	UserDeviceID       string
	Trigger            string
	Name               string
	Value              string
}

type UpdateTestSignalCommandResponse struct {
	ID string
}

func (h UpdateTestSignalCommandHandler) Execute(ctx context.Context, command *UpdateTestSignalCommandRequest) (*UpdateTestSignalCommandResponse, error) {

	test, err := models.TestSignals(models.TestSignalWhere.ID.EQ(command.ID)).One(ctx, h.DBS().Reader)

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

	test.SignalName = command.Name
	test.DeviceDefinitionID = command.DeviceDefinitionID
	test.DBCCodesID = command.DBCCodesID
	test.Trigger = command.Trigger
	test.Value = command.Value

	if _, err := test.Update(ctx, h.DBS().Writer.DB, boil.Infer()); err != nil {
		return nil, &exceptions.InternalError{
			Err: err,
		}
	}

	return &UpdateTestSignalCommandResponse{ID: test.ID}, nil
}

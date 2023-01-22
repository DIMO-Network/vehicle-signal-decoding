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
	Trigger            string
	Name               string
	Value              string
}

type CreateTestSignalCommandResponse struct {
	ID string
}

func (h CreateTestSignalCommandHandler) Execute(ctx context.Context, command *CreateTestSignalCommandRequest) (*CreateTestSignalCommandResponse, error) {

	test := models.TestSignal{}
	test.ID = ksuid.New().String()
	test.SignalName = command.Name
	test.DeviceDefinitionID = command.DeviceDefinitionID
	test.DBCCodesID = command.DBCCodesID
	test.Trigger = command.Trigger
	test.Value = command.Value

	err := test.Insert(ctx, h.DBS().Writer, boil.Infer())

	if err != nil {
		return nil, &exceptions.InternalError{Err: errors.Wrapf(err, "error inserting test signal: %s", command.Name)}
	}

	return &CreateTestSignalCommandResponse{ID: test.ID}, nil
}

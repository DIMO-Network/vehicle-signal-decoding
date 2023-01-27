package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	"github.com/pkg/errors"

	"github.com/DIMO-Network/shared/db"
	p_grpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
)

type GetTestSignalByIDQueryHandler struct {
	DBS    func() *db.ReaderWriter
	logger *zerolog.Logger
}

func NewGetTestSignalByIDQueryHandler(dbs func() *db.ReaderWriter, logger *zerolog.Logger) GetTestSignalByIDQueryHandler {
	return GetTestSignalByIDQueryHandler{
		DBS:    dbs,
		logger: logger,
	}
}

type GetTestSignalByIDQueryRequest struct {
	ID string
}

func (h GetTestSignalByIDQueryHandler) Handle(ctx context.Context, query *GetTestSignalByIDQueryRequest) (*p_grpc.GetTestSignalResponse, error) {

	item, err := models.TestSignals(models.TestSignalWhere.ID.EQ(query.ID)).One(ctx, h.DBS().Reader)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &exceptions.NotFoundError{
				Err: fmt.Errorf("could not find dbc_code id: %s", query.ID),
			}
		}

		return nil, &exceptions.InternalError{
			Err: fmt.Errorf("failed to get dbc_codes"),
		}
	}

	result := &p_grpc.GetTestSignalResponse{
		Id:                 item.ID,
		Name:               item.SignalName,
		UserDeviceId:       item.UserDeviceID,
		DeviceDefinitionId: item.DeviceDefinitionID,
		DbcCodesId:         item.DBCCodesID,
		Trigger:            item.Trigger,
		Value:              item.Value,
	}

	return result, nil
}

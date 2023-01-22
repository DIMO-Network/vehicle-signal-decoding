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

type GetTestSignalByIdQueryHandler struct {
	DBS    func() *db.ReaderWriter
	logger *zerolog.Logger
}

func NewGetTestSignalByIdQueryHandler(dbs func() *db.ReaderWriter, logger *zerolog.Logger) GetTestSignalByIdQueryHandler {
	return GetTestSignalByIdQueryHandler{
		DBS:    dbs,
		logger: logger,
	}
}

type GetTestSignalByIdQueryRequest struct {
	ID string
}

func (h GetTestSignalByIdQueryHandler) Handle(ctx context.Context, query *GetTestSignalByIdQueryRequest) (*p_grpc.GetTestSignalResponse, error) {

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

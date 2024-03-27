package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/common"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	"github.com/pkg/errors"

	"github.com/DIMO-Network/shared/db"
	pgrpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
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

func (h GetTestSignalByIDQueryHandler) Handle(ctx context.Context, query *GetTestSignalByIDQueryRequest) (*pgrpc.GetTestSignalResponse, error) {

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

	result := &pgrpc.GetTestSignalResponse{
		Id:                 item.ID,
		UserDeviceId:       item.UserDeviceID,
		DeviceDefinitionId: item.DeviceDefinitionID,
		DbcCodesId:         item.DBCCodesID,
		Value:              item.Value,
		AutopiUnitId:       item.AutopiUnitID,
		Approved:           item.Approved,
		Signals:            string(common.JSONOrDefault(item.Signals)),
	}

	return result, nil
}

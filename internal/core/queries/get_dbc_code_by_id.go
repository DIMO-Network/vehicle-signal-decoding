package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	"github.com/pkg/errors"

	"github.com/DIMO-Network/shared/pkg/db"
	pgrpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
)

type GetDBCCodeByIDQueryHandler struct {
	DBS    func() *db.ReaderWriter
	logger *zerolog.Logger
}

func NewGetDBCCodeByIDQueryHandler(dbs func() *db.ReaderWriter, logger *zerolog.Logger) GetDBCCodeByIDQueryHandler {
	return GetDBCCodeByIDQueryHandler{
		DBS:    dbs,
		logger: logger,
	}
}

type GetDBCCodeByIDQueryRequest struct {
	ID string
}

func (h GetDBCCodeByIDQueryHandler) Handle(ctx context.Context, query *GetDBCCodeByIDQueryRequest) (*pgrpc.GetDBCCodeResponse, error) {

	item, err := models.DBCCodes(models.DBCCodeWhere.ID.EQ(query.ID)).One(ctx, h.DBS().Reader)
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

	result := &pgrpc.GetDBCCodeResponse{
		Id:               item.ID,
		Name:             item.Name,
		DbcContents:      item.DBCContents.String,
		Header:           int32(item.Header.Int),
		Trigger:          item.Trigger,
		MaxSampleSize:    int32(item.MaxSampleSize),
		RecordingEnabled: item.RecordingEnabled,
	}

	return result, nil
}

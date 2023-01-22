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

type GetDBCCodeByIdQueryHandler struct {
	DBS    func() *db.ReaderWriter
	logger *zerolog.Logger
}

func NewGetDBCCodeByIdQueryHandler(dbs func() *db.ReaderWriter, logger *zerolog.Logger) GetDBCCodeByIdQueryHandler {
	return GetDBCCodeByIdQueryHandler{
		DBS:    dbs,
		logger: logger,
	}
}

type GetDBCCodeByIdQueryRequest struct {
	ID string
}

func (h GetDBCCodeByIdQueryHandler) Handle(ctx context.Context, query *GetDBCCodeByIdQueryRequest) (*p_grpc.GetDBCCodeResponse, error) {

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

	result := &p_grpc.GetDBCCodeResponse{
		Id:          item.ID,
		Name:        item.Name,
		DbcContents: item.DBCContents,
	}

	return result, nil
}

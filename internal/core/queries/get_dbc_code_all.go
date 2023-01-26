package queries

import (
	"context"
	"fmt"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"

	"github.com/DIMO-Network/shared/db"
	p_grpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
)

type GetDBCCodeAllQueryHandler struct {
	DBS    func() *db.ReaderWriter
	logger *zerolog.Logger
}

func NewGetDBCCodeAllQueryHandler(dbs func() *db.ReaderWriter, logger *zerolog.Logger) GetDBCCodeAllQueryHandler {
	return GetDBCCodeAllQueryHandler{
		DBS:    dbs,
		logger: logger,
	}
}

type GetDBCCodeAllQueryRequest struct {
}

func (h GetDBCCodeAllQueryHandler) Handle(ctx context.Context, query *GetDBCCodeAllQueryRequest) (*p_grpc.GetDBCCodeListResponse, error) {

	all, err := models.DBCCodes().All(ctx, h.DBS().Reader)
	if err != nil {
		return nil, &exceptions.InternalError{
			Err: fmt.Errorf("failed to get dbc_codes"),
		}
	}

	result := &p_grpc.GetDBCCodeListResponse{}

	for _, item := range all {
		result.Items = append(result.Items, &p_grpc.GetDBCCodeResponse{
			Id:          item.ID,
			Name:        item.Name,
			DbcContents: item.DBCContents,
		})
	}

	return result, nil
}

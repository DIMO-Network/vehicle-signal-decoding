package queries

import (
	"context"

	"github.com/DIMO-Network/shared/db"
	"github.com/rs/zerolog"
)

type BaseQueryHandler struct {
	dbs    func() *db.ReaderWriter
	logger *zerolog.Logger
}

func NewBaseQueryHandler(dbs func() *db.ReaderWriter, logger *zerolog.Logger) BaseQueryHandler {
	return BaseQueryHandler{
		dbs:    dbs,
		logger: logger,
	}
}

type BaseQueryRequest struct {
	Input string
}

type BaseQueryResponse struct {
	Input string
}

func (h BaseQueryHandler) Handle(ctx context.Context, query *BaseQueryRequest) (*BaseQueryResponse, error) {
	return &BaseQueryResponse{}, nil
}

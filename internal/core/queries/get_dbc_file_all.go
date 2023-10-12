package queries

import (
	"context"
	"fmt"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	grpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
)

type GetDbcAllQueryHandler struct {
	DBS    func() *db.ReaderWriter
	logger *zerolog.Logger
}

func NewGetDbcAllQueryHandler(dbs func() *db.ReaderWriter, logger *zerolog.Logger) GetDbcAllQueryHandler {
	return GetDbcAllQueryHandler{
		DBS:    dbs,
		logger: logger,
	}
}

type GetDbcAllQueryRequest struct {
	TemplateName string
}

func (h GetDbcAllQueryHandler) Handle(ctx context.Context, _ *GetDbcAllQueryRequest) (*grpc.GetDbcListResponse, error) {

	all, err := models.DBCFiles().All(ctx, h.DBS().Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to get DBC File: %w", err)
	}

	dbcFileSummaries := make([]*grpc.DbcSummary, 0, len(all))

	for _, item := range all {
		dbcFileSummaries = append(dbcFileSummaries, &grpc.DbcSummary{
			TemplateName: item.TemplateName,
		})
	}

	result := &grpc.GetDbcListResponse{
		Dbc: dbcFileSummaries,
	}

	return result, nil
}

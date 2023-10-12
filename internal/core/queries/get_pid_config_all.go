package queries

import (
	"context"
	"fmt"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	grpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
)

type GetPidAllQueryHandler struct {
	DBS    func() *db.ReaderWriter
	logger *zerolog.Logger
}

func NewGetPidAllQueryHandler(dbs func() *db.ReaderWriter, logger *zerolog.Logger) GetPidAllQueryHandler {
	return GetPidAllQueryHandler{
		DBS:    dbs,
		logger: logger,
	}
}

type GetPidAllQueryRequest struct {
	ID int64
}

func (h GetPidAllQueryHandler) Handle(ctx context.Context, _ *GetPidAllQueryRequest) (*grpc.GetPidListResponse, error) {

	all, err := models.PidConfigs().All(ctx, h.DBS().Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to get PidConfigs: %w", err)
	}
	pidSummaries := make([]*grpc.PidSummary, 0, len(all))
	for _, item := range all {
		pidSummaries = append(pidSummaries, &grpc.PidSummary{
			Id:              item.ID,
			TemplateName:    item.TemplateName,
			Header:          item.Header,
			Mode:            item.Mode,
			Pid:             item.Pid,
			Formula:         item.Formula,
			IntervalSeconds: int32(item.IntervalSeconds),
			Protocol:        item.Protocol,
			SignalName:      item.SignalName,
		})
	}

	result := &grpc.GetPidListResponse{
		Pid: pidSummaries,
	}

	return result, nil
}

package queries

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	grpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetPidByIDQueryHandler struct {
	DBS    func() *db.ReaderWriter
	logger *zerolog.Logger
}

func NewGetPidByIDQueryHandler(dbs func() *db.ReaderWriter, logger *zerolog.Logger) GetPidByIDQueryHandler {
	return GetPidByIDQueryHandler{
		DBS:    dbs,
		logger: logger,
	}
}

type GetPidByIDQueryRequest struct {
	ID string
}

func (h GetPidByIDQueryHandler) Handle(ctx context.Context, query *GetPidByIDQueryRequest) (*grpc.GetPidByIDResponse, error) {

	item, err := models.PidConfigs(models.PidConfigWhere.TemplateName.EQ(query.ID)).One(ctx, h.DBS().Reader)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("template not found id: %s", query.ID) // You may replace it with your custom NotFoundError
		}

		return nil, fmt.Errorf("failed to get template")
	}

	result := &grpc.GetPidByIDResponse{
		Pid: &grpc.PidConfig{
			TemplateName:    item.TemplateName,
			Header:          item.Header,
			Mode:            item.Mode,
			Pid:             item.Pid,
			Formula:         item.Formula,
			IntervalSeconds: int32(item.IntervalSeconds),
			Protocol:        item.Protocol,
			SignalName:      item.SignalName,
			CreatedAt:       timestamppb.New(item.CreatedAt),
			UpdatedAt:       timestamppb.New(item.UpdatedAt),
		},
	}

	return result, nil
}

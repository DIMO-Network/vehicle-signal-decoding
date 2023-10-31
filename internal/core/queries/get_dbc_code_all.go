package queries

import (
	"context"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"google.golang.org/protobuf/types/known/timestamppb"

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

func (h GetDBCCodeAllQueryHandler) Handle(ctx context.Context, _ *GetDBCCodeAllQueryRequest) (*p_grpc.GetDBCCodeListResponse, error) {

	all, err := models.DBCCodes(qm.OrderBy("created_at desc"), qm.Limit(100)).All(ctx, h.DBS().Reader)
	if err != nil {
		return nil, &exceptions.InternalError{
			Err: fmt.Errorf("failed to get dbc_codes"),
		}
	}

	result := &p_grpc.GetDBCCodeListResponse{}

	for _, item := range all {
		result.Items = append(result.Items, &p_grpc.GetDBCCodeResponse{
			Id:               item.ID,
			Name:             item.Name,
			DbcContents:      item.DBCContents.String,
			Header:           int32(item.Header.Int),
			Trigger:          item.Trigger,
			MaxSampleSize:    int32(item.MaxSampleSize),
			RecordingEnabled: item.RecordingEnabled,
			CreatedAt:        timestamppb.New(item.CreatedAt),
			UpdatedAt:        timestamppb.New(item.UpdatedAt),
		})
	}

	return result, nil
}

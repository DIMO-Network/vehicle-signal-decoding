package queries

import (
	"context"
	"fmt"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/ethutil"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/DIMO-Network/shared/pkg/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	pgrpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
)

type GetJobByEthereumAddressQueryHandler struct {
	DBS    func() *db.ReaderWriter
	logger *zerolog.Logger
}

func NewGetJobByEthereumAddressQueryHandler(dbs func() *db.ReaderWriter, logger *zerolog.Logger) GetJobByEthereumAddressQueryHandler {
	return GetJobByEthereumAddressQueryHandler{
		DBS:    dbs,
		logger: logger,
	}
}

type GetJobByyEthereumAddressQueryRequest struct {
	EtherumAddress string
}

func (h GetJobByEthereumAddressQueryHandler) Handle(ctx context.Context, query *GetJobByyEthereumAddressQueryRequest) (*pgrpc.GetJobsByEtherumAddressResponse, error) {

	ethAddrBytes, err := ethutil.ResolveEtherumAddressFromString(query.EtherumAddress)
	if err != nil {
		return nil, &exceptions.ValidationError{
			Err: fmt.Errorf("invalid ethereum address: %w", err),
		}
	}

	jobs, err := models.Jobs(models.JobWhere.DeviceEthereumAddress.EQ(ethAddrBytes)).All(ctx, h.DBS().Reader)
	if err != nil {
		return nil, &exceptions.InternalError{
			Err: fmt.Errorf("failed to get jobs"),
		}
	}

	result := &pgrpc.GetJobsByEtherumAddressResponse{}

	for _, item := range jobs {
		jobItem := &pgrpc.GetJobsByEtherumAddressItemResponse{
			Id:        item.ID,
			Command:   item.Command,
			Status:    item.Status,
			CreatedAt: timestamppb.New(item.CreatedAt),
		}

		if item.LastExecution.Valid {
			jobItem.LastExecution = timestamppb.New(item.LastExecution.Time)
		}

		result.Items = append(result.Items, jobItem)
	}

	return result, nil
}

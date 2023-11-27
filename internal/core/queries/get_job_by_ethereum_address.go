package queries

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/common"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	p_grpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
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

func (h GetJobByEthereumAddressQueryHandler) Handle(ctx context.Context, query *GetJobByyEthereumAddressQueryRequest) (*p_grpc.GetJobByEtherumAddressResponse, error) {

	ethAddrBytes, err := common.ResolveEtherumAddressFromString(query.EtherumAddress)
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

	result := &p_grpc.GetJobByEtherumAddressResponse{}

	for _, item := range jobs {
		result.Items = append(result.Items, &p_grpc.GetJobByEtherumAddressItemResponse{
			Id:            item.ID,
			Command:       item.Command,
			Status:        item.Status,
			CreatedAt:     timestamppb.New(item.CreatedAt),
			LastExecution: timestamppb.New(item.LastExecution),
		})
	}

	return result, nil
}

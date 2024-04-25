package queries

import (
	"context"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetAftermarketDeviceTemplateByEthereumAddress struct {
	DBS func() *db.ReaderWriter
}

func NewGetAftermarketDeviceTemplateByEthereumAddress(dbs func() *db.ReaderWriter) *GetAftermarketDeviceTemplateByEthereumAddress {
	return &GetAftermarketDeviceTemplateByEthereumAddress{DBS: dbs}
}

func (h *GetAftermarketDeviceTemplateByEthereumAddress) Handle(ctx context.Context, ethAddress common.Address) (*grpc.AftermarketDeviceTemplate, error) {
	aftermarketDeviceTemplate, err := models.AftermarketDeviceToTemplates(models.AftermarketDeviceToTemplateWhere.AftermarketDeviceEthereumAddress.EQ(ethAddress.Bytes())).One(ctx, h.DBS().Reader)

	if err != nil {
		return nil, err
	}

	response := &grpc.AftermarketDeviceTemplate{
		EthereumAddress: aftermarketDeviceTemplate.AftermarketDeviceEthereumAddress,
		TemplateName:    aftermarketDeviceTemplate.TemplateName,
		CreatedAt:       timestamppb.New(aftermarketDeviceTemplate.CreatedAt),
		UpdatedAt:       timestamppb.New(aftermarketDeviceTemplate.UpdatedAt),
	}

	return response, nil
}

package queries

import (
	"context"

	"github.com/DIMO-Network/shared/pkg/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetAftermarketDeviceTemplateAll struct {
	DBS func() *db.ReaderWriter
}

func NewGetAftermarketDeviceTemplateAll(dbs func() *db.ReaderWriter) *GetAftermarketDeviceTemplateAll {
	return &GetAftermarketDeviceTemplateAll{DBS: dbs}
}

func (h *GetAftermarketDeviceTemplateAll) Handle(ctx context.Context) (*grpc.AftermarketDeviceTemplates, error) {
	aftermarketDeviceTemplate, err := models.AftermarketDeviceToTemplates().All(ctx, h.DBS().Reader)

	if err != nil {
		return nil, err
	}

	aftermarketDeviceTemplates := make([]*grpc.AftermarketDeviceTemplate, len(aftermarketDeviceTemplate))

	for i, template := range aftermarketDeviceTemplate {
		aftermarketDeviceTemplates[i] = &grpc.AftermarketDeviceTemplate{
			EthereumAddress: template.AftermarketDeviceEthereumAddress,
			TemplateName:    template.TemplateName,
			CreatedAt:       timestamppb.New(template.CreatedAt),
			UpdatedAt:       timestamppb.New(template.UpdatedAt),
		}
	}

	response := &grpc.AftermarketDeviceTemplates{
		Items: aftermarketDeviceTemplates,
	}

	return response, nil
}

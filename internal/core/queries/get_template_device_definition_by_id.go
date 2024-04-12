package queries

import (
	"context"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetTemplateDeviceDefinitionByIDQuery struct {
	ID int64
}

type GetTemplateDeviceDefinitionByIDQueryHandler struct {
	DBS func() *db.ReaderWriter
}

func NewGetTemplateDeviceDefinitionByIDQueryHandler(dbs func() *db.ReaderWriter) *GetTemplateDeviceDefinitionByIDQueryHandler {
	return &GetTemplateDeviceDefinitionByIDQueryHandler{DBS: dbs}
}

func (h *GetTemplateDeviceDefinitionByIDQueryHandler) Handle(ctx context.Context, q GetTemplateDeviceDefinitionByIDQuery) (*grpc.TemplateDeviceDefinition, error) {
	templateDeviceDefinition, err := models.TemplateDeviceDefinitions(
		models.TemplateDeviceDefinitionWhere.ID.EQ(q.ID),
	).One(ctx, h.DBS().Reader)

	if err != nil {
		return nil, err
	}

	response := &grpc.TemplateDeviceDefinition{
		Id:                 templateDeviceDefinition.ID,
		DeviceDefinitionId: templateDeviceDefinition.DeviceDefinitionID,
		DeviceStyleId:      templateDeviceDefinition.DeviceStyleID.Ptr(),
		TemplateName:       templateDeviceDefinition.TemplateName,
		CreatedAt:          timestamppb.New(templateDeviceDefinition.CreatedAt),
		UpdatedAt:          timestamppb.New(templateDeviceDefinition.UpdatedAt),
	}

	return response, nil
}

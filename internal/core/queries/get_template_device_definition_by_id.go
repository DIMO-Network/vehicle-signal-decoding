package queries

import (
	"context"
	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetTemplateDeviceDefinitionByIdQuery struct {
	ID int64
}

type GetTemplateDeviceDefinitionByIdQueryHandler struct {
	DBS func() *db.ReaderWriter
}

func NewGetTemplateDeviceDefinitionByIdQueryHandler(dbs func() *db.ReaderWriter) *GetTemplateDeviceDefinitionByIdQueryHandler {
	return &GetTemplateDeviceDefinitionByIdQueryHandler{DBS: dbs}
}

func (h *GetTemplateDeviceDefinitionByIdQueryHandler) Handle(ctx context.Context, q GetTemplateDeviceDefinitionByIdQuery) (*grpc.TemplateDeviceDefinition, error) {
	templateDeviceDefinition, err := models.TemplateDeviceDefinitions(
		models.TemplateDeviceDefinitionWhere.ID.EQ(q.ID),
	).One(ctx, h.DBS().Reader)

	if err != nil {
		return nil, err
	}

	response := &grpc.TemplateDeviceDefinition{
		Id:                 templateDeviceDefinition.ID,
		DeviceDefinitionId: templateDeviceDefinition.DeviceDefinitionID,
		DeviceStyleId:      templateDeviceDefinition.DeviceStyleID.String,
		TemplateName:       templateDeviceDefinition.TemplateName,
		CreatedAt:          timestamppb.New(templateDeviceDefinition.CreatedAt),
		UpdatedAt:          timestamppb.New(templateDeviceDefinition.UpdatedAt),
	}

	return response, nil
}

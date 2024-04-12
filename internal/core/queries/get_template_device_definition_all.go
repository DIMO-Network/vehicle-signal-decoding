package queries

import (
	"context"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetTemplateDeviceDefinitionAllQueryHandler struct {
	DBS func() *db.ReaderWriter
}

type GetTemplateDeviceDefinitionAllQuery struct {
}

func NewGetTemplateDeviceDefinitionAllQueryHandler(dbs func() *db.ReaderWriter) *GetTemplateDeviceDefinitionAllQueryHandler {
	return &GetTemplateDeviceDefinitionAllQueryHandler{DBS: dbs}
}

func (h *GetTemplateDeviceDefinitionAllQueryHandler) Handle(ctx context.Context, _ GetTemplateDeviceDefinitionAllQuery) (*grpc.GetTemplateDeviceDefinitionResponse, error) {
	templateDeviceDefinitions, err := models.TemplateDeviceDefinitions(
		qm.OrderBy("template_name, update_at asc"),
	).All(ctx, h.DBS().Reader)

	if err != nil {
		return nil, err
	}

	templateDd := make([]*grpc.TemplateDeviceDefinition, len(templateDeviceDefinitions))

	for i, templateDeviceDefinition := range templateDeviceDefinitions {
		templateDd[i] = &grpc.TemplateDeviceDefinition{
			Id:                 templateDeviceDefinition.ID,
			DeviceDefinitionId: templateDeviceDefinition.DeviceDefinitionID,
			DeviceStyleId:      templateDeviceDefinition.DeviceStyleID.Ptr(),
			TemplateName:       templateDeviceDefinition.TemplateName,
			CreatedAt:          timestamppb.New(templateDeviceDefinition.CreatedAt),
			UpdatedAt:          timestamppb.New(templateDeviceDefinition.UpdatedAt),
		}
	}

	response := &grpc.GetTemplateDeviceDefinitionResponse{
		Items: templateDd,
	}

	return response, nil
}

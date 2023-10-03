package queries

import (
	"context"
	"fmt"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	grpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
)

type GetTemplatesAllQueryHandler struct {
	DBS    func() *db.ReaderWriter
	logger *zerolog.Logger
}

func NewGetTemplatesAllQueryHandler(dbs func() *db.ReaderWriter, logger *zerolog.Logger) GetTemplatesAllQueryHandler {
	return GetTemplatesAllQueryHandler{
		DBS:    dbs,
		logger: logger,
	}
}

type GetTemplatesAllQueryRequest struct {
	Protocol   string
	Powertrain string
}

func (h GetTemplatesAllQueryHandler) Handle(ctx context.Context, _ *GetTemplatesAllQueryRequest) (*grpc.GetTemplateListResponse, error) {
	all, err := models.Templates().All(ctx, h.DBS().Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to get templates: %w", err)
	}

	result := &grpc.GetTemplateListResponse{}

	for _, item := range all {
		result.Templates = append(result.Templates, &grpc.TemplateSummary{
			Name:       item.TemplateName,
			Version:    item.Version,
			Protocol:   item.Protocol,
			Powertrain: item.Powertrain,
			HasDbc:     "", //need help
			PidsCount:  1,  //need help
		})
	}
	return result, nil
}

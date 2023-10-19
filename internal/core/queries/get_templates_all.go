package queries

import (
	"context"
	"fmt"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	"github.com/pkg/errors"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

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
	Protocol   *string
	Powertrain *string
}

func (h GetTemplatesAllQueryHandler) Handle(ctx context.Context, query *GetTemplatesAllQueryRequest) (*grpc.GetTemplateListResponse, error) {
	var mods []qm.QueryMod

	if query.Protocol != nil && len(*query.Protocol) > 0 {
		validProtocols := map[string]struct{}{
			models.CanProtocolTypeCAN11_500: {},
			models.CanProtocolTypeCAN29_500: {},
		}
		if _, isValid := validProtocols[*query.Protocol]; !isValid {
			return nil, &exceptions.ValidationError{
				Err: errors.Errorf("invalid protocol: %s", *query.Protocol),
			}
		}
		mods = append(mods,
			models.TemplateWhere.Protocol.EQ(*query.Protocol))
	}

	if query.Powertrain != nil && len(*query.Powertrain) > 0 {
		mods = append(mods,
			models.TemplateWhere.Powertrain.EQ(*query.Powertrain))
	}
	// future optimization, use raw sql. note that we're pulling in the entire relationship list just to use the count below and whether it as a dbc file present
	mods = append(mods, qm.Load(models.TemplateRels.TemplateNameDBCFile), qm.Load(models.TemplateRels.TemplateNamePidConfigs))

	all, err := models.Templates(mods...).All(ctx, h.DBS().Reader)
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
			HasDbc:     item.R.GetTemplateNameDBCFile().DBCFile,
			PidsCount:  int32(len(item.R.GetTemplateNamePidConfigs())),
		})
	}
	return result, nil
}

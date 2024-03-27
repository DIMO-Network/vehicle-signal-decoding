package queries

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
)

type GetDbcByTemplateNameQueryHandler struct {
	DBS    func() *db.ReaderWriter
	logger *zerolog.Logger
}

func NewGetDbcByTemplateNameQueryHandler(dbs func() *db.ReaderWriter, logger *zerolog.Logger) GetDbcByTemplateNameQueryHandler {
	return GetDbcByTemplateNameQueryHandler{
		DBS:    dbs,
		logger: logger,
	}
}

type GetDbcByTemplateNameQueryRequest struct {
	TemplateName string
}

func (h GetDbcByTemplateNameQueryHandler) Handle(ctx context.Context, query *GetDbcByTemplateNameQueryRequest) (*grpc.GetDbcByTemplateNameResponse, error) {

	item, err := models.DBCFiles(models.DBCFileWhere.TemplateName.EQ(query.TemplateName)).One(ctx, h.DBS().Reader)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.logger.Info().Msgf("failed to get DBC File by template name: %s", err)
			return &grpc.GetDbcByTemplateNameResponse{Dbc: &grpc.DbcConfig{}}, nil
		}
		return nil, fmt.Errorf("failed to get DBC File by template name: %s", err)
	}

	result := &grpc.GetDbcByTemplateNameResponse{
		Dbc: &grpc.DbcConfig{
			TemplateName: item.TemplateName,
			DbcFile:      item.DBCFile,
		},
	}

	return result, nil
}

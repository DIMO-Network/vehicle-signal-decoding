package queries

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	grpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetTemplateByNameQueryHandler struct {
	DBS    func() *db.ReaderWriter
	logger *zerolog.Logger
}

func NewGetTemplateByNameQueryHandler(dbs func() *db.ReaderWriter, logger *zerolog.Logger) GetTemplateByNameQueryHandler {
	return GetTemplateByNameQueryHandler{
		DBS:    dbs,
		logger: logger,
	}
}

type GetTemplateByNameQueryRequest struct {
	Name string
}

func (h GetTemplateByNameQueryHandler) Handle(ctx context.Context, query *GetTemplateByNameQueryRequest) (*grpc.GetTemplateByNameResponse, error) {

	item, err := models.Templates(
		models.TemplateWhere.TemplateName.EQ(query.Name),
		qm.Load(models.TemplateRels.TemplateNameDBCFile),
		qm.Load(models.TemplateRels.TemplateNamePidConfigs),
	).One(ctx, h.DBS().Reader)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("template not found name: %s", query.Name)
		}

		return nil, fmt.Errorf("failed to get template")
	}

	pidsCount := 0
	if item.R.TemplateNamePidConfigs != nil {
		pidsCount = len(item.R.GetTemplateNamePidConfigs())
	}

	hasDbc := "false"
	if item.R.TemplateNameDBCFile != nil {
		hasDbc = "true" // ! todo: this should not be a string
	}

	result := &grpc.GetTemplateByNameResponse{
		Template: &grpc.Template{
			Name:       item.TemplateName,
			Version:    item.Version,
			Protocol:   item.Protocol,
			Powertrain: item.Powertrain,
			HasDbc:     hasDbc,
			PidsCount:  int32(pidsCount),
			CreatedAt:  timestamppb.New(item.CreatedAt),
			UpdatedAt:  timestamppb.New(item.UpdatedAt),
		},
	}

	if item.ParentTemplateName.Valid {
		result.Template.ParentTemplateName = item.ParentTemplateName.String
	}

	return result, nil
}

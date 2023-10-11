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

	item, err := models.Templates(models.TemplateWhere.TemplateName.EQ(query.Name)).One(ctx, h.DBS().Reader)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("template not found name: %s", query.Name)
		}

		return nil, fmt.Errorf("failed to get template")
	}

	result := &grpc.GetTemplateByNameResponse{
		Template: &grpc.Template{
			Name:               item.TemplateName,
			ParentTemplateName: "",
			Version:            item.Version,
			Protocol:           item.Protocol,
			Powertrain:         item.Powertrain,
			HasDbc:             item.R.GetTemplateNameDBCFile().DBCFile,
			PidsCount:          int32(len(item.R.GetTemplateNamePidConfigs())),
			Pids:               nil,
			CreatedAt:          timestamppb.New(item.CreatedAt),
			UpdatedAt:          timestamppb.New(item.UpdatedAt),
		},
	}
	if item.ParentTemplateName.Valid {
		result.Template.ParentTemplateName = item.ParentTemplateName.String
	}

	return result, nil
}

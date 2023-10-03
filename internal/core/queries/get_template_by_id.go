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

type GetTemplateByIDQueryHandler struct {
	DBS    func() *db.ReaderWriter
	logger *zerolog.Logger
}

func NewGetTemplateByIDQueryHandler(dbs func() *db.ReaderWriter, logger *zerolog.Logger) GetTemplateByIDQueryHandler {
	return GetTemplateByIDQueryHandler{
		DBS:    dbs,
		logger: logger,
	}
}

type GetTemplateByIDQueryRequest struct {
	ID string
}

func (h GetTemplateByIDQueryHandler) Handle(ctx context.Context, query *GetTemplateByIDQueryRequest) (*grpc.GetTemplateByIDResponse, error) {

	item, err := models.Templates(models.TemplateWhere.TemplateName.EQ(query.ID)).One(ctx, h.DBS().Reader)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("template not found id: %s", query.ID) // You may replace it with your custom NotFoundError
		}

		return nil, fmt.Errorf("failed to get template")
	}

	result := &grpc.GetTemplateByIDResponse{
		Template: &grpc.Template{
			Name:               item.TemplateName,
			ParentTemplateName: "",
			Version:            item.Version,
			Protocol:           item.Protocol,
			Powertrain:         item.Powertrain,
			//Need help from here
			HasDbc:    "", //need help
			PidsCount: 1,  //need help
			// Pids:
			Dbc: "",
			//Template Vehicles:
			// to here
			CreatedAt: timestamppb.New(item.CreatedAt),
			UpdatedAt: timestamppb.New(item.UpdatedAt),
		},
	}
	if item.ParentTemplateName.Valid {
		result.Template.ParentTemplateName = item.ParentTemplateName.String
	}

	return result, nil
}

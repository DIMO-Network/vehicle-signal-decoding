package commands

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UpdateDbcCommandHandler struct {
	DBS func() *db.ReaderWriter
}

func NewUpdateDbcCommandHandler(dbs func() *db.ReaderWriter) UpdateDbcCommandHandler {
	return UpdateDbcCommandHandler{DBS: dbs}
}

type UpdateDbcCommandRequest struct {
	TemplateName string
	DbcFile      string
}

type UpdateDbcCommandResponse struct {
	TemplateName string
}

func (h UpdateDbcCommandHandler) Execute(ctx context.Context, req *UpdateDbcCommandRequest) (*UpdateDbcCommandResponse, error) {

	dbc, err := models.DBCFiles(models.DBCFileWhere.TemplateName.EQ(req.TemplateName)).One(ctx, h.DBS().Reader)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &exceptions.NotFoundError{
				Err: fmt.Errorf("dbc file not found with template name: %s", req.TemplateName),
			}
		}
		return nil, &exceptions.InternalError{
			Err: err,
		}
	}

	dbc.TemplateName = req.TemplateName

	if _, err := dbc.Update(ctx, h.DBS().Writer.DB, boil.Infer()); err != nil {
		return nil, &exceptions.InternalError{
			Err: err,
		}
	}

	return &UpdateDbcCommandResponse{TemplateName: dbc.TemplateName}, nil
}

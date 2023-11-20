package commands

import (
	"context"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/DIMO-Network/shared/db"
)

type CreateDbcCommandHandler struct {
	DBS func() *db.ReaderWriter
}

func NewCreateDbcCommandHandler(dbs func() *db.ReaderWriter) CreateDbcCommandHandler {
	return CreateDbcCommandHandler{DBS: dbs}
}

type CreateDbcCommandRequest struct {
	TemplateName string
	DbcFile      string
}

type CreateDbcCommandResponse struct {
	Name string
}

func (h CreateDbcCommandHandler) Execute(ctx context.Context, req *CreateDbcCommandRequest) (*CreateDbcCommandResponse, error) {

	exists, err := models.DBCFiles(models.DBCFileWhere.TemplateName.EQ(req.TemplateName)).Exists(ctx, h.DBS().Reader)
	if err != nil {
		return nil, &exceptions.InternalError{
			Err: errors.Wrapf(err, "error checking if DbcConfig exists: %s", req.TemplateName),
		}
	}
	if exists {
		return nil, &exceptions.ConflictError{
			Err: errors.Errorf("DbcConfig already exists: %s", req.TemplateName),
		}
	}

	dbcConfig := &models.DBCFile{
		TemplateName: req.TemplateName,
		DBCFile:      req.DbcFile,
	}

	err = dbcConfig.Insert(ctx, h.DBS().Writer, boil.Infer())
	if err != nil {
		return nil, &exceptions.InternalError{
			Err: errors.Wrapf(err, "error inserting DbcConfig with template name: %s", req.TemplateName),
		}
	}

	return &CreateDbcCommandResponse{Name: dbcConfig.TemplateName}, nil
}

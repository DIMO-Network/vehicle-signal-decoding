package commands

import (
	"context"

	"github.com/DIMO-Network/shared/pkg/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UpdateDbcCommandHandler struct {
	DBS func() *db.ReaderWriter
}

func NewUpsertDbcCommandHandler(dbs func() *db.ReaderWriter) UpdateDbcCommandHandler {
	return UpdateDbcCommandHandler{DBS: dbs}
}

type UpsertDbcCommandRequest struct {
	TemplateName string
	DbcFile      string
}

func (h UpdateDbcCommandHandler) Execute(ctx context.Context, req *UpsertDbcCommandRequest) (*emptypb.Empty, error) {

	dbcFile := models.DBCFile{
		TemplateName: req.TemplateName,
		DBCFile:      req.DbcFile,
	}

	err := dbcFile.Upsert(ctx, h.DBS().Writer, true, []string{"template_name"}, boil.Whitelist("dbc_file"), boil.Infer())

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

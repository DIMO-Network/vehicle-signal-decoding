package api

import (
	"context"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/commands"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/queries"

	"github.com/DIMO-Network/shared/pkg/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DbcConfigService struct {
	grpc.DbcConfigServiceServer
	logger *zerolog.Logger
	dbs    func() *db.ReaderWriter
}

func NewDbcConfigService(logger *zerolog.Logger, dbs func() *db.ReaderWriter) grpc.DbcConfigServiceServer {
	return &DbcConfigService{logger: logger, dbs: dbs}
}

func (s *DbcConfigService) UpsertDbc(ctx context.Context, in *grpc.UpdateDbcRequest) (*emptypb.Empty, error) {
	service := commands.NewUpsertDbcCommandHandler(s.dbs)
	_, err := service.Execute(ctx, &commands.UpsertDbcCommandRequest{
		TemplateName: in.Dbc.TemplateName,
		DbcFile:      in.Dbc.DbcFile,
	})

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DbcConfigService) GetDbcList(ctx context.Context, in *grpc.GetDbcListRequest) (*grpc.GetDbcListResponse, error) {
	service := queries.NewGetDbcAllQueryHandler(s.dbs, s.logger)
	response, err := service.Handle(ctx, &queries.GetDbcAllQueryRequest{
		TemplateName: *in.TemplateName,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *DbcConfigService) GetDbcByTemplateName(ctx context.Context, in *grpc.GetDbcByTemplateNameRequest) (*grpc.GetDbcByTemplateNameResponse, error) {
	service := queries.NewGetDbcByTemplateNameQueryHandler(s.dbs, s.logger)
	response, err := service.Handle(ctx, &queries.GetDbcByTemplateNameQueryRequest{
		TemplateName: in.TemplateName,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

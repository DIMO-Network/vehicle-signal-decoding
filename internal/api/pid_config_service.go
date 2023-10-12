package api

import (
	"context"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/commands"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/queries"

	db "github.com/DIMO-Network/shared/db"
	grpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PidConfigService struct {
	grpc.PidConfigServiceServer
	logger *zerolog.Logger
	dbs    func() *db.ReaderWriter
}

func NewPidConfigService(logger *zerolog.Logger, dbs func() *db.ReaderWriter) grpc.PidConfigServiceServer {
	return &PidConfigService{logger: logger, dbs: dbs}
}

func (s *PidConfigService) CreatePid(ctx context.Context, in *grpc.UpdatePidRequest) (*emptypb.Empty, error) {
	service := commands.NewCreatePidCommandHandler(s.dbs)
	_, err := service.Execute(ctx, &commands.CreatePidCommandRequest{
		ID:              in.Pid.Id,
		TemplateName:    in.Pid.TemplateName,
		Header:          in.Pid.Header,
		Mode:            in.Pid.Mode,
		Pid:             in.Pid.Pid,
		Formula:         in.Pid.Formula,
		IntervalSeconds: in.Pid.IntervalSeconds,
		Protocol:        in.Pid.Protocol,
		SignalName:      in.Pid.SignalName,
	})

	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *PidConfigService) UpdatePid(ctx context.Context, in *grpc.UpdatePidRequest) (*emptypb.Empty, error) {
	service := commands.NewUpdatePidCommandHandler(s.dbs)
	_, err := service.Execute(ctx, &commands.UpdatePidCommandRequest{
		ID:              in.Pid.Id,
		TemplateName:    in.Pid.TemplateName,
		Header:          in.Pid.Header,
		Mode:            in.Pid.Mode,
		Pid:             in.Pid.Pid,
		Formula:         in.Pid.Formula,
		IntervalSeconds: in.Pid.IntervalSeconds,
		Protocol:        in.Pid.Protocol,
		SignalName:      in.Pid.SignalName,
	})

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *PidConfigService) GetPidList(ctx context.Context, in *grpc.GetPidListRequest) (*grpc.GetPidListResponse, error) {
	service := queries.NewGetPidAllQueryHandler(s.dbs, s.logger)
	response, err := service.Handle(ctx, &queries.GetPidAllQueryRequest{
		ID: *in.Id,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *PidConfigService) GetPidByID(ctx context.Context, in *grpc.GetPidByIDRequest) (*grpc.GetPidByIDResponse, error) {
	service := queries.NewGetPidByIDQueryHandler(s.dbs, s.logger)
	response, err := service.Handle(ctx, &queries.GetPidByIDQueryRequest{
		ID: in.Id,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}
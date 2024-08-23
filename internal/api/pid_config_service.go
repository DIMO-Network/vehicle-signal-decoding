package api

import (
	"context"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/commands"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/queries"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
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
		ID:                   in.Pid.Id,
		TemplateName:         in.Pid.TemplateName,
		Header:               in.Pid.Header,
		Mode:                 in.Pid.Mode,
		Pid:                  in.Pid.Pid,
		Formula:              in.Pid.Formula,
		IntervalSeconds:      in.Pid.IntervalSeconds,
		Protocol:             in.Pid.Protocol,
		SignalName:           in.Pid.SignalName,
		CanFlowControlClear:  in.Pid.CanFlowControlClear,
		CanFlowControlIDPair: in.Pid.CanFlowControlIdPair,
		VSSCovesaSignalName:  in.Pid.VssCovesaName,
		Unit:                 in.Pid.Unit,
	})

	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *PidConfigService) UpdatePid(ctx context.Context, in *grpc.UpdatePidRequest) (*emptypb.Empty, error) {
	service := commands.NewUpdatePidCommandHandler(s.dbs)
	_, err := service.Execute(ctx, &commands.UpdatePidCommandRequest{
		ID:                   in.Pid.Id,
		TemplateName:         in.Pid.TemplateName,
		Header:               in.Pid.Header,
		Mode:                 in.Pid.Mode,
		Pid:                  in.Pid.Pid,
		Formula:              in.Pid.Formula,
		IntervalSeconds:      in.Pid.IntervalSeconds,
		Protocol:             in.Pid.Protocol,
		SignalName:           in.Pid.SignalName,
		CanFlowControlClear:  in.Pid.CanFlowControlClear,
		CanFlowControlIDPair: in.Pid.CanFlowControlIdPair,
		Enabled:              in.Pid.Enabled,
		VSSCovesaSignalName:  in.Pid.VssCovesaName,
		Unit:                 in.Pid.Unit,
	})

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *PidConfigService) GetPidList(ctx context.Context, in *grpc.GetPidListRequest) (*grpc.GetPidListResponse, error) {
	service := queries.NewGetPidAllQueryHandler(s.dbs, s.logger)
	response, err := service.Handle(ctx, &queries.GetPidAllQueryRequest{
		TemplateName: in.TemplateName,
	})
	if err != nil {
		return nil, err
	}

	pidSummaries := make([]*grpc.PidSummary, 0)

	for _, item := range response {
		pidSummaries = append(pidSummaries, &grpc.PidSummary{
			Id:                   item.ID,
			TemplateName:         item.TemplateName,
			Header:               item.Header,
			Mode:                 item.Mode,
			Pid:                  item.Pid,
			Formula:              item.Formula,
			IntervalSeconds:      int32(item.IntervalSeconds),
			Protocol:             item.Protocol.String,
			SignalName:           item.SignalName,
			CanFlowControlClear:  item.CanFlowControlClear.Bool,
			CanFlowControlIdPair: item.CanFlowControlIDPair.String,
			Enabled:              item.Enabled,
			VssCovesaName:        item.VSSCovesaName.String,
			Unit:                 item.Unit.String,
			IsParentPid:          item.TemplateName != in.TemplateName,
		})
	}

	result := &grpc.GetPidListResponse{
		Pid: pidSummaries,
	}
	return result, nil
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

func (s *PidConfigService) DeletePid(ctx context.Context, in *grpc.DeletePidRequest) (*emptypb.Empty, error) {
	service := commands.NewDeletePidCommandHandler(s.dbs)
	err := service.Execute(ctx, &commands.DeletePidCommandRequest{
		ID:           in.Id,
		TemplateName: in.TemplateName,
	})

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *PidConfigService) ChangePidEnableStatus(ctx context.Context, in *grpc.ChangePidEnableStatusRequest) (*emptypb.Empty, error) {
	service := commands.NewUpdateEnableStatusPidCommandHandler(s.dbs)
	err := service.Execute(ctx, &commands.UpdateEnableStatusPidCommandRequest{
		ID: in.Id,
	})

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

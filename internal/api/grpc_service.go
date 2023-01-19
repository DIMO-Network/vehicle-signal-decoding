package api

import (
	"context"
	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/commands"

	p_grpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
)

type GrpcService struct {
	p_grpc.VehicleSignalDecodingServiceServer
	logger *zerolog.Logger
	DBS    func() *db.ReaderWriter
}

func NewGrpcService(logger *zerolog.Logger) p_grpc.VehicleSignalDecodingServiceServer {
	return &GrpcService{logger: logger}
}

func (s *GrpcService) ToDo(ctx context.Context, in *p_grpc.BaseRequest) (*p_grpc.BaseResponse, error) {
	service := commands.NewBaseCommandHandler(s.DBS)
	response, err := service.Execute(ctx, &commands.BaseCommandRequest{Input: in.Input})

	if err != nil {
		return nil, err
	}

	return &p_grpc.BaseResponse{Result: response.Result}, nil
}

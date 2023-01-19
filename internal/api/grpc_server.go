package api

import (
	"net"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/api/common"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	pkggrpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

func StartGrpcServer(logger zerolog.Logger, s *config.Settings) {
	lis, err := net.Listen("tcp", ":"+s.GRPCPort)
	if err != nil {
		logger.Fatal().Msgf("Failed to listen on port %v: %v", s.GRPCPort, err)
	}

	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(common.GrpcConfig),
	}

	vehicleSignalDecodingService := NewGrpcService(&logger)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(opts...),
		)),
	)

	pkggrpc.RegisterVehicleSignalDecodingServiceServer(server, vehicleSignalDecodingService)

	logger.Info().Str("port", s.GRPCPort).Msgf("started grpc server on port: %v", s.GRPCPort)

	if err := server.Serve(lis); err != nil {
		logger.Fatal().Msgf("Failed to serve over port %v: %v", s.GRPCPort, err)
	}
}

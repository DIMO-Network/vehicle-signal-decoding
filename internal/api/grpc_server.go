package api

import (
	"fmt"
	"net"
	"runtime/debug"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/DIMO-Network/shared/db"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	pkggrpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

func StartGrpcServer(logger zerolog.Logger, dbs func() *db.ReaderWriter, s *config.Settings) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", s.GRPCPort))
	if err != nil {
		logger.Fatal().Msgf("Failed to listen on port %v: %v", s.GRPCPort, err)
	}

	vehicleSignalDecodingService := NewGrpcService(&logger, dbs)
	templateConfigService := NewTemplateConfigService(&logger, dbs)
	pidConfigService := NewPidConfigService(&logger, dbs)
	deviceSettingsService := NewDeviceSettingsConfigService(&logger, dbs)
	dbcConfigService := NewDbcConfigService(&logger, dbs)
	vehicleTemplateService := NewVehicleTemplateService(&logger, dbs)

	grpcRecovery := GRPCPanicker{Logger: &logger}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcRecovery.GRPCPanicRecoveryHandler)),
		)),
	)

	pkggrpc.RegisterVehicleSignalDecodingServiceServer(server, vehicleSignalDecodingService)
	pkggrpc.RegisterTemplateConfigServiceServer(server, templateConfigService)
	pkggrpc.RegisterPidConfigServiceServer(server, pidConfigService)
	pkggrpc.RegisterDeviceSettingsServiceServer(server, deviceSettingsService)
	pkggrpc.RegisterDbcConfigServiceServer(server, dbcConfigService)
	pkggrpc.RegisterVehicleTemplateServiceServer(server, vehicleTemplateService)

	logger.Info().Str("port", s.GRPCPort).Msgf("started grpc server on port: %v", s.GRPCPort)

	if err := server.Serve(lis); err != nil {
		logger.Fatal().Msgf("Failed to serve over port %v: %v", s.GRPCPort, err)
	}
}

type GRPCPanicker struct {
	Logger *zerolog.Logger
}

func (pr *GRPCPanicker) GRPCPanicRecoveryHandler(p any) (err error) {
	//appmetrics.GRPCPanicsCount.Inc()

	pr.Logger.Err(fmt.Errorf("%s", p)).Str("stack", string(debug.Stack())).Msg("grpc recovered from panic")
	return status.Errorf(codes.Internal, "%s", p)
}

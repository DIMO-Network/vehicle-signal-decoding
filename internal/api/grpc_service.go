package api

import (
	"context"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/commands"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/queries"

	p_grpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type GrpcService struct {
	p_grpc.VehicleSignalDecodingServiceServer
	logger *zerolog.Logger
	DBS    func() *db.ReaderWriter
}

func NewGrpcService(logger *zerolog.Logger, dbs func() *db.ReaderWriter) p_grpc.VehicleSignalDecodingServiceServer {
	return &GrpcService{logger: logger, DBS: dbs}
}

func (s *GrpcService) CreateDBCCode(ctx context.Context, in *p_grpc.CreateDBCCodeRequest) (*p_grpc.VehicleSignalBaseResponse, error) {
	service := commands.NewCreateDBCCodeCommandHandler(s.DBS)
	response, err := service.Execute(ctx, &commands.CreateDBCCodeCommandRequest{
		Name:             in.Name,
		Trigger:          in.Trigger,
		DBCContents:      in.DbcContents,
		MaxSampleSize:    in.MaxSampleSize,
		Header:           int(in.Header),
		RecordingEnabled: in.RecordingEnabled,
	})

	if err != nil {
		return nil, err
	}

	return &p_grpc.VehicleSignalBaseResponse{Id: response.ID}, nil
}

func (s *GrpcService) UpdateDBCCode(ctx context.Context, in *p_grpc.UpdateDBCCodeRequest) (*p_grpc.VehicleSignalBaseResponse, error) {
	service := commands.NewUpdateDBCCodeCommandHandler(s.DBS)
	response, err := service.Execute(ctx, &commands.UpdateDBCCodeCommandRequest{
		ID:               in.Id,
		Name:             in.Name,
		Trigger:          in.Trigger,
		DBCContents:      in.DbcContents,
		MaxSampleSize:    in.MaxSampleSize,
		Header:           int(in.Header),
		RecordingEnabled: in.RecordingEnabled,
	})

	if err != nil {
		return nil, err
	}

	return &p_grpc.VehicleSignalBaseResponse{Id: response.ID}, nil
}

func (s *GrpcService) GetDBCCodes(ctx context.Context, _ *emptypb.Empty) (*p_grpc.GetDBCCodeListResponse, error) {
	service := queries.NewGetDBCCodeAllQueryHandler(s.DBS, s.logger)
	response, err := service.Handle(ctx, &queries.GetDBCCodeAllQueryRequest{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *GrpcService) GetDBCCodesByID(ctx context.Context, in *p_grpc.GetByIdRequest) (*p_grpc.GetDBCCodeResponse, error) {
	service := queries.NewGetDBCCodeByIDQueryHandler(s.DBS, s.logger)
	response, err := service.Handle(ctx, &queries.GetDBCCodeByIDQueryRequest{
		ID: in.Id,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *GrpcService) CreateTestSignal(ctx context.Context, in *p_grpc.CreateTestSignalRequest) (*p_grpc.VehicleSignalBaseResponse, error) {
	service := commands.NewCreateTestSignalCommandHandler(s.DBS)
	response, err := service.Execute(ctx, &commands.CreateTestSignalCommandRequest{
		DBCCodesID:         in.DbcCodesId,
		DeviceDefinitionID: in.DeviceDefinitionId,
		UserDeviceID:       in.UserDeviceId,
		Value:              in.Value,
		Approved:           in.Approved,
		AutoPIUnitID:       in.AutopiUnitId,
	})

	if err != nil {
		return nil, err
	}

	return &p_grpc.VehicleSignalBaseResponse{Id: response.ID}, nil
}

func (s *GrpcService) UpdateTestSignal(ctx context.Context, in *p_grpc.UpdateTestSignalRequest) (*p_grpc.VehicleSignalBaseResponse, error) {
	service := commands.NewUpdateTestSignalCommandHandler(s.DBS)
	response, err := service.Execute(ctx, &commands.UpdateTestSignalCommandRequest{
		ID:                 in.Id,
		DBCCodesID:         in.DbcCodesId,
		DeviceDefinitionID: in.DeviceDefinitionId,
		UserDeviceID:       in.UserDeviceId,
		Value:              in.Value,
		Approved:           in.Approved,
		AutoPIUnitID:       in.AutopiUnitId,
	})

	if err != nil {
		return nil, err
	}

	return &p_grpc.VehicleSignalBaseResponse{Id: response.ID}, nil
}

func (s *GrpcService) GetTestSignals(ctx context.Context, _ *emptypb.Empty) (*p_grpc.GetTestSignalListResponse, error) {
	service := queries.NewGetTestSignalAllQueryHandler(s.DBS, s.logger)
	response, err := service.Handle(ctx, &queries.GetTestSignalAllQueryRequest{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *GrpcService) GetTestSignalByID(ctx context.Context, in *p_grpc.GetByIdRequest) (*p_grpc.GetTestSignalResponse, error) {
	service := queries.NewGetTestSignalByIDQueryHandler(s.DBS, s.logger)
	response, err := service.Handle(ctx, &queries.GetTestSignalByIDQueryRequest{
		ID: in.Id,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *GrpcService) GetTestSignalsByDeviceDefinitionID(ctx context.Context, in *p_grpc.GetByIdRequest) (*p_grpc.GetTestSignalListResponse, error) {
	service := queries.NewGetTestSignalFilterQueryHandler(s.DBS, s.logger)
	response, err := service.Handle(ctx, &queries.GetTestSignalFilterQueryRequest{
		DeviceDefinitionID: in.Id,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *GrpcService) GetTestSignalsByUserDeviceID(ctx context.Context, in *p_grpc.GetByIdRequest) (*p_grpc.GetTestSignalListResponse, error) {
	service := queries.NewGetTestSignalFilterQueryHandler(s.DBS, s.logger)
	response, err := service.Handle(ctx, &queries.GetTestSignalFilterQueryRequest{
		UserDeviceID: in.Id,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *GrpcService) GetTestSignalsByDBCCodeID(ctx context.Context, in *p_grpc.GetByIdRequest) (*p_grpc.GetTestSignalListResponse, error) {
	service := queries.NewGetTestSignalFilterQueryHandler(s.DBS, s.logger)
	response, err := service.Handle(ctx, &queries.GetTestSignalFilterQueryRequest{
		DBCCodeID: in.Id,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

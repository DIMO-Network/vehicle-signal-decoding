package api

import (
	"context"
	
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/commands"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/queries"

	p_grpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type GrpcService struct {
	p_grpc.VehicleSignalDecodingServiceServer
	logger   *zerolog.Logger
	DBS      func() *db.ReaderWriter
	s3Client *s3.Client
	settings *config.Settings
}

func NewGrpcService(logger *zerolog.Logger, dbs func() *db.ReaderWriter, s3Client *s3.Client, settings *config.Settings) p_grpc.VehicleSignalDecodingServiceServer {
	return &GrpcService{logger: logger, DBS: dbs, settings: settings, s3Client: s3Client}
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

func (s *GrpcService) GetCanBusDumpFiles(ctx context.Context, in *p_grpc.GetCanBusDumpFileRequest) (*p_grpc.GetCanBusDumpFileResponse, error) {
	service := queries.NewGetCanBusDumpFileByEthAddressQueryHandler(s.logger, s.s3Client, s.settings)
	response, err := service.Handle(ctx, &queries.GetCanBusDumpFileByEthAddressQueryRequest{
		EthAddress: in.EthAddr,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *GrpcService) DownloadCanBusDumpFile(ctx context.Context, in *p_grpc.DownloadCanBusDumpFileContentRequest) (*p_grpc.DownloadCanBusDumpFileContentResponse, error) {
	service := queries.NewDownloadCanBusDumpFileByFileNameQueryHandler(s.logger, s.s3Client, s.settings)
	response, err := service.Handle(ctx, &queries.DownloadCanBusDumpFileByFileNameQueryRequest{
		FileName: in.Id,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *GrpcService) GetJobsByEtherumAddress(ctx context.Context, in *p_grpc.GetJobsByEtherumAddressRequest) (*p_grpc.GetJobsByEtherumAddressResponse, error) {
	service := queries.NewGetJobByEthereumAddressQueryHandler(s.DBS, s.logger)
	response, err := service.Handle(ctx, &queries.GetJobByyEthereumAddressQueryRequest{
		EtherumAddress: in.EtherumAddress,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *GrpcService) CreateJobsByEtherumAddress(ctx context.Context, in *p_grpc.CreateJobByEtherumAddressRequest) (*p_grpc.GetJobsByEtherumAddressItemResponse, error) {
	service := commands.NewCreateJobByEtherumAddressCommandHandler(s.DBS)
	response, err := service.Execute(ctx, &commands.CreateJobCommandRequest{
		EtherumAddress: in.EtherumAddress,
		Command:        in.Command,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *GrpcService) DeleteJobsByEtherumAddress(ctx context.Context, in *p_grpc.DeleteJobByEtherumAddressRequest) (*emptypb.Empty, error) {
	service := commands.NewDeleteJobByEtherumAddressCommandHandler(s.DBS)
	_, err := service.Execute(ctx, &commands.DeleteJobCommandRequest{
		ID: in.Id,
	})

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

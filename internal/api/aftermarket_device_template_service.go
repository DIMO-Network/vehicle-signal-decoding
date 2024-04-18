package api

import (
	"context"
	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AftermarketDeviceTemplateService struct {
	grpc.AftermarketDeviceTemplateServiceServer
	logger *zerolog.Logger
	dbs    func() *db.ReaderWriter
}

func NewAftermarketDeviceTemplateService(logger *zerolog.Logger, dbs func() *db.ReaderWriter) grpc.AftermarketDeviceTemplateServiceServer {
	return &AftermarketDeviceTemplateService{
		logger: logger,
		dbs:    dbs,
	}
}

func (s *AftermarketDeviceTemplateService) CreateAftermarketDeviceTemplate(ctx context.Context, request *grpc.AftermarketDeviceTemplateRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *AftermarketDeviceTemplateService) GetAftermarketDeviceTemplates(ctx context.Context, _ *emptypb.Empty) (*grpc.AftermarketDeviceTemplates, error) {
	return nil, nil
}

func (s *AftermarketDeviceTemplateService) UpdateAftermarketDeviceTemplate(ctx context.Context, request *grpc.AftermarketDeviceTemplateRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *AftermarketDeviceTemplateService) DeleteAftermarketDeviceTemplate(ctx context.Context, request *grpc.AftermarketDeviceTemplateRequest) (*emptypb.Empty, error) {
	return nil, nil
}

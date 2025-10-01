package api

import (
	"context"

	common2 "github.com/ethereum/go-ethereum/common"

	"github.com/DIMO-Network/shared/pkg/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/commands"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/queries"
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
	service := commands.NewCreateAftermarketDeviceTemplateCommandHandler(s.dbs)

	err := service.Execute(ctx, commands.CreateAftermarketDeviceTemplateCommand{
		AftermarketDeviceEthereumAddress: request.EthereumAddress,
		TemplateName:                     request.TemplateName,
	})

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *AftermarketDeviceTemplateService) GetAftermarketDeviceTemplates(ctx context.Context, _ *emptypb.Empty) (*grpc.AftermarketDeviceTemplates, error) {
	service := queries.NewGetAftermarketDeviceTemplateAll(s.dbs)

	response, err := service.Handle(ctx)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *AftermarketDeviceTemplateService) GetAftermarketDeviceTemplate(ctx context.Context, request *grpc.AftermarketDeviceTemplateRequest) (*grpc.AftermarketDeviceTemplate, error) {
	service := queries.NewGetAftermarketDeviceTemplateByEthereumAddress(s.dbs)

	ethAddress := common2.BytesToAddress(request.EthereumAddress)

	response, err := service.Handle(ctx, ethAddress)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *AftermarketDeviceTemplateService) DeleteAftermarketDeviceTemplate(ctx context.Context, request *grpc.AftermarketDeviceTemplateRequest) (*emptypb.Empty, error) {
	service := commands.NewDeleteAftermarketDeviceTemplateCommandHandler(s.dbs)

	err := service.Execute(ctx, commands.DeleteAftermarketDeviceTemplateCommand{
		AftermarketDeviceEthereumAddress: request.EthereumAddress,
		TemplateName:                     request.TemplateName,
	})

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

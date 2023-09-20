package services

import (
	"context"

	p_grpc "github.com/DIMO-Network/device-definitions-api/pkg/grpc"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//go:generate mockgen -source device_definitions_service.go -destination mocks/device_definitions_service_mock.go
type DeviceDefinitionsService interface {
	GetDeviceDefinitionByID(ctx context.Context, id string) (*p_grpc.GetDeviceDefinitionResponse, error)
	// Add other methods as required.
}

type deviceDefinitionsService struct {
	definitionsGRPCAddr string
}

func NewDeviceDefinitionsService(settings *config.Settings) DeviceDefinitionsService {
	return &deviceDefinitionsService{
		definitionsGRPCAddr: settings.DefinitionsGRPCAddr,
	}
}

func (d *deviceDefinitionsService) GetDeviceDefinitionByID(ctx context.Context, id string) (*p_grpc.GetDeviceDefinitionResponse, error) {
	definitionsClient, conn, err := d.getDefinitionsGrpcClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	request := &p_grpc.GetDeviceDefinitionRequest{
		Ids: []string{id},
	}
	response, err := definitionsClient.GetDeviceDefinitionByID(ctx, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (d *deviceDefinitionsService) getDefinitionsGrpcClient() (p_grpc.DeviceDefinitionServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(d.definitionsGRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, conn, err
	}
	definitionsClient := p_grpc.NewDeviceDefinitionServiceClient(conn)
	return definitionsClient, conn, nil
}
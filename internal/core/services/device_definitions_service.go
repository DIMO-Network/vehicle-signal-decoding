package services

import (
	"context"
	"fmt"

	p_grpc "github.com/DIMO-Network/device-definitions-api/pkg/grpc"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//go:generate mockgen -source device_definitions_service.go -destination mocks/device_definitions_service_mock.go
type DeviceDefinitionsService interface {
	GetDeviceDefinitionByID(ctx context.Context, id string) (*p_grpc.GetDeviceDefinitionItemResponse, error)
	DecodeVIN(ctx context.Context, vin string) (*p_grpc.DecodeVinResponse, error)
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

func (d *deviceDefinitionsService) GetDeviceDefinitionByID(ctx context.Context, id string) (*p_grpc.GetDeviceDefinitionItemResponse, error) {
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
	if len(response.DeviceDefinitions) == 0 {
		return nil, &exceptions.NotFoundError{
			Err: fmt.Errorf("no definition found with id: %s", id)}
	}

	return response.DeviceDefinitions[0], nil
}

func (d *deviceDefinitionsService) DecodeVIN(ctx context.Context, vin string) (*p_grpc.DecodeVinResponse, error) {
	_, conn, err := d.getDefinitionsGrpcClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	request := &p_grpc.DecodeVinRequest{
		Vin: vin,
	}
	decoderClient := p_grpc.NewVinDecoderServiceClient(conn)
	response, err := decoderClient.DecodeVin(ctx, request)
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

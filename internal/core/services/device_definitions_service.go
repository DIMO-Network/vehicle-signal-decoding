package services

import (
	"context"
	"fmt"

	pgrpc "github.com/DIMO-Network/device-definitions-api/pkg/grpc"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
)

//go:generate mockgen -source device_definitions_service.go -destination mocks/device_definitions_service_mock.go
type DeviceDefinitionsService interface {
	GetDeviceDefinitionByID(ctx context.Context, id string) (*pgrpc.GetDeviceDefinitionItemResponse, error)
	DecodeVIN(ctx context.Context, vin string) (*pgrpc.DecodeVinResponse, error)
}

type deviceDefinitionsService struct {
	definitionsClient pgrpc.DeviceDefinitionServiceClient
	vinDecoderClient  pgrpc.VinDecoderServiceClient
}

func NewDeviceDefinitionsService(definitionsClient pgrpc.DeviceDefinitionServiceClient, vinDecoderClient pgrpc.VinDecoderServiceClient) DeviceDefinitionsService {
	return &deviceDefinitionsService{
		definitionsClient: definitionsClient,
		vinDecoderClient:  vinDecoderClient,
	}
}

func (d *deviceDefinitionsService) GetDeviceDefinitionByID(ctx context.Context, id string) (*pgrpc.GetDeviceDefinitionItemResponse, error) {
	response, err := d.definitionsClient.GetDeviceDefinitionByID(ctx, &pgrpc.GetDeviceDefinitionRequest{
		Ids: []string{id},
	})
	if err != nil {
		return nil, err
	}
	if len(response.DeviceDefinitions) == 0 {
		return nil, &exceptions.NotFoundError{
			Err: fmt.Errorf("no definition found with id: %s", id)}
	}

	return response.DeviceDefinitions[0], nil
}

func (d *deviceDefinitionsService) DecodeVIN(ctx context.Context, vin string) (*pgrpc.DecodeVinResponse, error) {
	response, err := d.vinDecoderClient.DecodeVin(ctx, &pgrpc.DecodeVinRequest{
		Vin: vin,
	})
	if err != nil {
		return nil, err
	}

	return response, nil
}

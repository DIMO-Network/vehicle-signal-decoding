package services

import (
	"context"

	grpc "google.golang.org/grpc"

	pgrpc "github.com/DIMO-Network/device-definitions-api/pkg/grpc"
)

//go:generate mockgen -source device_definitions_service.go -destination mocks/device_definitions_service_mock.go
type DeviceDefinitionsService interface {
	DecodeVIN(ctx context.Context, vin string) (*pgrpc.DecodeVinResponse, error)
}

type deviceDefinitionsService struct {
	vinDecoderClient pgrpc.VinDecoderServiceClient
}

func NewDeviceDefinitionsService(definitionsConn *grpc.ClientConn) DeviceDefinitionsService {
	return &deviceDefinitionsService{
		vinDecoderClient: pgrpc.NewVinDecoderServiceClient(definitionsConn),
	}
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

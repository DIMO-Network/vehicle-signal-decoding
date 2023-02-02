package services

import (
	"context"

	pb "github.com/DIMO-Network/shared/api/devices"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/services/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//go:generate mockgen -source user_device_service.go -destination mocks/user_device_service_mock.go
type UserDeviceService interface {
	GetUserDeviceServiceByAutoPIUnitID(ctx context.Context, id string) (*models.UserDeviceAutoPIUnit, error)
}

type userDeviceService struct {
	deviceGRPCAddr string
}

func NewUserDeviceService(settings *config.Settings) UserDeviceService {
	return &userDeviceService{
		deviceGRPCAddr: settings.DeviceGRPCAddr,
	}
}

func (a *userDeviceService) GetUserDeviceServiceByAutoPIUnitID(ctx context.Context, id string) (*models.UserDeviceAutoPIUnit, error) {

	deviceClient, conn, err := a.getDeviceGrpcClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	userDevice, err := deviceClient.GetUserDeviceByAutoPIUnitId(ctx, &pb.GetUserDeviceByAutoPIUnitIdRequest{Id: id})

	if err != nil {
		return nil, err
	}

	return &models.UserDeviceAutoPIUnit{
		UserDeviceID:       userDevice.UserDeviceId,
		DeviceDefinitionID: userDevice.DeviceDefinitionId,
		DeviceStyleID:      userDevice.DeviceStyleId,
	}, nil
}

func (a *userDeviceService) getDeviceGrpcClient() (pb.UserDeviceServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(a.deviceGRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, conn, err
	}
	definitionsClient := pb.NewUserDeviceServiceClient(conn)
	return definitionsClient, conn, nil
}

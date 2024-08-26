package services

import (
	"context"
	"fmt"

	grpc "google.golang.org/grpc"

	"github.com/pkg/errors"

	common2 "github.com/ethereum/go-ethereum/common"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/appmodels"

	pb "github.com/DIMO-Network/devices-api/pkg/grpc"
)

//go:generate mockgen -source user_device_service.go -destination mocks/user_device_service_mock.go
type UserDeviceService interface {
	GetUserDeviceServiceByAutoPIUnitID(ctx context.Context, id string) (*appmodels.UserDeviceAutoPIUnit, error)
	GetUserDeviceByVIN(ctx context.Context, vin string) (*pb.UserDevice, error)
	GetUserDeviceByEthAddr(ctx context.Context, address common2.Address) (*pb.UserDevice, error)
	GetUserDevice(ctx context.Context, userDeviceID string) (*pb.UserDevice, error)
}

type userDeviceService struct {
	devicesClient pb.UserDeviceServiceClient
}

func NewUserDeviceService(devicesGrpcConn *grpc.ClientConn) UserDeviceService {
	return &userDeviceService{
		devicesClient: pb.NewUserDeviceServiceClient(devicesGrpcConn),
	}
}

// GetUserDevice gets the userDevice from devices-api, helpful to get the eth addr of the owner
func (a *userDeviceService) GetUserDevice(ctx context.Context, userDeviceID string) (*pb.UserDevice, error) {
	if len(userDeviceID) == 0 {
		return nil, fmt.Errorf("user device id was empty - invalid")
	}

	ud, err := a.devicesClient.GetUserDevice(ctx, &pb.GetUserDeviceRequest{
		Id: userDeviceID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to GetUserDevice")
	}
	return ud, nil
}

func (a *userDeviceService) GetUserDeviceServiceByAutoPIUnitID(ctx context.Context, id string) (*appmodels.UserDeviceAutoPIUnit, error) {
	userDevice, err := a.devicesClient.GetUserDeviceByAutoPIUnitId(ctx, &pb.GetUserDeviceByAutoPIUnitIdRequest{Id: id})
	if err != nil {
		return nil, errors.Wrap(err, "failed to GetUserDeviceByAutoPIUnitId")
	}

	return &appmodels.UserDeviceAutoPIUnit{
		UserDeviceID:       userDevice.UserDeviceId,
		DeviceDefinitionID: userDevice.DeviceDefinitionId,
		DeviceStyleID:      userDevice.DeviceStyleId,
	}, nil
}

func (a *userDeviceService) GetUserDeviceByVIN(ctx context.Context, vin string) (*pb.UserDevice, error) {
	userDevice, err := a.devicesClient.GetUserDeviceByVIN(ctx, &pb.GetUserDeviceByVINRequest{Vin: vin})
	if err != nil {
		return nil, errors.Wrap(err, "failed to GetUserDeviceByVIN")
	}

	return userDevice, nil
}

func (a *userDeviceService) GetUserDeviceByEthAddr(ctx context.Context, address common2.Address) (*pb.UserDevice, error) {
	userDevice, err := a.devicesClient.GetUserDeviceByEthAddr(ctx, &pb.GetUserDeviceByEthAddrRequest{EthAddr: address.Bytes()})
	if err != nil {
		return nil, errors.Wrap(err, "failed to GetUserDeviceByEthAddr")
	}

	return userDevice, nil
}

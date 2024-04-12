package services

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	common2 "github.com/ethereum/go-ethereum/common"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/appmodels"

	pb "github.com/DIMO-Network/devices-api/pkg/grpc"

	gdata "github.com/DIMO-Network/device-data-api/pkg/grpc"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//go:generate mockgen -source user_device_service.go -destination mocks/user_device_service_mock.go
type UserDeviceService interface {
	GetUserDeviceServiceByAutoPIUnitID(ctx context.Context, id string) (*appmodels.UserDeviceAutoPIUnit, error)
	GetUserDeviceByVIN(ctx context.Context, vin string) (*pb.UserDevice, error)
	GetUserDeviceByEthAddr(ctx context.Context, address common2.Address) (*pb.UserDevice, error)
	GetRawDeviceData(ctx context.Context, userDeviceID string) (*gdata.RawDeviceDataResponse, error)
	GetUserDevice(ctx context.Context, userDeviceID string) (*pb.UserDevice, error)
}

type userDeviceService struct {
	deviceGRPCAddr     string
	deviceDataGRPCAddr string
}

func NewUserDeviceService(settings *config.Settings) UserDeviceService {
	return &userDeviceService{
		deviceGRPCAddr:     settings.DeviceGRPCAddr,
		deviceDataGRPCAddr: settings.DeviceDataGRPCAddr,
	}
}

// GetUserDevice gets the userDevice from devices-api, helpful to get the eth addr of the owner
func (a *userDeviceService) GetUserDevice(ctx context.Context, userDeviceID string) (*pb.UserDevice, error) {
	if len(userDeviceID) == 0 {
		return nil, fmt.Errorf("user device id was empty - invalid")
	}
	deviceClient, conn, err := a.getDeviceGrpcClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	ud, err := deviceClient.GetUserDevice(ctx, &pb.GetUserDeviceRequest{
		Id: userDeviceID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to GetUserDevice")
	}
	return ud, nil
}

func (a *userDeviceService) GetUserDeviceServiceByAutoPIUnitID(ctx context.Context, id string) (*appmodels.UserDeviceAutoPIUnit, error) {

	deviceClient, conn, err := a.getDeviceGrpcClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	userDevice, err := deviceClient.GetUserDeviceByAutoPIUnitId(ctx, &pb.GetUserDeviceByAutoPIUnitIdRequest{Id: id})

	if err != nil {
		return nil, err
	}

	return &appmodels.UserDeviceAutoPIUnit{
		UserDeviceID:       userDevice.UserDeviceId,
		DeviceDefinitionID: userDevice.DeviceDefinitionId,
		DeviceStyleID:      userDevice.DeviceStyleId,
	}, nil
}

func (a *userDeviceService) GetUserDeviceByVIN(ctx context.Context, vin string) (*pb.UserDevice, error) {

	deviceClient, conn, err := a.getDeviceGrpcClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	userDevice, err := deviceClient.GetUserDeviceByVIN(ctx, &pb.GetUserDeviceByVINRequest{Vin: vin})

	if err != nil {
		return nil, err
	}

	return userDevice, nil
}

func (a *userDeviceService) GetUserDeviceByEthAddr(ctx context.Context, address common2.Address) (*pb.UserDevice, error) {
	// todo: better connection handling as singleton, with check to refresh, on sighup call close...
	deviceClient, conn, err := a.getDeviceGrpcClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	userDevice, err := deviceClient.GetUserDeviceByEthAddr(ctx, &pb.GetUserDeviceByEthAddrRequest{EthAddr: address.Bytes()})
	if err != nil {
		return nil, errors.Wrap(err, "failed to GetUserDeviceByEthAddr")
	}

	return userDevice, nil
}

func (a *userDeviceService) GetRawDeviceData(ctx context.Context, userDeviceID string) (*gdata.RawDeviceDataResponse, error) {
	conn, err := grpc.Dial(a.deviceDataGRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := gdata.NewUserDeviceDataServiceClient(conn)

	data, err := client.GetRawDeviceData(ctx, &gdata.RawDeviceDataRequest{
		UserDeviceId:  userDeviceID,
		IntegrationId: nil, // not needed, this will return all
	})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *userDeviceService) getDeviceGrpcClient() (pb.UserDeviceServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(a.deviceGRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, conn, err
	}
	userDeviceClient := pb.NewUserDeviceServiceClient(conn)
	return userDeviceClient, conn, nil
}

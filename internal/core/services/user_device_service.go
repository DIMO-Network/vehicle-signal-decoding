package services

import (
	"context"
	"encoding/hex"
	"fmt"

	pb "github.com/DIMO-Network/devices-api/pkg/grpc"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//go:generate mockgen -source user_device_service.go -destination mocks/user_device_service_mock.go
type UserDeviceService interface {
	GetUserDeviceServiceByAutoPIUnitID(ctx context.Context, id string) (*UserDeviceAutoPIUnit, error)
	GetUserDeviceByVIN(ctx context.Context, vin string) (*pb.UserDevice, error)
	GetUserDeviceByEthAddr(ctx context.Context, ethAddr string) (*pb.UserDevice, error)
}

type userDeviceService struct {
	deviceGRPCAddr string
}

func NewUserDeviceService(settings *config.Settings) UserDeviceService {
	return &userDeviceService{
		deviceGRPCAddr: settings.DeviceGRPCAddr,
	}
}

func (a *userDeviceService) GetUserDeviceServiceByAutoPIUnitID(ctx context.Context, id string) (*UserDeviceAutoPIUnit, error) {

	deviceClient, conn, err := a.getDeviceGrpcClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	userDevice, err := deviceClient.GetUserDeviceByAutoPIUnitId(ctx, &pb.GetUserDeviceByAutoPIUnitIdRequest{Id: id})

	if err != nil {
		return nil, err
	}

	return &UserDeviceAutoPIUnit{
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
func (a *userDeviceService) GetUserDeviceByEthAddr(ctx context.Context, ethAddr string) (*pb.UserDevice, error) {
	deviceClient, conn, err := a.getDeviceGrpcClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if len(ethAddr) > 2 && ethAddr[:2] == "0x" {
		ethAddr = ethAddr[2:]
	}

	ethAddrBytes, err := hex.DecodeString(ethAddr)
	if err != nil {
		return nil, fmt.Errorf("invalid ethereum address: %w", err)
	}

	userDevice, err := deviceClient.GetUserDeviceByEthAddr(ctx, &pb.GetUserDeviceByEthAddrRequest{EthAddr: ethAddrBytes})
	if err != nil {
		return nil, err
	}

	return userDevice, nil
}

func (a *userDeviceService) getDeviceGrpcClient() (pb.UserDeviceServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(a.deviceGRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, conn, err
	}
	definitionsClient := pb.NewUserDeviceServiceClient(conn)
	return definitionsClient, conn, nil
}

type UserDeviceAutoPIUnit struct {
	UserDeviceID       string
	DeviceDefinitionID string
	DeviceStyleID      string
}

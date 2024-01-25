package services

import (
	"context"
	"encoding/hex"
	"fmt"

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
	GetUserDeviceByEthAddr(ctx context.Context, ethAddr string) (*pb.UserDevice, error)
	GetRawDeviceData(ctx context.Context, userDeviceID string) (*gdata.RawDeviceDataResponse, error)
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
	definitionsClient := pb.NewUserDeviceServiceClient(conn)
	return definitionsClient, conn, nil
}

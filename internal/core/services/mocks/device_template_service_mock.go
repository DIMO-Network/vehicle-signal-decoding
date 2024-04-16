// Code generated by MockGen. DO NOT EDIT.
// Source: device_template_service.go
//
// Generated by this command:
//
//	mockgen -source device_template_service.go -destination mocks/device_template_service_mock.go
//
// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	grpc "github.com/DIMO-Network/devices-api/pkg/grpc"
	appmodels "github.com/DIMO-Network/vehicle-signal-decoding/internal/core/appmodels"
	gateways "github.com/DIMO-Network/vehicle-signal-decoding/internal/gateways"
	models "github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	common "github.com/ethereum/go-ethereum/common"
	fiber "github.com/gofiber/fiber/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockDeviceTemplateService is a mock of DeviceTemplateService interface.
type MockDeviceTemplateService struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceTemplateServiceMockRecorder
}

// MockDeviceTemplateServiceMockRecorder is the mock recorder for MockDeviceTemplateService.
type MockDeviceTemplateServiceMockRecorder struct {
	mock *MockDeviceTemplateService
}

// NewMockDeviceTemplateService creates a new mock instance.
func NewMockDeviceTemplateService(ctrl *gomock.Controller) *MockDeviceTemplateService {
	mock := &MockDeviceTemplateService{ctrl: ctrl}
	mock.recorder = &MockDeviceTemplateServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeviceTemplateService) EXPECT() *MockDeviceTemplateServiceMockRecorder {
	return m.recorder
}

// ResolveDeviceConfiguration mocks base method.
func (m *MockDeviceTemplateService) ResolveDeviceConfiguration(c *fiber.Ctx, ud *grpc.UserDevice, vehicle *gateways.VehicleInfo) (*appmodels.DeviceConfigResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveDeviceConfiguration", c, ud, vehicle)
	ret0, _ := ret[0].(*appmodels.DeviceConfigResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveDeviceConfiguration indicates an expected call of ResolveDeviceConfiguration.
func (mr *MockDeviceTemplateServiceMockRecorder) ResolveDeviceConfiguration(c, ud, vehicle any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveDeviceConfiguration", reflect.TypeOf((*MockDeviceTemplateService)(nil).ResolveDeviceConfiguration), c, ud, vehicle)
}

// StoreDeviceConfigUsed mocks base method.
func (m *MockDeviceTemplateService) StoreDeviceConfigUsed(ctx context.Context, address common.Address, dbcURL, pidURL, settingURL, firmwareVersion *string) (*models.DeviceTemplateStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreDeviceConfigUsed", ctx, address, dbcURL, pidURL, settingURL, firmwareVersion)
	ret0, _ := ret[0].(*models.DeviceTemplateStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreDeviceConfigUsed indicates an expected call of StoreDeviceConfigUsed.
func (mr *MockDeviceTemplateServiceMockRecorder) StoreDeviceConfigUsed(ctx, address, dbcURL, pidURL, settingURL, firmwareVersion any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreDeviceConfigUsed", reflect.TypeOf((*MockDeviceTemplateService)(nil).StoreDeviceConfigUsed), ctx, address, dbcURL, pidURL, settingURL, firmwareVersion)
}

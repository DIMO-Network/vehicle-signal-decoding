// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/services/user_device_service.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	grpc "github.com/DIMO-Network/devices-api/pkg/grpc"
	services "github.com/DIMO-Network/vehicle-signal-decoding/internal/core/services"
	gomock "github.com/golang/mock/gomock"
)

// MockUserDeviceService is a mock of UserDeviceService interface.
type MockUserDeviceService struct {
	ctrl     *gomock.Controller
	recorder *MockUserDeviceServiceMockRecorder
}

// MockUserDeviceServiceMockRecorder is the mock recorder for MockUserDeviceService.
type MockUserDeviceServiceMockRecorder struct {
	mock *MockUserDeviceService
}

// NewMockUserDeviceService creates a new mock instance.
func NewMockUserDeviceService(ctrl *gomock.Controller) *MockUserDeviceService {
	mock := &MockUserDeviceService{ctrl: ctrl}
	mock.recorder = &MockUserDeviceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDeviceService) EXPECT() *MockUserDeviceServiceMockRecorder {
	return m.recorder
}

// GetUserDeviceByEthAddr mocks base method.
func (m *MockUserDeviceService) GetUserDeviceByEthAddr(ctx context.Context, ethAddr string) (*grpc.UserDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDeviceByEthAddr", ctx, ethAddr)
	ret0, _ := ret[0].(*grpc.UserDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserDeviceByEthAddr indicates an expected call of GetUserDeviceByEthAddr.
func (mr *MockUserDeviceServiceMockRecorder) GetUserDeviceByEthAddr(ctx, ethAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDeviceByEthAddr", reflect.TypeOf((*MockUserDeviceService)(nil).GetUserDeviceByEthAddr), ctx, ethAddr)
}

// GetUserDeviceByVIN mocks base method.
func (m *MockUserDeviceService) GetUserDeviceByVIN(ctx context.Context, vin string) (*grpc.UserDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDeviceByVIN", ctx, vin)
	ret0, _ := ret[0].(*grpc.UserDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserDeviceByVIN indicates an expected call of GetUserDeviceByVIN.
func (mr *MockUserDeviceServiceMockRecorder) GetUserDeviceByVIN(ctx, vin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDeviceByVIN", reflect.TypeOf((*MockUserDeviceService)(nil).GetUserDeviceByVIN), ctx, vin)
}

// GetUserDeviceServiceByAutoPIUnitID mocks base method.
func (m *MockUserDeviceService) GetUserDeviceServiceByAutoPIUnitID(ctx context.Context, id string) (*services.UserDeviceAutoPIUnit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDeviceServiceByAutoPIUnitID", ctx, id)
	ret0, _ := ret[0].(*services.UserDeviceAutoPIUnit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserDeviceServiceByAutoPIUnitID indicates an expected call of GetUserDeviceServiceByAutoPIUnitID.
func (mr *MockUserDeviceServiceMockRecorder) GetUserDeviceServiceByAutoPIUnitID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDeviceServiceByAutoPIUnitID", reflect.TypeOf((*MockUserDeviceService)(nil).GetUserDeviceServiceByAutoPIUnitID), ctx, id)
}

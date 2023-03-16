// Code generated by MockGen. DO NOT EDIT.
// Source: user_device_service.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

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
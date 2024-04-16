// Code generated by MockGen. DO NOT EDIT.
// Source: identity_api.go
//
// Generated by this command:
//
//	mockgen -source identity_api.go -destination mocks/identity_api_mock.go
//
// Package mock_gateways is a generated GoMock package.
package mock_gateways

import (
	reflect "reflect"

	gateways "github.com/DIMO-Network/vehicle-signal-decoding/internal/gateways"
	common "github.com/ethereum/go-ethereum/common"
	gomock "go.uber.org/mock/gomock"
)

// MockIdentityAPI is a mock of IdentityAPI interface.
type MockIdentityAPI struct {
	ctrl     *gomock.Controller
	recorder *MockIdentityAPIMockRecorder
}

// MockIdentityAPIMockRecorder is the mock recorder for MockIdentityAPI.
type MockIdentityAPIMockRecorder struct {
	mock *MockIdentityAPI
}

// NewMockIdentityAPI creates a new mock instance.
func NewMockIdentityAPI(ctrl *gomock.Controller) *MockIdentityAPI {
	mock := &MockIdentityAPI{ctrl: ctrl}
	mock.recorder = &MockIdentityAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIdentityAPI) EXPECT() *MockIdentityAPIMockRecorder {
	return m.recorder
}

// QueryIdentityAPIForVehicle mocks base method.
func (m *MockIdentityAPI) QueryIdentityAPIForVehicle(ethAddress common.Address) (*gateways.VehicleInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryIdentityAPIForVehicle", ethAddress)
	ret0, _ := ret[0].(*gateways.VehicleInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryIdentityAPIForVehicle indicates an expected call of QueryIdentityAPIForVehicle.
func (mr *MockIdentityAPIMockRecorder) QueryIdentityAPIForVehicle(ethAddress any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryIdentityAPIForVehicle", reflect.TypeOf((*MockIdentityAPI)(nil).QueryIdentityAPIForVehicle), ethAddress)
}

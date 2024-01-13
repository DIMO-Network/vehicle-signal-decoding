// Code generated by MockGen. DO NOT EDIT.
// Source: user_device_template_service.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserDeviceTemplateService is a mock of UserDeviceTemplateService interface.
type MockUserDeviceTemplateService struct {
	ctrl     *gomock.Controller
	recorder *MockUserDeviceTemplateServiceMockRecorder
}

// MockUserDeviceTemplateServiceMockRecorder is the mock recorder for MockUserDeviceTemplateService.
type MockUserDeviceTemplateServiceMockRecorder struct {
	mock *MockUserDeviceTemplateService
}

// NewMockUserDeviceTemplateService creates a new mock instance.
func NewMockUserDeviceTemplateService(ctrl *gomock.Controller) *MockUserDeviceTemplateService {
	mock := &MockUserDeviceTemplateService{ctrl: ctrl}
	mock.recorder = &MockUserDeviceTemplateServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDeviceTemplateService) EXPECT() *MockUserDeviceTemplateServiceMockRecorder {
	return m.recorder
}

// AssociateTemplate mocks base method.
func (m *MockUserDeviceTemplateService) AssociateTemplate(ctx context.Context, vin, templateDbcUrl, templatePidUrl, templateSettingUrl, version string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssociateTemplate", ctx, vin, templateDbcUrl, templatePidUrl, templateSettingUrl, version)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssociateTemplate indicates an expected call of AssociateTemplate.
func (mr *MockUserDeviceTemplateServiceMockRecorder) AssociateTemplate(ctx, vin, templateDbcUrl, templatePidUrl, templateSettingUrl, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssociateTemplate", reflect.TypeOf((*MockUserDeviceTemplateService)(nil).AssociateTemplate), ctx, vin, templateDbcUrl, templatePidUrl, templateSettingUrl, version)
}
// Code generated by MockGen. DO NOT EDIT.
// Source: rabbit_instances_service.go

// Package mock_instance is a generated GoMock package.
package mock_instance

import (
	context "context"
	model "maas/maas-service/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRabbitInstanceService is a mock of RabbitInstanceService interface.
type MockRabbitInstanceService struct {
	ctrl     *gomock.Controller
	recorder *MockRabbitInstanceServiceMockRecorder
}

// MockRabbitInstanceServiceMockRecorder is the mock recorder for MockRabbitInstanceService.
type MockRabbitInstanceServiceMockRecorder struct {
	mock *MockRabbitInstanceService
}

// NewMockRabbitInstanceService creates a new mock instance.
func NewMockRabbitInstanceService(ctrl *gomock.Controller) *MockRabbitInstanceService {
	mock := &MockRabbitInstanceService{ctrl: ctrl}
	mock.recorder = &MockRabbitInstanceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRabbitInstanceService) EXPECT() *MockRabbitInstanceServiceMockRecorder {
	return m.recorder
}

// AddInstanceUpdateCallback mocks base method.
func (m *MockRabbitInstanceService) AddInstanceUpdateCallback(arg0 func(context.Context, *model.RabbitInstance)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddInstanceUpdateCallback", arg0)
}

// AddInstanceUpdateCallback indicates an expected call of AddInstanceUpdateCallback.
func (mr *MockRabbitInstanceServiceMockRecorder) AddInstanceUpdateCallback(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddInstanceUpdateCallback", reflect.TypeOf((*MockRabbitInstanceService)(nil).AddInstanceUpdateCallback), arg0)
}

// CheckHealth mocks base method.
func (m *MockRabbitInstanceService) CheckHealth(arg0 context.Context, arg1 *model.RabbitInstance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckHealth", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckHealth indicates an expected call of CheckHealth.
func (mr *MockRabbitInstanceServiceMockRecorder) CheckHealth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckHealth", reflect.TypeOf((*MockRabbitInstanceService)(nil).CheckHealth), arg0, arg1)
}

// DeleteRabbitInstanceDesignatorByNamespace mocks base method.
func (m *MockRabbitInstanceService) DeleteRabbitInstanceDesignatorByNamespace(ctx context.Context, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRabbitInstanceDesignatorByNamespace", ctx, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRabbitInstanceDesignatorByNamespace indicates an expected call of DeleteRabbitInstanceDesignatorByNamespace.
func (mr *MockRabbitInstanceServiceMockRecorder) DeleteRabbitInstanceDesignatorByNamespace(ctx, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRabbitInstanceDesignatorByNamespace", reflect.TypeOf((*MockRabbitInstanceService)(nil).DeleteRabbitInstanceDesignatorByNamespace), ctx, namespace)
}

// GetById mocks base method.
func (m *MockRabbitInstanceService) GetById(ctx context.Context, id string) (*model.RabbitInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(*model.RabbitInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockRabbitInstanceServiceMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockRabbitInstanceService)(nil).GetById), ctx, id)
}

// GetDefault mocks base method.
func (m *MockRabbitInstanceService) GetDefault(ctx context.Context) (*model.RabbitInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefault", ctx)
	ret0, _ := ret[0].(*model.RabbitInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDefault indicates an expected call of GetDefault.
func (mr *MockRabbitInstanceServiceMockRecorder) GetDefault(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefault", reflect.TypeOf((*MockRabbitInstanceService)(nil).GetDefault), ctx)
}

// GetRabbitInstanceDesignatorByNamespace mocks base method.
func (m *MockRabbitInstanceService) GetRabbitInstanceDesignatorByNamespace(ctx context.Context, namespace string) (*model.InstanceDesignatorRabbit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRabbitInstanceDesignatorByNamespace", ctx, namespace)
	ret0, _ := ret[0].(*model.InstanceDesignatorRabbit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRabbitInstanceDesignatorByNamespace indicates an expected call of GetRabbitInstanceDesignatorByNamespace.
func (mr *MockRabbitInstanceServiceMockRecorder) GetRabbitInstanceDesignatorByNamespace(ctx, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRabbitInstanceDesignatorByNamespace", reflect.TypeOf((*MockRabbitInstanceService)(nil).GetRabbitInstanceDesignatorByNamespace), ctx, namespace)
}

// GetRabbitInstances mocks base method.
func (m *MockRabbitInstanceService) GetRabbitInstances(ctx context.Context) (*[]model.RabbitInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRabbitInstances", ctx)
	ret0, _ := ret[0].(*[]model.RabbitInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRabbitInstances indicates an expected call of GetRabbitInstances.
func (mr *MockRabbitInstanceServiceMockRecorder) GetRabbitInstances(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRabbitInstances", reflect.TypeOf((*MockRabbitInstanceService)(nil).GetRabbitInstances), ctx)
}

// Register mocks base method.
func (m *MockRabbitInstanceService) Register(ctx context.Context, instance *model.RabbitInstance) (*model.RabbitInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, instance)
	ret0, _ := ret[0].(*model.RabbitInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockRabbitInstanceServiceMockRecorder) Register(ctx, instance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRabbitInstanceService)(nil).Register), ctx, instance)
}

// SetDefault mocks base method.
func (m *MockRabbitInstanceService) SetDefault(ctx context.Context, instance *model.RabbitInstance) (*model.RabbitInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDefault", ctx, instance)
	ret0, _ := ret[0].(*model.RabbitInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetDefault indicates an expected call of SetDefault.
func (mr *MockRabbitInstanceServiceMockRecorder) SetDefault(ctx, instance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDefault", reflect.TypeOf((*MockRabbitInstanceService)(nil).SetDefault), ctx, instance)
}

// Unregister mocks base method.
func (m *MockRabbitInstanceService) Unregister(ctx context.Context, instanceId string) (*model.RabbitInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unregister", ctx, instanceId)
	ret0, _ := ret[0].(*model.RabbitInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unregister indicates an expected call of Unregister.
func (mr *MockRabbitInstanceServiceMockRecorder) Unregister(ctx, instanceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unregister", reflect.TypeOf((*MockRabbitInstanceService)(nil).Unregister), ctx, instanceId)
}

// Update mocks base method.
func (m *MockRabbitInstanceService) Update(ctx context.Context, instance *model.RabbitInstance) (*model.RabbitInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, instance)
	ret0, _ := ret[0].(*model.RabbitInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRabbitInstanceServiceMockRecorder) Update(ctx, instance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRabbitInstanceService)(nil).Update), ctx, instance)
}

// UpsertRabbitInstanceDesignator mocks base method.
func (m *MockRabbitInstanceService) UpsertRabbitInstanceDesignator(ctx context.Context, instanceDesignator model.InstanceDesignatorRabbit, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertRabbitInstanceDesignator", ctx, instanceDesignator, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertRabbitInstanceDesignator indicates an expected call of UpsertRabbitInstanceDesignator.
func (mr *MockRabbitInstanceServiceMockRecorder) UpsertRabbitInstanceDesignator(ctx, instanceDesignator, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertRabbitInstanceDesignator", reflect.TypeOf((*MockRabbitInstanceService)(nil).UpsertRabbitInstanceDesignator), ctx, instanceDesignator, namespace)
}

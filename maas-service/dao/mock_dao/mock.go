// Code generated by MockGen. DO NOT EDIT.
// Source: dao.go

// Package mock_dao is a generated GoMock package.
package mock_dao

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockBaseDao is a mock of BaseDao interface.
type MockBaseDao struct {
	ctrl     *gomock.Controller
	recorder *MockBaseDaoMockRecorder
}

// MockBaseDaoMockRecorder is the mock recorder for MockBaseDao.
type MockBaseDaoMockRecorder struct {
	mock *MockBaseDao
}

// NewMockBaseDao creates a new mock instance.
func NewMockBaseDao(ctrl *gomock.Controller) *MockBaseDao {
	mock := &MockBaseDao{ctrl: ctrl}
	mock.recorder = &MockBaseDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBaseDao) EXPECT() *MockBaseDaoMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockBaseDao) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockBaseDaoMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockBaseDao)(nil).Close))
}

// DSN mocks base method.
func (m *MockBaseDao) DSN() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DSN")
	ret0, _ := ret[0].(string)
	return ret0
}

// DSN indicates an expected call of DSN.
func (mr *MockBaseDaoMockRecorder) DSN() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DSN", reflect.TypeOf((*MockBaseDao)(nil).DSN))
}

// IsAvailable mocks base method.
func (m *MockBaseDao) IsAvailable() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAvailable")
	ret0, _ := ret[0].(error)
	return ret0
}

// IsAvailable indicates an expected call of IsAvailable.
func (mr *MockBaseDaoMockRecorder) IsAvailable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAvailable", reflect.TypeOf((*MockBaseDao)(nil).IsAvailable))
}

// IsForeignKeyIntegrityViolation mocks base method.
func (m *MockBaseDao) IsForeignKeyIntegrityViolation(err error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsForeignKeyIntegrityViolation", err)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsForeignKeyIntegrityViolation indicates an expected call of IsForeignKeyIntegrityViolation.
func (mr *MockBaseDaoMockRecorder) IsForeignKeyIntegrityViolation(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsForeignKeyIntegrityViolation", reflect.TypeOf((*MockBaseDao)(nil).IsForeignKeyIntegrityViolation), err)
}

// IsUniqIntegrityViolation mocks base method.
func (m *MockBaseDao) IsUniqIntegrityViolation(err error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUniqIntegrityViolation", err)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUniqIntegrityViolation indicates an expected call of IsUniqIntegrityViolation.
func (mr *MockBaseDaoMockRecorder) IsUniqIntegrityViolation(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUniqIntegrityViolation", reflect.TypeOf((*MockBaseDao)(nil).IsUniqIntegrityViolation), err)
}

// StartCache mocks base method.
func (m *MockBaseDao) StartCache(ctx context.Context, fullCacheResyncInterval time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartCache", ctx, fullCacheResyncInterval)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartCache indicates an expected call of StartCache.
func (mr *MockBaseDaoMockRecorder) StartCache(ctx, fullCacheResyncInterval interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartCache", reflect.TypeOf((*MockBaseDao)(nil).StartCache), ctx, fullCacheResyncInterval)
}

// StartMonitor mocks base method.
func (m *MockBaseDao) StartMonitor(ctx context.Context, masterMonitorCheckInterval time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartMonitor", ctx, masterMonitorCheckInterval)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartMonitor indicates an expected call of StartMonitor.
func (mr *MockBaseDaoMockRecorder) StartMonitor(ctx, masterMonitorCheckInterval interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartMonitor", reflect.TypeOf((*MockBaseDao)(nil).StartMonitor), ctx, masterMonitorCheckInterval)
}

// UsingDb mocks base method.
func (m *MockBaseDao) UsingDb(ctx context.Context, f func(*gorm.DB) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UsingDb", ctx, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// UsingDb indicates an expected call of UsingDb.
func (mr *MockBaseDaoMockRecorder) UsingDb(ctx, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsingDb", reflect.TypeOf((*MockBaseDao)(nil).UsingDb), ctx, f)
}

// WithLock mocks base method.
func (m *MockBaseDao) WithLock(ctx context.Context, lockId string, f func(context.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithLock", ctx, lockId, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithLock indicates an expected call of WithLock.
func (mr *MockBaseDaoMockRecorder) WithLock(ctx, lockId, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithLock", reflect.TypeOf((*MockBaseDao)(nil).WithLock), ctx, lockId, f)
}

// WithTx mocks base method.
func (m *MockBaseDao) WithTx(ctx context.Context, f func(context.Context, *gorm.DB) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTx", ctx, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithTx indicates an expected call of WithTx.
func (mr *MockBaseDaoMockRecorder) WithTx(ctx, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTx", reflect.TypeOf((*MockBaseDao)(nil).WithTx), ctx, f)
}

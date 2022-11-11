// Code generated by MockGen. DO NOT EDIT.
// Source: personia/domain/hr (interfaces: HierarchyRepo)

// Package hrmock is a generated GoMock package.
package hrmock

import (
	context "context"
	hr "personia/domain/hr"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHierarchyRepo is a mock of HierarchyRepo interface.
type MockHierarchyRepo struct {
	ctrl     *gomock.Controller
	recorder *MockHierarchyRepoMockRecorder
}

// MockHierarchyRepoMockRecorder is the mock recorder for MockHierarchyRepo.
type MockHierarchyRepoMockRecorder struct {
	mock *MockHierarchyRepo
}

// NewMockHierarchyRepo creates a new mock instance.
func NewMockHierarchyRepo(ctrl *gomock.Controller) *MockHierarchyRepo {
	mock := &MockHierarchyRepo{ctrl: ctrl}
	mock.recorder = &MockHierarchyRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHierarchyRepo) EXPECT() *MockHierarchyRepoMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockHierarchyRepo) Get(arg0 context.Context) (hr.Hierarchy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(hr.Hierarchy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockHierarchyRepoMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHierarchyRepo)(nil).Get), arg0)
}

// Update mocks base method.
func (m *MockHierarchyRepo) Update(arg0 context.Context, arg1 hr.Hierarchy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockHierarchyRepoMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockHierarchyRepo)(nil).Update), arg0, arg1)
}
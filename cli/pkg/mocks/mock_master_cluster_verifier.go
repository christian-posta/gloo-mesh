// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/mesh-projects/cli/pkg/common/config (interfaces: MasterKubeConfigVerifier)

// Package cli_mocks is a generated GoMock package.
package cli_mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMasterKubeConfigVerifier is a mock of MasterKubeConfigVerifier interface
type MockMasterKubeConfigVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockMasterKubeConfigVerifierMockRecorder
}

// MockMasterKubeConfigVerifierMockRecorder is the mock recorder for MockMasterKubeConfigVerifier
type MockMasterKubeConfigVerifierMockRecorder struct {
	mock *MockMasterKubeConfigVerifier
}

// NewMockMasterKubeConfigVerifier creates a new mock instance
func NewMockMasterKubeConfigVerifier(ctrl *gomock.Controller) *MockMasterKubeConfigVerifier {
	mock := &MockMasterKubeConfigVerifier{ctrl: ctrl}
	mock.recorder = &MockMasterKubeConfigVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMasterKubeConfigVerifier) EXPECT() *MockMasterKubeConfigVerifierMockRecorder {
	return m.recorder
}

// Verify mocks base method
func (m *MockMasterKubeConfigVerifier) Verify(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify
func (mr *MockMasterKubeConfigVerifierMockRecorder) Verify(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockMasterKubeConfigVerifier)(nil).Verify), arg0, arg1)
}

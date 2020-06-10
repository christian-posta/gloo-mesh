// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_cluster_registration is a generated GoMock package.
package mock_cluster_registration

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	cluster_registration "github.com/solo-io/service-mesh-hub/pkg/cluster-registration"
	clientcmd "k8s.io/client-go/tools/clientcmd"
)

// MockClusterRegistrationClient is a mock of ClusterRegistrationClient interface.
type MockClusterRegistrationClient struct {
	ctrl     *gomock.Controller
	recorder *MockClusterRegistrationClientMockRecorder
}

// MockClusterRegistrationClientMockRecorder is the mock recorder for MockClusterRegistrationClient.
type MockClusterRegistrationClientMockRecorder struct {
	mock *MockClusterRegistrationClient
}

// NewMockClusterRegistrationClient creates a new mock instance.
func NewMockClusterRegistrationClient(ctrl *gomock.Controller) *MockClusterRegistrationClient {
	mock := &MockClusterRegistrationClient{ctrl: ctrl}
	mock.recorder = &MockClusterRegistrationClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterRegistrationClient) EXPECT() *MockClusterRegistrationClientMockRecorder {
	return m.recorder
}

// Register mocks base method.
func (m *MockClusterRegistrationClient) Register(ctx context.Context, remoteConfig clientcmd.ClientConfig, remoteClusterName, remoteWriteNamespace, remoteContextName, discoverySource string, registerOpts cluster_registration.ClusterRegisterOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, remoteConfig, remoteClusterName, remoteWriteNamespace, remoteContextName, discoverySource, registerOpts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockClusterRegistrationClientMockRecorder) Register(ctx, remoteConfig, remoteClusterName, remoteWriteNamespace, remoteContextName, discoverySource, registerOpts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockClusterRegistrationClient)(nil).Register), ctx, remoteConfig, remoteClusterName, remoteWriteNamespace, remoteContextName, discoverySource, registerOpts)
}

// MockClusterDeregistrationClient is a mock of ClusterDeregistrationClient interface.
type MockClusterDeregistrationClient struct {
	ctrl     *gomock.Controller
	recorder *MockClusterDeregistrationClientMockRecorder
}

// MockClusterDeregistrationClientMockRecorder is the mock recorder for MockClusterDeregistrationClient.
type MockClusterDeregistrationClientMockRecorder struct {
	mock *MockClusterDeregistrationClient
}

// NewMockClusterDeregistrationClient creates a new mock instance.
func NewMockClusterDeregistrationClient(ctrl *gomock.Controller) *MockClusterDeregistrationClient {
	mock := &MockClusterDeregistrationClient{ctrl: ctrl}
	mock.recorder = &MockClusterDeregistrationClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterDeregistrationClient) EXPECT() *MockClusterDeregistrationClientMockRecorder {
	return m.recorder
}

// Deregister mocks base method.
func (m *MockClusterDeregistrationClient) Deregister(ctx context.Context, kubeCluster *v1alpha1.KubernetesCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deregister", ctx, kubeCluster)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deregister indicates an expected call of Deregister.
func (mr *MockClusterDeregistrationClientMockRecorder) Deregister(ctx, kubeCluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deregister", reflect.TypeOf((*MockClusterDeregistrationClient)(nil).Deregister), ctx, kubeCluster)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_nodeattestor is a generated GoMock package.
package mock_nodeattestor

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	plugin "github.com/spiffe/spire/proto/common/plugin"
	"github.com/spiffe/spire/proto/server/nodeattestor"
)

// MockNodeAttestor is a mock of NodeAttestor interface
type MockNodeAttestor struct {
	ctrl     *gomock.Controller
	recorder *MockNodeAttestorMockRecorder
}

// MockNodeAttestorMockRecorder is the mock recorder for MockNodeAttestor
type MockNodeAttestorMockRecorder struct {
	mock *MockNodeAttestor
}

// NewMockNodeAttestor creates a new mock instance
func NewMockNodeAttestor(ctrl *gomock.Controller) *MockNodeAttestor {
	mock := &MockNodeAttestor{ctrl: ctrl}
	mock.recorder = &MockNodeAttestorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeAttestor) EXPECT() *MockNodeAttestorMockRecorder {
	return m.recorder
}

// Attest mocks base method
func (m *MockNodeAttestor) Attest(arg0 *nodeattestor.AttestRequest) (*nodeattestor.AttestResponse, error) {
	ret := m.ctrl.Call(m, "Attest", arg0)
	ret0, _ := ret[0].(*nodeattestor.AttestResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Attest indicates an expected call of Attest
func (mr *MockNodeAttestorMockRecorder) Attest(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attest", reflect.TypeOf((*MockNodeAttestor)(nil).Attest), arg0)
}

// Configure mocks base method
func (m *MockNodeAttestor) Configure(arg0 *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error) {
	ret := m.ctrl.Call(m, "Configure", arg0)
	ret0, _ := ret[0].(*plugin.ConfigureResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Configure indicates an expected call of Configure
func (mr *MockNodeAttestorMockRecorder) Configure(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*MockNodeAttestor)(nil).Configure), arg0)
}

// GetPluginInfo mocks base method
func (m *MockNodeAttestor) GetPluginInfo(arg0 *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error) {
	ret := m.ctrl.Call(m, "GetPluginInfo", arg0)
	ret0, _ := ret[0].(*plugin.GetPluginInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPluginInfo indicates an expected call of GetPluginInfo
func (mr *MockNodeAttestorMockRecorder) GetPluginInfo(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPluginInfo", reflect.TypeOf((*MockNodeAttestor)(nil).GetPluginInfo), arg0)
}

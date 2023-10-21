// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/oxyno-zeta/s3-proxy/pkg/s3-proxy/bucket (interfaces: Client)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_Client.go -package=mocks github.com/oxyno-zeta/s3-proxy/pkg/s3-proxy/bucket Client
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	bucket "github.com/oxyno-zeta/s3-proxy/pkg/s3-proxy/bucket"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockClient) Delete(arg0 context.Context, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", arg0, arg1)
}

// Delete indicates an expected call of Delete.
func (mr *MockClientMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClient)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockClient) Get(arg0 context.Context, arg1 *bucket.GetInput) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Get", arg0, arg1)
}

// Get indicates an expected call of Get.
func (mr *MockClientMockRecorder) Get(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClient)(nil).Get), arg0, arg1)
}

// LoadFileContent mocks base method.
func (m *MockClient) LoadFileContent(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadFileContent", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadFileContent indicates an expected call of LoadFileContent.
func (mr *MockClientMockRecorder) LoadFileContent(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadFileContent", reflect.TypeOf((*MockClient)(nil).LoadFileContent), arg0, arg1)
}

// Put mocks base method.
func (m *MockClient) Put(arg0 context.Context, arg1 *bucket.PutInput) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Put", arg0, arg1)
}

// Put indicates an expected call of Put.
func (mr *MockClientMockRecorder) Put(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockClient)(nil).Put), arg0, arg1)
}

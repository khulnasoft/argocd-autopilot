// Code generated by MockGen. DO NOT EDIT.
// Source: ./provider_gitlab.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gitlab "github.com/xanzy/go-gitlab"
)

// MockGitlabClient is a mock of GitlabClient interface.
type MockGitlabClient struct {
	ctrl     *gomock.Controller
	recorder *MockGitlabClientMockRecorder
}

// MockGitlabClientMockRecorder is the mock recorder for MockGitlabClient.
type MockGitlabClientMockRecorder struct {
	mock *MockGitlabClient
}

// NewMockGitlabClient creates a new mock instance.
func NewMockGitlabClient(ctrl *gomock.Controller) *MockGitlabClient {
	mock := &MockGitlabClient{ctrl: ctrl}
	mock.recorder = &MockGitlabClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGitlabClient) EXPECT() *MockGitlabClientMockRecorder {
	return m.recorder
}

// CreateProject mocks base method.
func (m *MockGitlabClient) CreateProject(opt *gitlab.CreateProjectOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Project, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateProject", varargs...)
	ret0, _ := ret[0].(*gitlab.Project)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockGitlabClientMockRecorder) CreateProject(opt interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockGitlabClient)(nil).CreateProject), varargs...)
}

// CurrentUser mocks base method.
func (m *MockGitlabClient) CurrentUser(options ...gitlab.RequestOptionFunc) (*gitlab.User, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CurrentUser", varargs...)
	ret0, _ := ret[0].(*gitlab.User)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CurrentUser indicates an expected call of CurrentUser.
func (mr *MockGitlabClientMockRecorder) CurrentUser(options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentUser", reflect.TypeOf((*MockGitlabClient)(nil).CurrentUser), options...)
}

// ListGroups mocks base method.
func (m *MockGitlabClient) ListGroups(opt *gitlab.ListGroupsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Group, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGroups", varargs...)
	ret0, _ := ret[0].([]*gitlab.Group)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListGroups indicates an expected call of ListGroups.
func (mr *MockGitlabClientMockRecorder) ListGroups(opt interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroups", reflect.TypeOf((*MockGitlabClient)(nil).ListGroups), varargs...)
}
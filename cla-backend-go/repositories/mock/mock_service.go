// Copyright The Linux Foundation and each contributor to CommunityBridge.
// SPDX-License-Identifier: MIT
//

// Code generated by MockGen. DO NOT EDIT.
// Source: repositories/service.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	models "github.com/communitybridge/easycla/cla-backend-go/gen/models"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// AddGithubRepository mocks base method
func (m *MockService) AddGithubRepository(ctx context.Context, externalProjectID string, input *models.GithubRepositoryInput) (*models.GithubRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddGithubRepository", ctx, externalProjectID, input)
	ret0, _ := ret[0].(*models.GithubRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddGithubRepository indicates an expected call of AddGithubRepository
func (mr *MockServiceMockRecorder) AddGithubRepository(ctx, externalProjectID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddGithubRepository", reflect.TypeOf((*MockService)(nil).AddGithubRepository), ctx, externalProjectID, input)
}

// GetRepositoryByName mocks base method
func (m *MockService) GetRepositoryByName(ctx context.Context, repositoryName string) (*models.GithubRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepositoryByName", ctx, repositoryName)
	ret0, _ := ret[0].(*models.GithubRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableRepository mocks base method
func (m *MockService) EnableRepository(ctx context.Context, repositoryID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableRepository", ctx, repositoryID)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableRepositoryWithCLAGroupID mocks base method
func (m *MockService) EnableRepositoryWithCLAGroupID(ctx context.Context, repositoryID, claGroupID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableRepositoryWithCLAGroupID", ctx, repositoryID, claGroupID)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableRepository indicates an expected call of EnableRepository
func (mr *MockServiceMockRecorder) EnableRepository(ctx, repositoryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableRepository", reflect.TypeOf((*MockService)(nil).EnableRepository), ctx, repositoryID)
}

// DisableRepository mocks base method
func (m *MockService) DisableRepository(ctx context.Context, repositoryID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableRepository", ctx, repositoryID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableRepository indicates an expected call of DisableRepository
func (mr *MockServiceMockRecorder) DisableRepository(ctx, repositoryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableRepository", reflect.TypeOf((*MockService)(nil).DisableRepository), ctx, repositoryID)
}

// UpdateClaGroupID mocks base method
func (m *MockService) UpdateClaGroupID(ctx context.Context, repositoryID, claGroupID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClaGroupID", ctx, repositoryID, claGroupID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClaGroupID indicates an expected call of UpdateClaGroupID
func (mr *MockServiceMockRecorder) UpdateClaGroupID(ctx, repositoryID, claGroupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClaGroupID", reflect.TypeOf((*MockService)(nil).UpdateClaGroupID), ctx, repositoryID, claGroupID)
}

// ListProjectRepositories mocks base method
func (m *MockService) ListProjectRepositories(ctx context.Context, externalProjectID string) (*models.ListGithubRepositories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjectRepositories", ctx, externalProjectID)
	ret0, _ := ret[0].(*models.ListGithubRepositories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjectRepositories indicates an expected call of ListProjectRepositories
func (mr *MockServiceMockRecorder) ListProjectRepositories(ctx, externalProjectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectRepositories", reflect.TypeOf((*MockService)(nil).ListProjectRepositories), ctx, externalProjectID)
}

// GetRepository mocks base method
func (m *MockService) GetRepository(ctx context.Context, repositoryID string) (*models.GithubRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepository", ctx, repositoryID)
	ret0, _ := ret[0].(*models.GithubRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepository indicates an expected call of GetRepository
func (mr *MockServiceMockRecorder) GetRepository(ctx, repositoryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepository", reflect.TypeOf((*MockService)(nil).GetRepository), ctx, repositoryID)
}

// DisableRepositoriesByProjectID mocks base method
func (m *MockService) DisableRepositoriesByProjectID(ctx context.Context, projectID string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableRepositoriesByProjectID", ctx, projectID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableRepositoriesByProjectID indicates an expected call of DisableRepositoriesByProjectID
func (mr *MockServiceMockRecorder) DisableRepositoriesByProjectID(ctx, projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableRepositoriesByProjectID", reflect.TypeOf((*MockService)(nil).DisableRepositoriesByProjectID), ctx, projectID)
}

// GetRepositoriesByCLAGroup mocks base method
func (m *MockService) GetRepositoriesByCLAGroup(ctx context.Context, claGroupID string) ([]*models.GithubRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepositoriesByCLAGroup", ctx, claGroupID)
	ret0, _ := ret[0].([]*models.GithubRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoriesByCLAGroup indicates an expected call of GetRepositoriesByCLAGroup
func (mr *MockServiceMockRecorder) GetRepositoriesByCLAGroup(ctx, claGroupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepositoriesByCLAGroup", reflect.TypeOf((*MockService)(nil).GetRepositoriesByCLAGroup), ctx, claGroupID)
}

// GetRepositoriesByOrganizationName mocks base method
func (m *MockService) GetRepositoriesByOrganizationName(ctx context.Context, gitHubOrgName string) ([]*models.GithubRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepositoriesByOrganizationName", ctx, gitHubOrgName)
	ret0, _ := ret[0].([]*models.GithubRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoriesByOrganizationName indicates an expected call of GetRepositoriesByOrganizationName
func (mr *MockServiceMockRecorder) GetRepositoriesByOrganizationName(ctx, gitHubOrgName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepositoriesByOrganizationName", reflect.TypeOf((*MockService)(nil).GetRepositoriesByOrganizationName), ctx, gitHubOrgName)
}

// MockGithubOrgRepo is a mock of GithubOrgRepo interface
type MockGithubOrgRepo struct {
	ctrl     *gomock.Controller
	recorder *MockGithubOrgRepoMockRecorder
}

// MockGithubOrgRepoMockRecorder is the mock recorder for MockGithubOrgRepo
type MockGithubOrgRepoMockRecorder struct {
	mock *MockGithubOrgRepo
}

// NewMockGithubOrgRepo creates a new mock instance
func NewMockGithubOrgRepo(ctrl *gomock.Controller) *MockGithubOrgRepo {
	mock := &MockGithubOrgRepo{ctrl: ctrl}
	mock.recorder = &MockGithubOrgRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGithubOrgRepo) EXPECT() *MockGithubOrgRepoMockRecorder {
	return m.recorder
}

// GetGithubOrganizationByName mocks base method
func (m *MockGithubOrgRepo) GetGithubOrganizationByName(ctx context.Context, githubOrganizationName string) (*models.GithubOrganizations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGithubOrganizationByName", ctx, githubOrganizationName)
	ret0, _ := ret[0].(*models.GithubOrganizations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGithubOrganizationByName indicates an expected call of GetGithubOrganizationByName
func (mr *MockGithubOrgRepoMockRecorder) GetGithubOrganizationByName(ctx, githubOrganizationName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGithubOrganizationByName", reflect.TypeOf((*MockGithubOrgRepo)(nil).GetGithubOrganizationByName), ctx, githubOrganizationName)
}

// GetGithubOrganization mocks base method
func (m *MockGithubOrgRepo) GetGithubOrganization(ctx context.Context, githubOrganizationName string) (*models.GithubOrganization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGithubOrganization", ctx, githubOrganizationName)
	ret0, _ := ret[0].(*models.GithubOrganization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGithubOrganization indicates an expected call of GetGithubOrganization
func (mr *MockGithubOrgRepoMockRecorder) GetGithubOrganization(ctx, githubOrganizationName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGithubOrganization", reflect.TypeOf((*MockGithubOrgRepo)(nil).GetGithubOrganization), ctx, githubOrganizationName)
}

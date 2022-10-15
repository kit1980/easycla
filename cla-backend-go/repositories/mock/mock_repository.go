// Copyright The Linux Foundation and each contributor to CommunityBridge.
// SPDX-License-Identifier: MIT
//

// Code generated by MockGen. DO NOT EDIT.
// Source: repositories/repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	models "github.com/communitybridge/easycla/cla-backend-go/gen/v1/models"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddGithubRepository mocks base method
func (m *MockRepository) GitHubAddRepository(ctx context.Context, externalProjectID, projectSFID string, input *models.GithubRepositoryInput) (*models.GithubRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubAddRepository", ctx, externalProjectID, projectSFID, input)
	ret0, _ := ret[0].(*models.GithubRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddGithubRepository indicates an expected call of AddGithubRepository
func (mr *MockRepositoryMockRecorder) AddGithubRepository(ctx, externalProjectID, projectSFID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubAddRepository", reflect.TypeOf((*MockRepository)(nil).GitHubAddRepository), ctx, externalProjectID, projectSFID, input)
}

// UpdateGithubRepository mocks base method
func (m *MockRepository) GitHubUpdateRepository(ctx context.Context, repositoryID string, input *models.GithubRepositoryInput) (*models.GithubRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubUpdateRepository", ctx, repositoryID, input)
	ret0, _ := ret[0].(*models.GithubRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateGithubRepository indicates an expected call of UpdateGithubRepository
func (mr *MockRepositoryMockRecorder) UpdateGithubRepository(ctx, repositoryID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubUpdateRepository", reflect.TypeOf((*MockRepository)(nil).GitHubUpdateRepository), ctx, repositoryID, input)
}

// UpdateClaGroupID mocks base method
func (m *MockRepository) GitHubUpdateClaGroupID(ctx context.Context, repositoryID, claGroupID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubUpdateClaGroupID", ctx, repositoryID, claGroupID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClaGroupID indicates an expected call of UpdateClaGroupID
func (mr *MockRepositoryMockRecorder) UpdateClaGroupID(ctx, repositoryID, claGroupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubUpdateClaGroupID", reflect.TypeOf((*MockRepository)(nil).GitHubUpdateClaGroupID), ctx, repositoryID, claGroupID)
}

// GitHubEnableRepository mocks base method
func (m *MockRepository) GitHubEnableRepository(ctx context.Context, repositoryID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubEnableRepository", ctx, repositoryID)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableRepository indicates an expected call of EnableRepository
func (mr *MockRepositoryMockRecorder) EnableRepository(ctx, repositoryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubEnableRepository", reflect.TypeOf((*MockRepository)(nil).GitHubEnableRepository), ctx, repositoryID)
}

// EnableRepositoryWithCLAGroupID mocks base method
func (m *MockRepository) GitHubEnableRepositoryWithCLAGroupID(ctx context.Context, repositoryID, claGroupID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubEnableRepositoryWithCLAGroupID", ctx, repositoryID, claGroupID)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableRepositoryWithCLAGroupID indicates an expected call of EnableRepositoryWithCLAGroupID
func (mr *MockRepositoryMockRecorder) EnableRepositoryWithCLAGroupID(ctx, repositoryID, claGroupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubEnableRepositoryWithCLAGroupID", reflect.TypeOf((*MockRepository)(nil).GitHubEnableRepositoryWithCLAGroupID), ctx, repositoryID, claGroupID)
}

// GitHubDisableRepository mocks base method
func (m *MockRepository) GitHubDisableRepository(ctx context.Context, repositoryID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubDisableRepository", ctx, repositoryID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableRepository indicates an expected call of DisableRepository
func (mr *MockRepositoryMockRecorder) DisableRepository(ctx, repositoryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubDisableRepository", reflect.TypeOf((*MockRepository)(nil).GitHubDisableRepository), ctx, repositoryID)
}

// DisableRepositoriesByProjectID mocks base method
func (m *MockRepository) GitHubDisableRepositoriesByProjectID(ctx context.Context, projectID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubDisableRepositoriesByProjectID", ctx, projectID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableRepositoriesByProjectID indicates an expected call of DisableRepositoriesByProjectID
func (mr *MockRepositoryMockRecorder) DisableRepositoriesByProjectID(ctx, projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubDisableRepositoriesByProjectID", reflect.TypeOf((*MockRepository)(nil).GitHubDisableRepositoriesByProjectID), ctx, projectID)
}

// DisableRepositoriesOfGithubOrganization mocks base method
func (m *MockRepository) GitHubDisableRepositoriesOfOrganization(ctx context.Context, externalProjectID, githubOrgName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubDisableRepositoriesOfOrganization", ctx, externalProjectID, githubOrgName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableRepositoriesOfGithubOrganization indicates an expected call of DisableRepositoriesOfGithubOrganization
func (mr *MockRepositoryMockRecorder) DisableRepositoriesOfGithubOrganization(ctx, externalProjectID, githubOrgName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubDisableRepositoriesOfOrganization", reflect.TypeOf((*MockRepository)(nil).GitHubDisableRepositoriesOfOrganization), ctx, externalProjectID, githubOrgName)
}

// GitHubGetRepository mocks base method
func (m *MockRepository) GitHubGetRepository(ctx context.Context, repositoryID string) (*models.GithubRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubGetRepository", ctx, repositoryID)
	ret0, _ := ret[0].(*models.GithubRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepository indicates an expected call of GetRepository
func (mr *MockRepositoryMockRecorder) GetRepository(ctx, repositoryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubGetRepository", reflect.TypeOf((*MockRepository)(nil).GitHubGetRepository), ctx, repositoryID)
}

// GitHubGetRepositoryByName mocks base method
func (m *MockRepository) GitHubGetRepositoryByName(ctx context.Context, repositoryName string) (*models.GithubRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubGetRepositoryByName", ctx, repositoryName)
	ret0, _ := ret[0].(*models.GithubRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GitHubGetRepositoryByExternalID mocks base method
func (m *MockRepository) GitHubGetRepositoryByExternalID(ctx context.Context, repositoryExternalID string) (*models.GithubRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubGetRepositoryByExternalID", ctx, repositoryExternalID)
	ret0, _ := ret[0].(*models.GithubRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoryByName indicates an expected call of GetRepositoryByName
func (mr *MockRepositoryMockRecorder) GetRepositoryByName(ctx, repositoryName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubGetRepositoryByName", reflect.TypeOf((*MockRepository)(nil).GitHubGetRepositoryByName), ctx, repositoryName)
}

// GetRepositoryByGithubID mocks base method
func (m *MockRepository) GitHubGetRepositoryByGithubID(ctx context.Context, externalID string, enabled bool) (*models.GithubRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubGetRepositoryByGithubID", ctx, externalID, enabled)
	ret0, _ := ret[0].(*models.GithubRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoryByGithubID indicates an expected call of GetRepositoryByGithubID
func (mr *MockRepositoryMockRecorder) GetRepositoryByGithubID(ctx, externalID, enabled interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubGetRepositoryByGithubID", reflect.TypeOf((*MockRepository)(nil).GitHubGetRepositoryByGithubID), ctx, externalID, enabled)
}

// GetRepositoriesByCLAGroup mocks base method
func (m *MockRepository) GitHubGetRepositoriesByCLAGroup(ctx context.Context, claGroup string, enabled bool) ([]*models.GithubRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubGetRepositoriesByCLAGroup", ctx, claGroup, enabled)
	ret0, _ := ret[0].([]*models.GithubRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoriesByCLAGroup indicates an expected call of GetRepositoriesByCLAGroup
func (mr *MockRepositoryMockRecorder) GetRepositoriesByCLAGroup(ctx, claGroup, enabled interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubGetRepositoriesByCLAGroup", reflect.TypeOf((*MockRepository)(nil).GitHubGetRepositoriesByCLAGroup), ctx, claGroup, enabled)
}

// GetRepositoriesByOrganizationName mocks base method
func (m *MockRepository) GitHubGetRepositoriesByOrganizationName(ctx context.Context, gitHubOrgName string) ([]*models.GithubRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubGetRepositoriesByOrganizationName", ctx, gitHubOrgName)
	ret0, _ := ret[0].([]*models.GithubRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoriesByOrganizationName indicates an expected call of GetRepositoriesByOrganizationName
func (mr *MockRepositoryMockRecorder) GetRepositoriesByOrganizationName(ctx, gitHubOrgName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubGetRepositoriesByOrganizationName", reflect.TypeOf((*MockRepository)(nil).GitHubGetRepositoriesByOrganizationName), ctx, gitHubOrgName)
}

// GetCLAGroupRepositoriesGroupByOrgs mocks base method
func (m *MockRepository) GitHubGetCLAGroupRepositoriesGroupByOrgs(ctx context.Context, projectID string, enabled bool) ([]*models.GithubRepositoriesGroupByOrgs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubGetCLAGroupRepositoriesGroupByOrgs", ctx, projectID, enabled)
	ret0, _ := ret[0].([]*models.GithubRepositoriesGroupByOrgs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCLAGroupRepositoriesGroupByOrgs indicates an expected call of GetCLAGroupRepositoriesGroupByOrgs
func (mr *MockRepositoryMockRecorder) GetCLAGroupRepositoriesGroupByOrgs(ctx, projectID, enabled interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubGetCLAGroupRepositoriesGroupByOrgs", reflect.TypeOf((*MockRepository)(nil).GitHubGetCLAGroupRepositoriesGroupByOrgs), ctx, projectID, enabled)
}

// GitHubListProjectRepositories mocks base method
func (m *MockRepository) GitHubListProjectRepositories(ctx context.Context, projectSFID string, enabled *bool) (*models.GithubListRepositories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubListProjectRepositories", ctx, projectSFID, enabled)
	ret0, _ := ret[0].(*models.GithubListRepositories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjectRepositories indicates an expected call of ListProjectRepositories
func (mr *MockRepositoryMockRecorder) ListProjectRepositories(ctx, projectSFID, enabled interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubListProjectRepositories", reflect.TypeOf((*MockRepository)(nil).GitHubListProjectRepositories), ctx, projectSFID, enabled)
}

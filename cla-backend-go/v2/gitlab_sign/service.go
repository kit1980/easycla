// Copyright The Linux Foundation and each contributor to CommunityBridge.
// SPDX-License-Identifier: MIT

package gitlab_sign

import (
	"context"
	"encoding/json"

	// "encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"time"

	"github.com/sirupsen/logrus"

	"github.com/communitybridge/easycla/cla-backend-go/events"
	"github.com/communitybridge/easycla/cla-backend-go/gen/v1/models"
	gitlab_api "github.com/communitybridge/easycla/cla-backend-go/gitlab_api"
	log "github.com/communitybridge/easycla/cla-backend-go/logging"
	"github.com/communitybridge/easycla/cla-backend-go/users"
	"github.com/communitybridge/easycla/cla-backend-go/v2/gitlab_organizations"
	"github.com/communitybridge/easycla/cla-backend-go/v2/repositories"
	"github.com/communitybridge/easycla/cla-backend-go/v2/store"
	"github.com/xanzy/go-gitlab"
)

type service struct {
	repoService   repositories.ServiceInterface
	gitlabOrgRepo gitlab_organizations.RepositoryInterface
	userService   users.Service
	gitlabApp     *gitlab_api.App
	storeRepo     store.Repository
}

type Service interface {
	GitlabSignRequest(ctx context.Context, req *http.Request, organizationID, repositoryID, mergeRequestID, contributorConsoleV2Base string, eventService events.Service) (*string, error)
}

func NewService(gitlabRepositoryService repositories.ServiceInterface, gitlabOrgRepository gitlab_organizations.RepositoryInterface, userService users.Service, storeRepo store.Repository, gitlabApp *gitlab_api.App) Service {
	return &service{
		repoService:   gitlabRepositoryService,
		gitlabOrgRepo: gitlabOrgRepository,
		userService:   userService,
		gitlabApp:     gitlabApp,
		storeRepo:     storeRepo,
	}
}

func (s service) GitlabSignRequest(ctx context.Context, req *http.Request, organizationID, repositoryID, mergeRequestID, contributorConsoleV2Base string, eventService events.Service) (*string, error) {
	f := logrus.Fields{
		"functionName":   "v2.gitlab_sign.service.GitlabSignRequest",
		"organizationID": organizationID,
		"repositoryID":   repositoryID,
		"mergeRequestID": mergeRequestID,
	}

	organization, err := s.gitlabOrgRepo.GetGitLabOrganization(ctx, organizationID)
	if err != nil {
		log.WithFields(f).Debugf("unable to get gitlab organiztion by ID: %s, error: %+v ", organizationID, err)
		return nil, err
	}

	if organization.AuthInfo == "" {
		msg := fmt.Sprintf("organization: %s  has no auth details", organizationID)
		log.WithFields(f).Debug(msg)
		return nil, errors.New(msg)
	}
	gitlabClient, err := gitlab_api.NewGitlabOauthClient(organization.AuthInfo, s.gitlabApp)
	if err != nil {
		log.WithFields(f).Debugf("initializaing gitlab client for gitlab org: %s failed: %v", organizationID, err)
		return nil, err
	}

	mergeRequestIDInt, err := strconv.Atoi(mergeRequestID)
	if err != nil {
		log.WithFields(f).Debugf("unable to convert organization string value : %s to Int", organizationID)
		return nil, err
	}

	log.WithFields(f).Debug("Determining return URL from the inbound request ...")
	mergeRequest, _, err := gitlabClient.MergeRequests.GetMergeRequest(repositoryID, mergeRequestIDInt, &gitlab.GetMergeRequestsOptions{})
	if err != nil || mergeRequest == nil {
		log.WithFields(f).Debugf("unable to fetch MR Web URL: mergeRequestID: %s ", mergeRequestID)
		return nil, err
	}

	originURL := mergeRequest.WebURL
	log.WithFields(f).Debugf("Return URL from the inbound request is : %s ", originURL)

	consoleURL, err := s.redirectToConsole(ctx, req, gitlabClient, repositoryID, mergeRequestID, originURL, contributorConsoleV2Base, eventService)
	if err != nil {
		log.WithFields(f).Debug("unable to redirect to contributor console")
		return nil, err
	}

	return consoleURL, nil
}

func (s service) redirectToConsole(ctx context.Context, req *http.Request, gitlabClient *gitlab.Client, repositoryID, mergeRequestID, originURL, contributorBaseURL string, eventService events.Service) (*string, error) {
	f := logrus.Fields{
		"functionName":   "v2.gitlab_sign.service.redirectToConsole",
		"repositoryID":   repositoryID,
		"mergeRequestID": mergeRequestID,
		"originURL":      originURL,
	}

	claUser, err := s.getOrCreateUser(ctx, gitlabClient, eventService)
	if err != nil {
		msg := fmt.Sprintf("unable to get or create user : %+v ", err)
		log.WithFields(f).Warn(msg)
		return nil, err
	}

	repoIDInt, err := strconv.Atoi(repositoryID)
	if err != nil {
		msg := fmt.Sprintf("unable to convert GitlabRepoID: %s to int", repositoryID)
		log.WithFields(f).Warn(msg)
		return nil, err
	}

	log.WithFields(f).Debugf("getting gitlab repository for: %d", repoIDInt)
	gitlabRepo, err := s.repoService.GitLabGetRepositoryByExternalID(ctx, int64(repoIDInt))
	if err != nil {
		msg := fmt.Sprintf("unable to find repository by ID: %s , error: %+v ", repositoryID, err)
		log.WithFields(f).Warn(msg)
		return nil, err
	}

	type StoreValue struct {
		UserID         string `json:"user_id"`
		ProjectID      string `json:"project_id"`
		RepositoryID   string `json:"repository_id"`
		MergeRequestID string `json:"merge_request_id"`
		ReturnURL      string `json:"return_url"`
	}

	log.WithFields(f).Debugf("setting active signature metadata: claUser: %+v, repository: %+v", claUser, gitlabRepo)
	// set active signature metadata to track the user signing process
	key := fmt.Sprintf("active_signature:%s", claUser.UserID)
	storeValue := StoreValue{
		UserID:         claUser.UserID,
		ProjectID:      gitlabRepo.RepositoryClaGroupID,
		RepositoryID:   repositoryID,
		MergeRequestID: mergeRequestID,
		ReturnURL:      originURL,
	}
	json_data, err := json.Marshal(storeValue)
	if err != nil {
		log.Fatal(err)
	}
	expire := time.Now().AddDate(0, 0, 1).Unix()

	// jsonVal, _ := json.Marshal(value)

	err = s.storeRepo.SetActiveSignatureMetaData(ctx, key, expire, string(json_data))
	if err != nil {
		log.WithFields(f).WithError(err).Warn("unable to save signature metadata")
	}

	params := "redirect=" + url.QueryEscape(originURL)
	consoleURL := fmt.Sprintf("https://%s/#/cla/project/%s/user/%s?%s", contributorBaseURL, gitlabRepo.RepositoryClaGroupID, claUser.UserID, params)
	_, err = http.Get(consoleURL)

	if err != nil {
		msg := fmt.Sprintf("unable to redirect to : %s , error: %+v ", consoleURL, err)
		log.WithFields(f).Warn(msg)
		return nil, err
	}

	return &consoleURL, nil
}

func (s service) getOrCreateUser(ctx context.Context, gitlabClient *gitlab.Client, eventsService events.Service) (*models.User, error) {

	f := logrus.Fields{
		"functionName": "v2.gitlab_sign.service.getOrCreateUser",
	}

	gitlabUser, _, err := gitlabClient.Users.CurrentUser()
	if err != nil {
		log.WithFields(f).Debugf("getting gitlab current user for failed : %v ", err)
		return nil, err
	}

	claUser, err := s.userService.GetUserByGitlabID(gitlabUser.ID)
	if err != nil {
		log.WithFields(f).Debugf("unable to get CLA user by github ID: %d , error: %+v ", gitlabUser.ID, err)
		log.WithFields(f).Infof("creating user record for gitlab user : %+v ", gitlabUser)
		user := &models.User{
			GitlabID:       fmt.Sprintf("%d", gitlabUser.ID),
			GitlabUsername: gitlabUser.Username,
			Emails:         []string{gitlabUser.Email},
		}
		claUser, userErr := s.userService.CreateUser(user, nil)
		if err != nil {
			log.WithFields(f).Debugf("unable to create claUser with details : %+v, error: %+v", user, userErr)
			return nil, userErr
		}

		// Log the event
		eventsService.LogEvent(&events.LogEventArgs{
			EventType: events.UserCreated,
			UserID:    user.UserID,
			UserModel: user,
			EventData: &events.UserCreatedEventData{},
		})
		return claUser, nil

	}

	return claUser, nil

}

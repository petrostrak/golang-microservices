package services

import (
	"golang-microservices/src/api/config"
	"golang-microservices/src/api/domain/github"
	"golang-microservices/src/api/domain/repositories"
	"golang-microservices/src/api/providers/github_provider"
	"golang-microservices/src/api/utils/errors"
	"strings"
)

type reposService struct{}

type reposServiceInterface interface {
	CreateRepo(req repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	req := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	res, err := github_provider.CreateRepo(config.GetGithubAccessToken(), req)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		ID:    res.ID,
		Name:  res.Name,
		Owner: res.Owner.Login,
	}

	return &result, nil
}

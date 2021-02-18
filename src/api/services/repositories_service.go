package services

import (
	"fmt"
	"golang-microservices/src/api/config"
	"golang-microservices/src/api/domain/github"
	"golang-microservices/src/api/domain/repositories"
	"golang-microservices/src/api/log"
	"golang-microservices/src/api/providers/github_provider"
	"golang-microservices/src/api/utils/errors"
	"net/http"
	"sync"
)

type reposService struct{}

type reposServiceInterface interface {
	CreateRepo(clientId string, req repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(req []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(clientId string, input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	req := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	log.Info("about to send request to external API", fmt.Sprintf("client_id: %s", clientId), "status:pending")
	res, err := github_provider.CreateRepo(config.GetGithubAccessToken(), req)
	if err != nil {
		log.Info("response obtained from external API", fmt.Sprintf("client_id: %s", clientId), "status:error")
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}
	log.Info("response obtained from external API", fmt.Sprintf("client_id: %s", clientId), "status:success")

	result := repositories.CreateRepoResponse{
		ID:    res.ID,
		Name:  res.Name,
		Owner: res.Owner.Login,
	}

	return &result, nil
}

func (s *reposService) CreateRepos(req []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)

	var wg sync.WaitGroup
	go s.handleRepoResults(&wg, input, output)

	for _, current := range req {
		wg.Add(1)
		go s.createRepoConcurrent(current, input)
	}

	wg.Wait()
	close(input)

	result := <-output

	succeeded := 0
	for _, current := range result.Results {
		if current.Response != nil {
			succeeded++
		}
	}
	if succeeded == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if succeeded == len(req) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}

	return result, nil
}

func (s *reposService) handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse

	for incomingEvent := range input {
		repoResult := repositories.CreateRepositoriesResult{
			Response: incomingEvent.Response,
			Error:    incomingEvent.Error,
		}
		results.Results = append(results.Results, repoResult)
		wg.Done()
	}
	output <- results
}

func (s *reposService) createRepoConcurrent(input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult) {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRepositoriesResult{
			Error: err,
		}
		return
	}

	result, err := s.CreateRepo("TODO_client_id", input)

	if err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}

	output <- repositories.CreateRepositoriesResult{Response: result}
}

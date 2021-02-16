package github_provider

import (
	"encoding/json"
	"fmt"
	"golang-microservices/src/api/clients/restclient"
	"golang-microservices/src/api/domain/github"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo             = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, req github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	res, err := restclient.Post(urlCreateRepo, req, headers)
	if err != nil {
		log.Printf("error while trying to create new repo in github: %s", err.Error())
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid response body"}
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		var errResp github.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errResp); err != nil {
			return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid json response body"}
		}
		errResp.StatusCode = res.StatusCode
		return nil, &errResp
	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Printf("error while trying to unmarshal create repo successful response: %s", err.Error())
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "error while trying to unmarshal github create RepoResponse"}
	}

	return &result, nil
}

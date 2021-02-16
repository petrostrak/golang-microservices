package github_provider

import (
	"errors"
	"fmt"
	"golang-microservices/src/api/clients/restclient"
	"golang-microservices/src/api/domain/github"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("asd123")
	assert.EqualValues(t, "token asd123", header)
}

func TestDefer(t *testing.T) {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	fmt.Println("function's body")
}

func TestCreateRepoErrorRestclient(t *testing.T) {
	restclient.StartMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Err:        errors.New("invalid restclient response"),
	})

	resp, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, resp)
	assert.NotNil(t, err)
}

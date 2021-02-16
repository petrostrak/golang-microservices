package github

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJSON(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "Golang microservices",
		Description: "A golang intro to micorservices",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   false,
		HasProjects: false,
		HasWiki:     true,
	}

	// Marshal takes an input interface and attempts to create a valid json
	// string.
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	fmt.Println(string(bytes))

	assert.EqualValues(t, `{"name":"Golang microservices","description":"A golang intro to micorservices","homepage":"https://github.com","private":true,"has_issues":false,"has_projects":false}`, string(bytes))

	var target CreateRepoRequest

	// Unmarshal takes an input byte array and a POINTER that we are trying to
	// populate using this json.
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)
	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.HasIssues, request.HasIssues)
}

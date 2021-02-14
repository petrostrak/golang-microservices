package services

import (
	"golang-microservices/mvc/dao"
	"golang-microservices/mvc/utils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	userDaoMock     usersDaoMock
	getUserFunction func(userID uint64) (*dao.User, *utils.ApplicationError)
)

func init() {
	dao.UserDao = &usersDaoMock{}
}

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userID uint64) (*dao.User, *utils.ApplicationError) {
	return getUserFunction(userID)
}

func TestGetUserNoUserFoundInDB(t *testing.T) {
	getUserFunction = func(userID uint64) (*dao.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "user 0 does not exists",
		}
	}
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 does not exists", err.Message)
}

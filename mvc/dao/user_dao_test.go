package dao

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := UserDao.GetUser(0)

	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 does not exists", err.Message)

	// if user != nil {
	// 	t.Error("we were not expecting a user with id 0")
	// }
	// if err == nil {
	// 	t.Error("we were expecting an error when if is 0")
	// }
	// if err.StatusCode != http.StatusNotFound {
	// 	t.Error("we were expecting 404 when user is not found")
	// }
}

func TestGetUserNoError(t *testing.T) {
	user, err := UserDao.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "Petros", user.FirstName)
	assert.EqualValues(t, "Trak", user.LastName)
	assert.EqualValues(t, "pit@gmail.com", user.Email)
}

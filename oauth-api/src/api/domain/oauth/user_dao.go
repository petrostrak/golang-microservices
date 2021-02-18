package oauth

import (
	"golang-microservices/src/api/utils/errors"
)

const (
	getUserByUsernamAndPassWord = "SELECT id, username FROM users WHERE username=? AND password=?;"
)

var (
	users = map[string]*User{
		"petros": &User{
			ID:       123,
			Username: "petros",
		},
	}
)

func GetUserByUsernamAndPassWord(username, password string) (*User, errors.ApiError) {
	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundError("no user found")
	}
	return nil, nil
}

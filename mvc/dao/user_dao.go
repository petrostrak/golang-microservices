package dao

import (
	"fmt"
	"golang-microservices/mvc/utils"
	"net/http"
)

var (
	users = map[uint64]*User{
		123: {ID: 1, FirstName: "Petros", LastName: "Trak", Email: "pit@gmail.com"},
	}
)

// GetUser checks the DB for the requested id
func GetUser(userID uint64) (*User, *utils.ApplicationError) {
	// validates if that id exists in DB
	if user := users[userID]; user != nil {
		return user, nil
	}

	// returns result to the service
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}

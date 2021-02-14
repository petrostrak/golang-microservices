package domain

import (
	"fmt"
)

var (
	users = map[uint64]*User{
		123: {ID: 1, FirstName: "Petros", LastName: "Trak", Email: "pit@gmail.com"},
	}
)

// GetUser func checks DB for the userID
func GetUser(userID uint64) (*User, error) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, fmt.Errorf("user %v was not found", userID)

}

package services

import (
	"golang-microservices/mvc/dao"
	"golang-microservices/mvc/utils"
)

// GetUser receives an id of type uint64 as a parameter from the users controller
func GetUser(userID uint64) (*dao.User, *utils.ApplicationError) {
	// calls DAO to check in the DB
	// and returns the user to controller
	return dao.GetUser(userID)
}

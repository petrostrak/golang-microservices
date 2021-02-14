package services

import "golang-microservices/mvc/domain"

func GetUser(userId uint64) (*domain.User, error) {
	return domain.GetUser(userId)
}

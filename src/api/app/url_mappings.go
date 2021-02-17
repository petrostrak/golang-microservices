package app

import (
	"golang-microservices/src/api/controllers/polo"
	"golang-microservices/src/api/controllers/repositories"
)

func mapURLs() {
	router.GET("/marco", polo.Polo)
	router.POST("/repositories", repositories.CreateRepo)
}

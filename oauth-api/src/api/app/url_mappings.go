package app

import (
	"golang-microservices/oauth-api/src/api/controllers/oauth"
	"golang-microservices/src/api/controllers/polo"
)

func mapURLs() {
	router.GET("/marco", polo.Polo)
	router.POST("/oauth/access_token", oauth.CreateAccessToken)
}

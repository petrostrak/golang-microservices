package app

import (
	"golang-microservices/src/api/log"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	log.Info("about to map urls", "step: 1", "status: pending")
	mapURLs()
	log.Info("urls successfully mapped", "step: 2", "status: success")

	if err := router.Run(":8000"); err != nil {
		panic(err)
	}
}

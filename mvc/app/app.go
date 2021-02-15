package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

// StartApp will initialize our app
func StartApp() {
	mapUrls()

	if err := router.Run(":8000"); err != nil {
		panic(err)
	}
}

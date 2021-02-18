package oauth

import (
	"golang-microservices/oauth-api/src/api/domain/oauth"
	"golang-microservices/src/api/utils/errors"

	"github.com/gin-gonic/gin"
)

func CreateAccessToken(c *gin.Context) {
	var request oauth.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}
}

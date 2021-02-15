package controllers

import (
	"golang-microservices/mvc/services"
	"golang-microservices/mvc/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUser will return all users
// GET
func GetUser(c *gin.Context) {
	// take the requested id from the URL query
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	// validate for errors
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a positive number",
			StatusCode: http.StatusNotFound,
			Code:       "bad_request",
		}
		// encode request to json
		utils.RespondError(c, apiErr)
		return
	}

	// now that we have a valid id, we
	// pass the requested id to the service
	user, apiErr := services.UserService.GetUser(uint64(userID))

	// validate for errors from service
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	// return user to client
	utils.Respond(c, http.StatusOK, user)
}

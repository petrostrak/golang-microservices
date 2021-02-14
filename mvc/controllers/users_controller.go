package controllers

import (
	"encoding/json"
	"golang-microservices/mvc/services"
	"golang-microservices/mvc/utils"
	"net/http"
	"strconv"
)

// GetUser will return all users
// GET
func GetUser(w http.ResponseWriter, r *http.Request) {
	// take the requested id from the URL query
	userID, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)

	// validate for errors
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a positive number",
			StatusCode: http.StatusNotFound,
			Code:       "bad_request",
		}

		// encode request to json
		json, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write(json)
		return
	}

	// pass the requested id to the service
	user, apiErr := services.GetUser(uint64(userID))

	// validate for errors from service
	if apiErr != nil {
		json, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write([]byte(json))
		return
	}

	// return user to client
	json, _ := json.Marshal(user)
	w.Write(json)
}

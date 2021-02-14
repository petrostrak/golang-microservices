package controllers

import (
	"encoding/json"
	"golang-microservices/mvc/services"
	"net/http"
	"strconv"
)

// GetUser will return all users
// GET
func GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("user_id must be a positive number"))
		return
	}

	user, err := services.GetUser(uint64(userID))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	// return user to client
	json, _ := json.Marshal(user)
	w.Write(json)
}

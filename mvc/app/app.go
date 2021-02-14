package app

import (
	"golang-microservices/mvc/controllers"
	"net/http"
)

// StartApp will initialize our app
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}

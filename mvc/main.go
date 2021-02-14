package main

import (
	"golang-microservices/mvc/app"
)

// curl localhost:8000/users?user_id=123A -v
func main() {
	app.StartApp()
}

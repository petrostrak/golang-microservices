package main

import (
	"golang-microservices/mvc/app"
)

// curl localhost:8000/users/123 -v
// curl localhost:8000/users/123 -H "Accept:application/xml" -v
func main() {
	app.StartApp()
}

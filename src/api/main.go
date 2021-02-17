package main

import "golang-microservices/mvc/app"

// curl -X GET localhost:8000/marco
// curl -X POST localhost:8000/repositories
func main() {
	app.StartApp()
}

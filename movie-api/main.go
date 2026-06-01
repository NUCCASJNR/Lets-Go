package main

import (
	"movie-api/config"
	"movie-api/routes"
)

func main() {
	config.ConnectDB()
	r := routes.SetupRouter()

	r.Run(":8080")
}
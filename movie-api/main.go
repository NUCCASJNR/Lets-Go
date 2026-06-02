package main

import (
	"movie-api/config"
	"movie-api/models"
	"movie-api/routes"
)

func main() {
	config.ConnectDB()

	// AUTO CREATE TABLE
	config.DB.AutoMigrate(&models.Movie{})

	r := routes.SetupRouter()
	r.Run(":8080")
}
package routes

import (
	"movie-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/movies", controllers.GetMovies)

	return r
}
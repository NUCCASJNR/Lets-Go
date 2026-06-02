package routes

import (
	"movie-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/movies", controllers.GetMovies)
	r.POST("/movies/add", controllers.CreateMovie)
	r.GET("/movies/:id", controllers.GetMovieByID)

	return r
}
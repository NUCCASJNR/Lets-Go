package routes

import (
	"movie-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	movies := r.Group("/movies")
	{
		movies.GET("", controllers.GetMovies)
		movies.POST("/add", controllers.CreateMovie)
		movies.GET("/:id", controllers.GetMovieByID)
		movies.DELETE("/:id", controllers.DeleteMovieByID)
		movies.PUT("/:id", controllers.UpdateMovieByID)
	}

	return r
}
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
		movies.POST("", controllers.CreateMovie)
		movies.GET("/:id", controllers.GetMovieByID)
		movies.PUT("/:id", controllers.UpdateMovieByID)
		movies.DELETE("/:id", controllers.DeleteMovieByID)
	}

	users := r.Group("/users")
	{
		users.POST("/register", controllers.Register)
		users.POST("/login", controllers.Login)
	}

	return r
}
package controllers

import (
	"movie-api/config"
	"movie-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMovies(c *gin.Context) {
	var movies []models.Movie

	config.DB.Find(&movies)

	c.JSON(http.StatusOK, movies)
}

func CreateMovie(c *gin.Context) {
	var movie models.Movie

	// Parse JSON request body
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := config.DB.Create(&movie).Error; err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "Failed to create movie",
		"error":   err.Error(),
	})
	return
}
c.JSON(http.StatusCreated, gin.H{
	"message": "Movie created successfully",
	"data":    movie,
})
}

func GetMovieByID(c *gin.Context) {
	id := c.Param("id")

	movieID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invlid ID",
		})
		return
	}
	var movie models.Movie
	result := config.DB.First(&movie, movieID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Movie not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Movie retrieved successfully",
		"data": movie,
	})
}
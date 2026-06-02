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

func DeleteMovieByID(c *gin.Context) {
	id := c.Param("id")

	movieID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
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

	config.DB.Delete(&movie)

	c.JSON(http.StatusOK, gin.H{
		"message": "Movie deleted successfully",
		// "data":    movie,
	})
}

func UpdateMovieByID(c *gin.Context) {
	id := c.Param("id")

	var movie models.Movie


	// 1. Find existing movie
	result := config.DB.First(&movie, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Movie not found",
		})
		return
	}

	// 2. Create a struct to hold incoming data
	var updatedData models.Movie

	// 3. Read JSON from request body
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 4. Update fields
	movie.Title = updatedData.Title
	movie.Year = updatedData.Year
	movie.Genre = updatedData.Genre
	movie.Rating = updatedData.Rating

	// 5. Save to database
	config.DB.Save(&movie)

	// 6. Return updated movie
	c.JSON(http.StatusOK, gin.H{
		"message": "Movie updated successfully",
		"data":    movie,
	})
}

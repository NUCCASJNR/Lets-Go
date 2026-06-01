package controllers

import (
	"net/http"
	"movie-api/config"
	"movie-api/models"
	"github.com/gin-gonic/gin"
)

func GetMovies(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, title, year, genre, rating FROM movies")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var movies []models.Movie

	for rows.Next() {
		var m models.Movie
		rows.Scan(&m.ID, &m.Title, &m.Year, &m.Genre, &m.Rating)
		movies = append(movies, m)
	}

	c.JSON(http.StatusOK, movies)
}
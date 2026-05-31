package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type Movie struct {
	Title    string	`json:"title"`
	Year     int	`json:"year"`
	Genre    string	`json:"genre"`
	Rating   float64  `json:"rating"`
}

func (m Movie) Summary() {
	fmt.Printf("%s (%d)\n", m.Title, m.Year)
	fmt.Printf("Genre: %s\n", m.Genre)
	fmt.Printf("Rating: %.1f\n", m.Rating)
}

func (m *Movie) UpdateRating(newRating float64) {
	m.Rating = newRating
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT title, year, genre, rating FROM movies")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	movies := []Movie{}

	for rows.Next() {
		var m Movie
		if err := rows.Scan(&m.Title, &m.Year, &m.Genre, &m.Rating); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		movies = append(movies, m)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	var m Movie

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = DB.Exec(
		"INSERT INTO movies(title, year, genre, rating) VALUES(?,?,?,?)",
		m.Title, m.Year, m.Genre, m.Rating,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"message": "Movie added successfully",
		"data":    m,
	}

	json.NewEncoder(w).Encode(response)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Movie API")
}

func startServer() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/movies", getMovies)
	http.HandleFunc("/movies/add", addMovie)

	http.ListenAndServe(":8080", nil)
}
package main

import (
	"encoding/json"
	"net/http"
)

type StudentApi struct {
	Name       string `json:"name"`
	Level      int    `json:"level"`
	Department string `json:"department"`
}

func getStudents(w http.ResponseWriter, r *http.Request) {
	students := []StudentApi{
		{Name: "Areef", Level: 100, Department: "Welding"},
		{Name: "John", Level: 200, Department: "Physics"},
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(students)
}

func PrintStudentApi() {
	http.HandleFunc("/students", getStudents)

	http.ListenAndServe(":8080", nil)
}
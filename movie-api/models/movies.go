package models

type Movie struct {
	ID     uint    `gorm:"primaryKey" json:"id"`
	Title  string  `json:"title" binding:"required"`
	Year   int     `json:"year" binding:"required"`
	Genre  string  `json:"genre" binding:"required"`
	Rating float64 `json:"rating" binding:"required"`
}
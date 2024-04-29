package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	BookID string
	Title string
	Author string
	GenresID uint `gorm:"foreignKey:GenresID"`
	Genres Genres
	Price float64
	Stock_quantity int
	Status string // in / out
}
package models

import "gorm.io/gorm"

type Genres struct {
	gorm.Model
	Name string
}
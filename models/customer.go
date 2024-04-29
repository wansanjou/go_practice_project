package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	CustomerID string
	Name string
	Email string
	Phone string
	Address string
}
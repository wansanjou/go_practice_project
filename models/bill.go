package models

import "gorm.io/gorm"

type Bill struct {
	gorm.Model
	CustomerID uint `gorm:"foreignKey:CustomerID"`
	Customer Customer
	OrderID uint `gorm:"foreignKey:OrderID"`
	Order Order
	TotalPrice int
	Type string // rent / sell
	LoanDate string
	ReturnDate string
}
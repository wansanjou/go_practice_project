package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerID uint `gorm:"foreignKey:CustomerID"`
	Customer Customer
	OrderName string
	OrderDate string
	Type string // rent / sell
}

type OrderDetail struct {
	gorm.Model
	OrderID uint `gorm:"foreignKey:OrderID"`
	Order Order
	BookID uint `gorm:"foreignKey:BookID"`
	Book Book
	Quantity int
	Price float64
	TotalPrice float64
}
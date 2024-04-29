package main

import (
	"github.com/wansanjou/book/initializers"
	"github.com/wansanjou/book/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Book{})
	initializers.DB.AutoMigrate(&models.Genres{})
	initializers.DB.AutoMigrate(&models.Order{})
	initializers.DB.AutoMigrate(&models.OrderDetail{})
	initializers.DB.AutoMigrate(&models.Bill{})
	initializers.DB.AutoMigrate(&models.Customer{})
}
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wansanjou/book/controllers"
	"github.com/wansanjou/book/initializers"
)

func init()  {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()	
	//book
	r.GET("/book" , controllers.SelectBook)
	r.GET("/book/:id" , controllers.BookGetByID)
	r.POST("/book" , controllers.CreateBook)
	r.PUT("/book/:id", controllers.UpdateBook) 
	r.DELETE("/book/:id" , controllers.DeleteBook)

	//genres
	r.GET("/genres" , controllers.SelectGenres)
	r.GET("/genres/:id" , controllers.GenresGetByID)
	r.POST("/genres" , controllers.CreateGenres)
	r.PUT("/genres/:id", controllers.UpdateGenres) 
	r.DELETE("/genres/:id" , controllers.DeleteGenres)

	//customer
	r.GET("/customer" , controllers.SelectCustomer)
	r.GET("/customer/:id" , controllers.CustomerGetByID)
	r.POST("/customer" , controllers.CreateCustomer)
	r.PUT("/customer/:id", controllers.UpdateCustomer) 
	r.DELETE("/customer/:id" , controllers.DeleteCustomer)

	//order
	// r.GET("/order" , controllers.SelectOrder)
	// r.GET("/order/:id" , controllers.OrderGetByID)
	r.POST("/order" , controllers.CreateOrder)
	// r.PUT("/order/:id", controllers.UpdateOrder) 
	// r.DELETE("/order/:id" , controllers.DeleteOrder)
	
	r.Run() // listen and serve on 0.0.0.0:8080
}
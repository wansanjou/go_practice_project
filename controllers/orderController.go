package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wansanjou/book/initializers"
	"github.com/wansanjou/book/models"
)

func OrderGetByID(c *gin.Context) {
	// Get id from URL
	id := c.Param("id")

	// Get the order
	var order models.Order
	if err := initializers.DB.Preload("OrderDetail").First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Order not found"})
		return
	}

	// Response with it
	c.JSON(http.StatusOK, gin.H{
		"Order": order,
	})
}

func CreateOrder(c *gin.Context) {
	// Get data from request body
	var order struct {
		CustomerID uint
		OrderName  string
		OrderDate  string
		Type       string // rent / sell
	}

	c.Bind(&order)
	
	//check id customer
	customer_id := models.Customer{}
	result_customer_id := initializers.DB.First(&customer_id , order.CustomerID)
	if result_customer_id.Error != nil {
		// genres_id not found 
		c.JSON(400, gin.H{"error" : "CustomerID not found"})
		return 
	}


	// Increment order counter
	// var orderCounter uint64 = 0
	// counter := atomic.AddUint64(&orderCounter, 1)

	// // Generate auto-incremental order name
	// orderDate := time.Now().Format("20060102")
	// ordername := fmt.Sprintf("order-%s%04d", orderDate, counter)

neworder := models.Order{
	CustomerID: order.CustomerID,
	OrderName:  order.OrderName,
	OrderDate:  order.OrderDate,
	Type:       order.Type,
}


	result := initializers.DB.Create(&neworder) // Pass pointer of data to Create
	if result.Error != nil {
		c.Status(400)
		return
	}  

	// Return success message
	c.JSON(200, gin.H{
		"Order": "Create Order " + order.OrderName + "success",
	})
}
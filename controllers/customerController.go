package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wansanjou/book/initializers"
	"github.com/wansanjou/book/models"
)

func SelectCustomer(c *gin.Context)  {
	//get all the posts
	var customer []models.Customer
	if err := initializers.DB.Find(&customer).Error ; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error" : "Customer not found" })
		return 
	}

	//response with them
	c.JSON(200, gin.H{
		"Customer": customer,
	})
}

func CustomerGetByID(c *gin.Context)  {
	//get id off url
	id := c.Param("id") 

	//get the posts
	var customer []models.Customer
	if err := initializers.DB.First(&customer , id).Error ; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error" : "Customer not found" })
		return 
	}

	//response with it
	c.JSON(200, gin.H{
		"Customer": customer,
	})
}

func CreateCustomer(c *gin.Context) {
	// Get data from request body
	var customer struct {
		CustomerID string 
		Name string
		Email string
		Phone string
		Address string
	}

	c.Bind(&customer)

	// Create customer

	newcustomer := models.Customer{
		CustomerID:    customer.CustomerID,
		Name:          customer.Name,
		Email:         customer.Email,
		Phone:      	 customer.Phone, 
		Address:       customer.Address,
	}

	result := initializers.DB.Create(&newcustomer) // Pass pointer of data to Create
	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return success message
	c.JSON(200, gin.H{
		"Customer": "Create customer "+customer.Name+" success",
	})
}

func UpdateCustomer(c *gin.Context)  {
	id := c.Param("id") 

	var body struct {
		CustomerID string 
		Name string
		Email string
		Phone string
		Address string
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var customer models.Customer
	if err := initializers.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	if err := initializers.DB.Model(&customer).Updates(models.Customer{
		CustomerID:    body.CustomerID,
		Name:          body.Name,
		Email:         body.Email,
		Phone:      	 body.Phone, 
		Address:       body.Address,
		}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update record"})
		return
	}

	c.JSON(200, gin.H{
		"Customer":  "Update Customer " + body.CustomerID + " successfully",
	})
}

func DeleteCustomer(c *gin.Context) {
	// get id
	id := c.Param("id")

	// delete item
	var customer models.Customer
	if err := initializers.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Customer not found"})
		return
	}

	if err := initializers.DB.Delete(&customer, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete customer"})
		return
	}

	// response with success message
	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted Customer ID: " + id + " successfully",
	})
}
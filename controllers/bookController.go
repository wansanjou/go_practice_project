package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wansanjou/book/initializers"
	"github.com/wansanjou/book/models"
)

func SelectBook(c *gin.Context)  {
	//get all the posts
	var book []models.Book
	if err := initializers.DB.Find(&book).Error ; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error" : "Book not found" })
		return 
	}

	//response with them
	c.JSON(200, gin.H{
		"Book": book,
	})
}

func BookGetByID(c *gin.Context)  {
	//get id off url
	id := c.Param("id") 

	//get the posts
	var book []models.Book
	if err := initializers.DB.First(&book , id).Error ; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error" : "Book not found" })
		return 
	}

	//response with it
	c.JSON(200, gin.H{
		"Book": book,
	})
}

func CreateBook(c *gin.Context) {
	// Get data from request body
	var book struct {
		BookID         string //b0001
		Title          string
		Author         string
		GenresID       uint
		Price          float64
		Stock_quantity int
		Status         string // in / out
	}

	c.Bind(&book)
	
	//check id genres
	genres_id := models.Genres{}
	result_genres_id := initializers.DB.First(&genres_id , book.GenresID)
	if result_genres_id.Error != nil {
		// genres_id not found 
		c.JSON(400, gin.H{"error" : "GenresID not found"})
		return 
	}

	// Create a book
	status := "out"
	if (book.Stock_quantity > 0)  {
		status = "in"
	} else if (book.Stock_quantity < 0 ) {
		c.JSON(404, gin.H{
			"Error": "Quantity must be more than 0 or equal 0",
		})
		return
	}

	newbook := models.Book{
		BookID:         book.BookID,
		Title:          book.Title,
		Author:         book.Author,
		GenresID:       book.GenresID, // Assuming GenresID is provided in the request
		Price:          book.Price,
		Stock_quantity: book.Stock_quantity,
		Status:         status,
	}

	result := initializers.DB.Create(&newbook) // Pass pointer of data to Create
	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return success message
	c.JSON(200, gin.H{
		"Book": "Create book "+book.Title+" success",
	})
}

func UpdateBook(c *gin.Context)  {
	id := c.Param("id") 

	var body struct {
		BookID string
		Title string
		Author string
		GenresID uint `gorm:"foreignKey:GenresID"`
		Price float64
		Stock_quantity int
		Status string // in / out
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book models.Book
	if err := initializers.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	if err := initializers.DB.Model(&book).Updates(models.Book{
		BookID:         body.BookID,
		Title:          body.Title,
		Author:         body.Author,
		GenresID:       body.GenresID, // Assuming GenresID is provided in the request
		Price:          body.Price,
		Stock_quantity: body.Stock_quantity,
		Status:         body.Status,
		}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update record"})
		return
	}

	c.JSON(200, gin.H{
		"Book":  "Update Book " + body.BookID + " successfully",
	})
}

func DeleteBook(c *gin.Context) {
	// get id
	id := c.Param("id")

	// delete item
	var book models.Book
	if err := initializers.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	if err := initializers.DB.Delete(&book, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete record"})
		return
	}

	// response with success message
	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted Book ID: " + id + " successfully",
	})
}
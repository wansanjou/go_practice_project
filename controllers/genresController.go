package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wansanjou/book/initializers"
	"github.com/wansanjou/book/models"
)


func SelectGenres(c *gin.Context)  {
	//get all the posts
	var genres []models.Genres
	if err := initializers.DB.Find(&genres).Error ; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error" : "Genres not found" })
		return 
	}
	

	//response with them
	c.JSON(200, gin.H{
		"Genres": genres,
	})
}

func GenresGetByID(c *gin.Context)  {
	//get id off url
	id := c.Param("id") 

	//get the posts
	var genres []models.Genres
	if err := initializers.DB.First(&genres , id).Error ; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error" : "Genres not found" })
		return 
	}

	//response with it
	c.JSON(200, gin.H{
		"Genres": genres,
	})
}

func CreateGenres(c *gin.Context) {
	// Get data from request body
	var genres struct {
		Name         string
	}

	c.Bind(&genres)

	// Create a genres
	newgenres := models.Genres{
		Name: genres.Name,
	}

	result := initializers.DB.Create(&newgenres) // Pass pointer of data to Create
	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return success message
	c.JSON(200, gin.H{
		"Genres":  "Create Genres " + newgenres.Name + " success",
	})
}


func UpdateGenres(c *gin.Context)  {
	//Get the id off the url
	id := c.Param("id") 

	//Get the data off req body
	var body struct{
		Name string
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	

	//Find the post were updating
	var genres models.Genres
	if err := initializers.DB.First(&genres, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	//Update it 
	if err := initializers.DB.Model(&genres).Updates(models.Genres{Name: body.Name}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update record"})
		return
	}

	//response with it
	c.JSON(200, gin.H{
		"Genres":  "Update Genres " + body.Name + " successfully",
	})

}

func DeleteGenres(c *gin.Context)  {
	//get id
	id := c.Param("id") 

	//delete item
	var genres models.Genres
	if err := initializers.DB.First(&genres, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	if err:= initializers.DB.Delete(&genres , id).Error ;err != nil {
		c.JSON(http.StatusNotFound , gin.H{"error": "Failed to Delete record"})
		return 
	}
	

	//response with it
	c.JSON(200, gin.H{
		"Genres":  "Delete Genres ID:" + id + " successfully",
	})
}
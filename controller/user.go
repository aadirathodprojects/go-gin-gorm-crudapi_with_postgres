package controller

import (
	"go-gin-gorm/config"
	"go-gin-gorm/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

	// Bind JSON from request to user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save to database
	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Return created user
	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {
	var users []models.User

	result := config.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		// If user not found
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	var updatedData models.User
	if err := c.BindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Name = updatedData.Name
	user.Email = updatedData.Email
	user.Password = updatedData.Password

	config.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User updated Successfully", "data": user})
}

func DeleteUser(c *gin.Context) {
	idparam := c.Param("id")
	id, err := strconv.Atoi(idparam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid user ID"})
	}
	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "ID not exist"})
	}
	config.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "successful deletion"})
}

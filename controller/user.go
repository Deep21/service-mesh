package controller

import (
	"net/http"

	"example.com/m/models"
	"github.com/gin-gonic/gin"
)

func NewUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&models.User{Name: user.Name, Email: user.Email})
}

func AllUsers() *[]models.User {
	var users []models.User
	models.DB.Find(&users)
	return &users
}

func FindUser() *models.User {
	var user models.User
	models.DB.Find(&user)
	return &user
}

func DeleteUser(c *gin.Context) {
	// Get model if exist
	var u models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&u)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

package controllers

import (
	"net/http"
	"taptoeat-be/models"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var users []models.User

	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"user": users})
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.ShouldBindJSON(&user)
	models.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfuly to created",
		"code":    http.StatusCreated,
	})
}

func PutUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := models.DB.Where("id", id).Find(&user); err.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Upss sorry, user with id " + id + " not found",
			"code":    http.StatusNotFound,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	c.ShouldBindJSON(&user)
	models.DB.Model(&user).Where("id", id).Updates(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success updated user",
		"code":    http.StatusOK,
	})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := models.DB.Where("id", id).Delete(&user, id); err.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Upss sorry, user with id " + id + " not found",
			"code":    http.StatusNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfuly to delete user with id " + id,
		"code":    http.StatusOK,
	})
}

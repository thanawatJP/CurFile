package handler

import (
	"curfile/database"
	"curfile/database/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUserHandler(c *gin.Context) {
	var newUser models.UserAuth

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var role models.Role
	if err := database.DB.First(&role, newUser.RoleID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role not found"})
		return
	}

	newUser.Role = role

	result := database.DB.Create(&newUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

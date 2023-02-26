package controllers

import (
	"gin-training/models"
	service "gin-training/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var user models.User

	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(strings.TrimSpace(input.Email)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The Email field is required.!"})
		return
	}

	if len(strings.TrimSpace(input.Password)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The Password field is required.!"})
		return
	}

	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user " + input.Email + " not found!"})
		return
	}

	if user.Password != input.Password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect password!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": service.JWTAuthService().GenerateToken(user.Email, true)})

}

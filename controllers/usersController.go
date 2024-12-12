package controllers

import (
	"github.com/cenpnt/Go-Gorm-PostgreSQL/initializers"
	"github.com/cenpnt/Go-Gorm-PostgreSQL/models"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

func UserCreate(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "error": "Invalid request data"})
		return
	}

	if err := initializers.DB.Where("email = ?", user.Email).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash the password"})
        return
	}

	user.Password = string(hashedPassword)

	if err := initializers.DB.Create(&user).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
        return
	}

	c.IndentedJSON(http.StatusCreated, user)
}

func GetUsers(c *gin.Context) {
	var users []models.User

	if err := initializers.DB.Preload("Posts").Find(&users).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Failed to get users"})
	}

	c.IndentedJSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := initializers.DB.Preload("Posts").First(&user, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}
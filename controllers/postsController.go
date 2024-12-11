package controllers

import (
	"github.com/cenpnt/Go-Gorm-PostgreSQL/initializers"
	"github.com/cenpnt/Go-Gorm-PostgreSQL/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func PostsCreate(c *gin.Context) {
	var post models.Post

	if err := c.BindJSON(&post); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "error": "Invalid request data"})
		return
	}

	if post.Title == "" || post.Body == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Title and Body are required"})
		return
	}

	result := initializers.DB.Create(&post)
	
	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "error": "Failed to create post"})
		return
	}
	c.IndentedJSON(http.StatusCreated, post)
}

func GetPosts(c *gin.Context) {
	var posts []models.Post

	initializers.DB.Find(&posts)

	c.IndentedJSON(http.StatusOK, posts)
}

func GetPostByID(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.IndentedJSON(http.StatusNotFound, gin.H{ "error": "Post not found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
            "error": "Server error",
        })
        return
	}

	c.IndentedJSON(http.StatusOK, post)
}
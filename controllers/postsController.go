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

    if err := initializers.DB.First(&post, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Post not found"})
            return
        }
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve post"})
        return
    }

	c.IndentedJSON(http.StatusOK, post)
}

func PostUpdates(c *gin.Context) {
	id := c.Param("id")
	var updateData models.Post

	if err := c.BindJSON(&updateData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

    var post models.Post
    if err := initializers.DB.First(&post, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Post not found"})
            return
        }
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve post"})
        return
    }

	updateResult := initializers.DB.Model(&post).Updates(updateData)

	if updateResult.Error != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
        return
    }

	var updatedPost models.Post
    if err := initializers.DB.First(&updatedPost, id).Error; err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated post"})
        return
    }

    c.IndentedJSON(http.StatusOK, updatedPost)
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	if err := initializers.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Post not found"})
            return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve post"})
        return
	}

	if err := initializers.DB.Delete(&post).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{ "error": "Failed to delete post" })
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{ "message": "Data has been deleted" })
}
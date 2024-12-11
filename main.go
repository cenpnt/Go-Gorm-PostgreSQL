package main

import (
	"github.com/cenpnt/Go-Gorm-PostgreSQL/controllers"
	"github.com/cenpnt/Go-Gorm-PostgreSQL/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPostByID)
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostUpdates)

	r.Run()
}
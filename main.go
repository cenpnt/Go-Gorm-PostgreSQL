package main

import (
	"github.com/cenpnt/Go-Gorm-PostgreSQL/controllers"
	"github.com/cenpnt/Go-Gorm-PostgreSQL/initializers"
	"github.com/cenpnt/Go-Gorm-PostgreSQL/middleware"
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
	r.POST("/posts", middleware.AuthMiddleware() ,controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostUpdates)
	r.DELETE("/posts/:id", controllers.PostDelete)

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUserByID)

	r.Run()
}
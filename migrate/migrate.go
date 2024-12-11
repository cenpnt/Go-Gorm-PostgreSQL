package main

import (
	"github.com/cenpnt/Go-Gorm-PostgreSQL/initializers"
	"github.com/cenpnt/Go-Gorm-PostgreSQL/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
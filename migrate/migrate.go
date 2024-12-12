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
	if err := initializers.DB.AutoMigrate(&models.User{}, &models.Post{}); err != nil {
		panic("Migration failed: " + err.Error())
	}
}
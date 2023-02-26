package main

import (
	"gin-training/models"
	"gin-training/routes"
)

func main() {
	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.User{})
	r := routes.SetupRoutes(db)
	r.Run("localhost:8000")
}

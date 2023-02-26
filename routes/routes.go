package routes

import (
	"gin-training/controllers"
	"gin-training/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/tasks", middleware.AuthorizeJWT(), controllers.FindTasks)
	r.POST("/tasks", middleware.AuthorizeJWT(), controllers.CreateTask)
	r.GET("/tasks/:id", middleware.AuthorizeJWT(), controllers.FindTask)
	r.PATCH("/tasks/:id", middleware.AuthorizeJWT(), controllers.UpdateTask)
	r.DELETE("tasks/:id", middleware.AuthorizeJWT(), controllers.DeleteTask)

	r.POST("/login", controllers.Login)
	return r
}

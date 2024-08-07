package router

import (
	"tasker/controllers"
	"tasker/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/api")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		// Example task endpoint
		auth.GET("/users", middleware.AdminOnlyMiddleware(), controllers.GetUsers)
		auth.POST("/tasks", controllers.CreateTask)
		auth.GET("/tasks", controllers.GetTasks)
		auth.GET("/tasks/:id", controllers.GetTaskByID)
		auth.PUT("/tasks/:id", controllers.UpdateTask)
		auth.DELETE("/tasks/:id", controllers.DeleteTask)
		

	}

	return r
}

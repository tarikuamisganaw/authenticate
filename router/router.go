package router

import (
	"tasker/controllers"
	"tasker/middleware"

	"github.com/gin-gonic/gin"
)

// router/router.go
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public routes
	r.POST("/api/register", controllers.Register)
	r.POST("/api/login", controllers.Login)

	// Protected routes
	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/tasks", controllers.CreateTask)
		auth.GET("/tasks", controllers.GetTasks)
		auth.GET("/tasks/:id", controllers.GetTaskByID)
		auth.PUT("/tasks/:id", controllers.UpdateTask)
		auth.DELETE("/tasks/:id", controllers.DeleteTask)

		// Admin only route
		admin := auth.Group("/admin")
		admin.Use(middleware.AdminMiddleware())
		{
			admin.GET("/users", controllers.GetAllUsers)
		}
	}

	return r
}

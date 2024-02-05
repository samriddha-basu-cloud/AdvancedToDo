package main

import (
	"github.com/gin-gonic/gin"
	"github.com/samriddha-basu-cloud/AdvancedToDo/api/handlers"
	"github.com/samriddha-basu-cloud/AdvancedToDo/cron"
	"github.com/samriddha-basu-cloud/AdvancedToDo/db"
	"github.com/samriddha-basu-cloud/AdvancedToDo/middleware"
)

func main() {
	r := gin.Default()

	// Initialize MongoDB
	db.InitMongoDB()

	// Middleware
	r.Use(middleware.AuthenticateJWT)

	// Routes
	v1 := r.Group("/api/v1")
	{
		v1.POST("/tasks", handlers.CreateTask)
		v1.POST("/subtasks", handlers.CreateSubTask)
		v1.GET("/tasks", handlers.GetAllUserTasks)
		v1.GET("/subtasks", handlers.GetAllUserSubTasks)
		v1.PUT("/tasks/:id", handlers.UpdateTask)
		v1.PUT("/subtasks/:id", handlers.UpdateSubTask)
		v1.DELETE("/tasks/:id", handlers.DeleteTask)
		v1.DELETE("/subtasks/:id", handlers.DeleteSubTask)
	}

	// Start server
	r.Run(":8080")

	// Schedule cron jobs
	cron.ScheduleTaskPriorityUpdate()
	cron.ScheduleVoiceCalling()
}

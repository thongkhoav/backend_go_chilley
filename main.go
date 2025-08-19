package main

import (
	"todo-backend/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/tasks", handlers.GetTasks)
	r.POST("/tasks", handlers.AddTask)
	r.Run(":8080")
}

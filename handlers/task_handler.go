package handlers

import (
	"net/http"
	"todo-backend/models"
	"todo-backend/storage"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	tasks := storage.GetTasks()
	c.JSON(http.StatusOK, tasks)
}

func AddTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTask := storage.AddTask(task)
	c.JSON(http.StatusCreated, gin.H{"id": newTask.ID})
}

package handlers

import (
	"net/http"
	"strconv"
	"todo-backend/models"
	"todo-backend/storage"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	tasks := storage.GetTasks()
	c.JSON(http.StatusOK, tasks)
}

func AddTask(c *gin.Context) {
	var taskDto models.CreateTaskRequest
	if err := c.ShouldBindJSON(&taskDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTask := storage.AddTask(taskDto)
	c.JSON(http.StatusCreated, gin.H{"id": newTask.ID})
}

func UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	// No id is passed
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var body models.UpdateTaskRequest

	// Check if the request body is valid
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the task in storage
	if body.Completed == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "completed field is required"})
		return
	}
	if err := storage.UpdateTask(id, *body.Completed); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	// No id is passed
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// Delete task from storage
	if err := storage.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

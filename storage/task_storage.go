package storage

import (
	"todo-backend/models"
)

var tasks []models.Task
var nextID = 1

func GetTasks() []models.Task {
	return tasks
}

func AddTask(task models.Task) models.Task {
	task.ID = nextID
	nextID++
	tasks = append(tasks, task)
	return task
}

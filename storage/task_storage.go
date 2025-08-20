package storage

import (
	"errors"
	"todo-backend/models"
)

var tasks []models.Task
var nextID = 1

func GetTasks() []models.Task {
	return tasks
}

func AddTask(task models.CreateTaskRequest) models.Task {
	newTask := models.Task{
		ID:          nextID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   false,
	}
	nextID++
	tasks = append(tasks, newTask)
	return newTask
}

func UpdateTask(id int, completed bool) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = completed
			return nil
		}
	}
	return errors.New("task not found")
}

func DeleteTask(id int) error {
	for i, task := range tasks {
		if task.ID == id {
			// Append next task to the current task
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

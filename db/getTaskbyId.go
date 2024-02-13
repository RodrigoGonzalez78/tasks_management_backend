package db

import "github.com/RodrigoGonzalez78/tasks_management_backend/models"

func GetTaskById(id string) models.Task {

	var task models.Task
	dbConnection.First(&task, id)
	return task
}

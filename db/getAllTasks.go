package db

import "github.com/RodrigoGonzalez78/tasks_management_backend/models"

func GetAllTaks() []models.Task {
	var tasks []models.Task
	dbConnection.Find(&tasks)
	return tasks
}

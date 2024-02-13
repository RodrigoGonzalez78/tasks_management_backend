package db

import "github.com/RodrigoGonzalez78/tasks_management_backend/models"

func GetAllUsers() []models.Task {
	var tasks []models.Task
	dbConnection.Find(&tasks)
	return tasks
}

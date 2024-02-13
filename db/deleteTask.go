package db

import "github.com/RodrigoGonzalez78/tasks_management_backend/models"

func DeleteTask(task *models.Task) {
	dbConnection.Unscoped().Delete(task)
}

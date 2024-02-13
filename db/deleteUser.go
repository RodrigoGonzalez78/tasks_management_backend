package db

import "github.com/RodrigoGonzalez78/tasks_management_backend/models"

func DeleteUser(user models.User) {
	dbConnection.Unscoped().Delete(&user)
}

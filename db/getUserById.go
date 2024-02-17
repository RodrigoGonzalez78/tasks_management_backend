package db

import "github.com/RodrigoGonzalez78/tasks_management_backend/models"

func GetUserById(id uint) models.User {
	var user models.User
	dbConnection.First(&user, id)
	return user
}

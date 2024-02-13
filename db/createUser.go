package db

import "github.com/RodrigoGonzalez78/tasks_management_backend/models"

func CreateUser(user *models.User) error {
	if err := dbConnection.Create(user).Error; err != nil {
		// Ocurrió un error al crear el usuario
		return err
	}
	// Usuario creado exitosamente
	return nil
}

package db

import "github.com/RodrigoGonzalez78/tasks_management_backend/models"

// UpdateUser actualiza los datos de un usuario en la base de datos.
func UpdateUser(user *models.User) error {
	if err := dbConnection.Save(user).Error; err != nil {
		return err // Ocurrió un error al actualizar el usuario
	}
	return nil // Usuario actualizado exitosamente
}

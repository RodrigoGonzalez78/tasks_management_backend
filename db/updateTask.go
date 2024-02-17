package db

import "github.com/RodrigoGonzalez78/tasks_management_backend/models"

// UpdateTask actualiza los datos de una tarea en la base de datos.
func UpdateTask(task *models.Task) error {
	if err := dbConnection.Save(task).Error; err != nil {
		return err // Ocurrió un error al actualizar la tarea
	}
	return nil // Tarea actualizada exitosamente
}

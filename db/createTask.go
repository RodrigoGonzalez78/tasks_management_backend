package db

import "github.com/RodrigoGonzalez78/tasks_management_backend/models"

func CreateTask(task *models.Task) error {
	if err := dbConnection.Create(task).Error; err != nil {
		// Ocurrió un error al crear la tarea
		return err
	}
	// Tarea creado exitosamente
	return nil
}

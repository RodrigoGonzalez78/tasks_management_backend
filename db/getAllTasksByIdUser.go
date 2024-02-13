package db

import "github.com/RodrigoGonzalez78/tasks_management_backend/models"

func GetAllTasksByUserId(userId uint) []models.Task {
	var tasks []models.Task
	// Consulta todas las tareas que tienen el UserId igual al proporcionado
	if err := dbConnection.Where("user_id = ?", userId).Find(&tasks).Error; err != nil {
		// Manejo de errores, puedes personalizarlo según tus necesidades
		panic(err)
	}
	return tasks
}

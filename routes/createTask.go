package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/tasks_management_backend/db"
	"github.com/RodrigoGonzalez78/tasks_management_backend/jwtMetods"
	"github.com/RodrigoGonzalez78/tasks_management_backend/models"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud: "+err.Error(), http.StatusBadRequest)
		return
	}

	if task.Title == "" || task.Description == "" {
		http.Error(w, "El título y la descripción son obligatorios", http.StatusBadRequest)
		return
	}

	task.UserId = jwtMetods.IDUser
	err = db.CreateTask(&task)
	if err != nil {
		http.Error(w, "Error al crear la tarea: "+err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&task)
}

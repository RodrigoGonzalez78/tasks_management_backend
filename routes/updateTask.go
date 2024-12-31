package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/tasks_management_backend/db"
	"github.com/RodrigoGonzalez78/tasks_management_backend/models"
	"github.com/gorilla/mux"
)

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	task := db.GetTaskById(params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var updatedTask models.Task
	err := json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud: "+err.Error(), http.StatusBadRequest)
		return
	}

	if updatedTask.Title == "" || updatedTask.Description == "" {
		http.Error(w, "El título y la descripción son obligatorios", http.StatusBadRequest)
		return
	}

	task.Title = updatedTask.Title
	task.Description = updatedTask.Description

	err = db.UpdateTask(&task)
	if err != nil {
		http.Error(w, "Error al actualizar la tarea: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(task)
}

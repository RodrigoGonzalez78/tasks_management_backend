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

	// Extraer el ID de la tarea de los parámetros de la URL
	params := mux.Vars(r)
	task := db.GetTaskById(params["id"])

	//Verificamos si existe el id en la tabla
	//Golang devuelve 0 por defecto, es decir todos los campos con ZERO value
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Decodificar la nueva información de la tarea del cuerpo de la solicitud
	var updatedTask models.Task
	err := json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Actualizar la información de la tarea
	task.Title = updatedTask.Title
	task.Description = updatedTask.Description
	// También puedes actualizar otros campos según sea necesario

	// Guardar la tarea actualizada en la base de datos
	err = db.UpdateTask(&task)

	if err != nil {
		http.Error(w, "Error al actualizar la tarea: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder con la tarea actualizada
	json.NewEncoder(w).Encode(task)
}

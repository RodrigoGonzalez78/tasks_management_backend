package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/tasks_management_backend/db"
	"github.com/RodrigoGonzalez78/tasks_management_backend/models"
	"github.com/gorilla/mux"
)

func GetTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	//Extraemos el parametro que nos indica el id de usuario
	params := mux.Vars(r)

	task = db.GetTaskById(params["id"])

	//Verificamos si existe el id en la tabla
	//Golang devuelve 0 por defecto, es decir todos los campos con ZERO value
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&task)
}

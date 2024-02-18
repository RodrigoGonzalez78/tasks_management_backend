package routes

import (
	"net/http"

	"github.com/RodrigoGonzalez78/tasks_management_backend/db"
	"github.com/RodrigoGonzalez78/tasks_management_backend/models"
	"github.com/gorilla/mux"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	//Extraemos el parametro que nos indica el id de usuario
	params := mux.Vars(r)

	task = db.GetTaskById(params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	db.DeleteTask(&task)
	w.WriteHeader(http.StatusOK)
}

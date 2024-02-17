package routes

import (
	"net/http"

	"github.com/RodrigoGonzalez78/tasks_management_backend/db"
	"github.com/RodrigoGonzalez78/tasks_management_backend/models"
	"github.com/gorilla/mux"
)

func DeleteTaskById(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	//Extraemos el parametro que nos indica el id de usuario
	params := mux.Vars(r)

	task = db.GetTaskById(params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//Cambia el valor de deleted_at, no elmina el elemento en si
	//db.DB.Delete(&task) //igual la libreria se encarga de no mostrar mas el elemento

	//Remueve totalamente de la tabla
	db.DeleteTask(&task)
	w.WriteHeader(http.StatusOK)
}

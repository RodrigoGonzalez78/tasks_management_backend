package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/tasks_management_backend/db"
	"github.com/RodrigoGonzalez78/tasks_management_backend/models"
	"github.com/gorilla/mux"
)

func GetAllTaks(w http.ResponseWriter, r *http.Request) {
	tasks := db.GetAllTaks()
	json.NewEncoder(w).Encode(&tasks)
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
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

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	json.NewDecoder(r.Body).Decode(&task)

	err := db.CreateTask(&task)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

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

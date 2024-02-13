package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/tasks_management_backend/db"
	"github.com/RodrigoGonzalez78/tasks_management_backend/models"
	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := db.GetAllUsers()
	json.NewEncoder(w).Encode(&users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	var user models.User

	//Extraemos el parametro que nos indica el id de usuario
	params := mux.Vars(r)
	user = db.GetUserById(params["id"])

	//Verificamos si existe el id en la tabla
	//Golang devuelve 0 por defecto, es decir todos los campos con ZERO value
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUserAccoutById(w http.ResponseWriter, r *http.Request) {
	var user models.User
	//Extraemos el parametro que nos indica el id de usuario
	params := mux.Vars(r)

	user = db.GetUserById(params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//Borramos las notas de la base de datos pertenecientes al usuario

	var tasks []models.Task = db.GetAllTasksByUserId(user.ID)

	for _, v := range tasks {
		db.DeleteTask(&v)
	}

	//Cambia el valor de deleted_at, no elmina el elemento en si
	//db.DB.Delete(&user) igual la libreria se encarga de no mostrar mas el elemento

	//Remueve totalamente de la tabla
	db.DeleteUser(user)

	w.WriteHeader(http.StatusOK)
}

package routes

import (
	"net/http"

	"github.com/RodrigoGonzalez78/tasks_management_backend/db"
	"github.com/RodrigoGonzalez78/tasks_management_backend/jwtMetods"
	"github.com/RodrigoGonzalez78/tasks_management_backend/models"
)

func DeleteUserAccout(w http.ResponseWriter, r *http.Request) {

	var user models.User = db.GetUserById(jwtMetods.IDUser)

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

package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/tasks_management_backend/db"
	"github.com/RodrigoGonzalez78/tasks_management_backend/jwtMetods"
	"github.com/RodrigoGonzalez78/tasks_management_backend/models"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := db.GetUserById(jwtMetods.IDUser)

	//Verificamos si existe el id en la tabla
	//Golang devuelve 0 por defecto, es decir todos los campos con ZERO value
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Decodificar la nueva información del usuario del cuerpo de la solicitud
	var updatedUser models.User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Actualizar la información del usuario
	user.FirstName = updatedUser.FirstName
	user.LastName = updatedUser.LastName
	user.Email = updatedUser.Email
	user.Password = updatedUser.Password
	// También puedes actualizar otros campos según sea necesario

	// Guardar el usuario actualizado en la base de datos
	err = db.UpdateUser(&user)

	if err != nil {
		http.Error(w, "Error al actualizar el usuario: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder con el usuario actualizado
	json.NewEncoder(w).Encode(user)
}

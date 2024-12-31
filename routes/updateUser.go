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

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud: "+err.Error(), http.StatusBadRequest)
		return
	}

	if user.FirstName == "" || user.LastName == "" {
		http.Error(w, "Todos los campos son obligatorios", http.StatusBadRequest)
		return
	}

	existingUser := db.GetUserById(jwtMetods.IDUser)

	if existingUser.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	existingUser.FirstName = user.FirstName
	existingUser.LastName = user.LastName

	if err := db.UpdateUser(&existingUser); err != nil {
		http.Error(w, "Error al actualizar el usuario: "+err.Error(), http.StatusInternalServerError)
		return
	}

	existingUser.Password = ""

	json.NewEncoder(w).Encode(existingUser)
}

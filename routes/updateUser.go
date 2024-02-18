package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/tasks_management_backend/db"
	"github.com/RodrigoGonzalez78/tasks_management_backend/jwtMetods"
	"github.com/RodrigoGonzalez78/tasks_management_backend/models"
)

// UpdateUserData actualiza los datos del usuario.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Verificar que los campos no estén vacíos
	if user.FirstName == "" || user.LastName == "" {
		http.Error(w, "Todos los campos son obligatorios", http.StatusBadRequest)
		return
	}

	// Verificar si el usuario existe en la base de datos
	existingUser := db.GetUserById(jwtMetods.IDUser)

	if existingUser.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Actualizar los datos del usuario
	existingUser.FirstName = user.FirstName
	existingUser.LastName = user.LastName

	// Guardar los datos actualizados del usuario en la base de datos
	if err := db.UpdateUser(&existingUser); err != nil {
		http.Error(w, "Error al actualizar el usuario: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder con el usuario actualizado
	json.NewEncoder(w).Encode(existingUser)
}

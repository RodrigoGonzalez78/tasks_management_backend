package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/tasks_management_backend/db"
	"github.com/RodrigoGonzalez78/tasks_management_backend/jwtMetods"
	"github.com/RodrigoGonzalez78/tasks_management_backend/models"
	"golang.org/x/crypto/bcrypt"
)

// UpdateUserPassword actualiza la contraseña del usuario.
func UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newUserPassword models.User
	err := json.NewDecoder(r.Body).Decode(&newUserPassword)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Verificar si el usuario existe en la base de datos
	existingUser := db.GetUserById(jwtMetods.IDUser)

	if existingUser.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Generar el hash de la nueva contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUserPassword.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error al generar el hash de la nueva contraseña: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Actualizar la contraseña del usuario
	existingUser.Password = string(hashedPassword)

	// Guardar la nueva contraseña en la base de datos
	if err := db.UpdateUser(&existingUser); err != nil {
		http.Error(w, "Error al actualizar la contraseña del usuario: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder con un mensaje de éxito
	response := map[string]string{"message": "Contraseña actualizada correctamente"}
	json.NewEncoder(w).Encode(response)
}

package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/tasks_management_backend/db"
	"github.com/RodrigoGonzalez78/tasks_management_backend/jwtMetods"
	"github.com/RodrigoGonzalez78/tasks_management_backend/models"
	"golang.org/x/crypto/bcrypt"
)

func UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newUserPassword models.User
	err := json.NewDecoder(r.Body).Decode(&newUserPassword)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud: "+err.Error(), http.StatusBadRequest)
		return
	}

	existingUser := db.GetUserById(jwtMetods.IDUser)

	if existingUser.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUserPassword.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error al generar el hash de la nueva contraseña: "+err.Error(), http.StatusInternalServerError)
		return
	}

	existingUser.Password = string(hashedPassword)

	if err := db.UpdateUser(&existingUser); err != nil {
		http.Error(w, "Error al actualizar la contraseña del usuario: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Contraseña actualizada correctamente"}
	json.NewEncoder(w).Encode(response)
}

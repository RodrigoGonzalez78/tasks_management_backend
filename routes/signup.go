package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/tasks_management_backend/db"
	"github.com/RodrigoGonzalez78/tasks_management_backend/models"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Verificar si el correo electrónico ya está en uso
	usedEmail, err := db.CheckExistUser(user.Email)
	if err != nil {
		http.Error(w, "Error al verificar el correo electrónico: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if usedEmail {
		http.Error(w, "El correo electrónico ya está en uso", http.StatusBadRequest)
		return
	}

	// Crear el usuario en la base de datos
	if err := db.CreateUser(&user); err != nil {
		http.Error(w, "Error al crear el usuario: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Devolver el usuario creado como respuesta
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

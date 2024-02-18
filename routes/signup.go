package routes

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/RodrigoGonzalez78/tasks_management_backend/db"
	"github.com/RodrigoGonzalez78/tasks_management_backend/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Verificar el formato del correo electrónico
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(user.Email) {
		http.Error(w, "Formato de correo electrónico inválido", http.StatusBadRequest)
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

	// Verificar si la contraseña cumple con los criterios (por ejemplo, longitud mínima)
	if len(user.Password) < 8 {
		http.Error(w, "La contraseña debe tener al menos 8 caracteres", http.StatusBadRequest)
		return
	}

	// Hash de la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error al generar el hash de la contraseña: "+err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Crear el usuario en la base de datos
	if err := db.CreateUser(&user); err != nil {
		http.Error(w, "Error al crear el usuario: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Devolver el usuario creado como respuesta
	w.WriteHeader(http.StatusCreated)
}

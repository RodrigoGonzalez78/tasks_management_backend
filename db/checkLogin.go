package db

import (
	"github.com/RodrigoGonzalez78/tasks_management_backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CheckLogin(email, password string) (models.User, bool, error) {
	// Buscar el usuario por su correo electrónico
	var user models.User
	result := dbConnection.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// El usuario no fue encontrado en la base de datos
			return user, false, nil
		}
		// Ocurrió un error al buscar el usuario
		return user, false, result.Error
	}

	// Verificar si la contraseña coincide
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// La contraseña no coincide
		return user, false, nil
	}

	// La contraseña coincide
	return user, true, nil
}

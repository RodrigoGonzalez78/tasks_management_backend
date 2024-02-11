package db

import (
	"github.com/RodrigoGonzalez78/tasks_management_backend/models"
	"gorm.io/gorm"
)

// Verifica si existe un usuario con ese email
func CheckExistUser(email string) (bool, error) {
	var user models.User
	result := DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

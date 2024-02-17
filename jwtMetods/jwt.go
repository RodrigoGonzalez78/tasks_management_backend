package jwtMetods

import (
	"errors"
	"strings"
	"time"

	"github.com/RodrigoGonzalez78/tasks_management_backend/db"
	"github.com/RodrigoGonzalez78/tasks_management_backend/models"
	"github.com/golang-jwt/jwt"
)

var SecretKey = []byte("NoteApp")

func CreateToken(t models.User) (string, error) {

	payload := jwt.MapClaims{
		"email":    t.Email,
		"name":     t.FirstName,
		"lastName": t.LastName,
		"id":       t.ID,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStrin, err := token.SignedString(SecretKey)

	if err != nil {
		return tokenStrin, err
	}

	return tokenStrin, nil
}

// valores para todos los endpoints
var Email string
var IDUser uint

// Proceso para extraer los datos del token
func ProcessToken(tk string) (*models.Claim, bool, uint, error) {

	claim := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claim, false, 0, errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claim, func(t *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err == nil {

		found, _ := db.CheckExistUser(claim.Email)

		if found {
			Email = claim.Email
			IDUser = claim.ID
		}

		return claim, found, IDUser, nil
	}

	if tkn.Valid {
		return claim, false, 0, errors.New("token invalido")
	}

	return claim, false, 0, err
}

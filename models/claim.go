package models

import "github.com/golang-jwt/jwt"

//Estructura para procesar el jwt
type Claim struct {
	Email string `json:"email"`
	ID    uint   `json:"id"`
	jwt.StandardClaims
}

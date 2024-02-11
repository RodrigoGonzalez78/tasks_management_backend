package middleware

import (
	"net/http"

	"github.com/RodrigoGonzalez78/tasks_management_backend/jwtMetods"
)

func CheckJwt(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := jwtMetods.ProcessToken(r.Header.Get("Authorization"))

		if err != nil {
			http.Error(w, "Erro en el token!"+err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}
}

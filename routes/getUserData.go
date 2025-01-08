package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/tasks_management_backend/db"
	"github.com/RodrigoGonzalez78/tasks_management_backend/jwtMetods"
	"github.com/RodrigoGonzalez78/tasks_management_backend/models"
)

func GetUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User = db.GetUserById(jwtMetods.IDUser)
	user.Password = ""

	json.NewEncoder(w).Encode(&user)
}

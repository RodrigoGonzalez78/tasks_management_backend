package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/tasks_management_backend/db"
	"github.com/RodrigoGonzalez78/tasks_management_backend/jwtMetods"
)

func GetAllTaksByUser(w http.ResponseWriter, r *http.Request) {
	tasks := db.GetAllTasksByUserId(jwtMetods.IDUser)
	json.NewEncoder(w).Encode(&tasks)
}

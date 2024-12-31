package main

import (
	"net/http"

	"github.com/RodrigoGonzalez78/tasks_management_backend/db"
	"github.com/RodrigoGonzalez78/tasks_management_backend/middleware"
	"github.com/RodrigoGonzalez78/tasks_management_backend/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	db.Migration()

	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)

	router.HandleFunc("/login", routes.Login).Methods("POST")
	router.HandleFunc("/signup", routes.SignUp).Methods("POST")
	router.HandleFunc("/user-data", middleware.CheckJwt(routes.GetUserData)).Methods("GET")
	router.HandleFunc("/delete-account", middleware.CheckJwt(routes.DeleteUserAccout)).Methods("DELETE")
	router.HandleFunc("/update-user", middleware.CheckJwt(routes.UpdateUser)).Methods("PUT")
	router.HandleFunc("/update-password", middleware.CheckJwt(routes.UpdateUserPassword)).Methods("PUT")

	router.HandleFunc("/tasks", middleware.CheckJwt(routes.GetAllTaksByUser)).Methods("GET")
	router.HandleFunc("/tasks/{id}", middleware.CheckJwt(routes.GetTask)).Methods("GET")
	router.HandleFunc("/tasks", middleware.CheckJwt(routes.CreateTask)).Methods("POST")
	router.HandleFunc("/tasks/{id}", middleware.CheckJwt(routes.DeleteTask)).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", middleware.CheckJwt(routes.UpdateTask)).Methods("PUT")

	http.ListenAndServe(":3000", router)
}

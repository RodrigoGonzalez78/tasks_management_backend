package main

import (
	"net/http"

	"github.com/RodrigoGonzalez78/tasks_management_backend/db"
	"github.com/RodrigoGonzalez78/tasks_management_backend/models"
	"github.com/RodrigoGonzalez78/tasks_management_backend/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	//Se crean las tablas users y task
	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Task{})

	router := mux.NewRouter()

	//Ruta home
	router.HandleFunc("/", routes.HomeHandler)

	router.HandleFunc("/login", routes.Login).Methods("POST")
	router.HandleFunc("/signup", routes.SignUp).Methods("POST")

	//Rutas del crud de tareas
	router.HandleFunc("/tasks", routes.GetTaksHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	router.HandleFunc("/tasks", routes.PostTaskHandler).Methods("POST")
	router.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")
	//Iniciamos el servidor
	http.ListenAndServe(":3000", router)
}

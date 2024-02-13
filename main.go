package main

import (
	"net/http"

	"github.com/RodrigoGonzalez78/tasks_management_backend/db"
	"github.com/RodrigoGonzalez78/tasks_management_backend/routes"
	"github.com/gorilla/mux"
)

func main() {

	//Iniciamos la base de datos
	db.DBConnection()

	//Migramos los modelos a la base de datos
	db.Migration()

	router := mux.NewRouter()

	//Ruta home
	router.HandleFunc("/", routes.HomeHandler)

	//Autenticacion
	router.HandleFunc("/login", routes.Login).Methods("POST")
	router.HandleFunc("/signup", routes.SignUp).Methods("POST")

	//Rutas para administrar usuarios
	router.HandleFunc("/users", routes.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserById).Methods("GET")
	router.HandleFunc("/users/{id}", routes.DeleteTaskById).Methods("DELETE")

	//Rutas para administrar tareas
	router.HandleFunc("/tasks", routes.GetAllTaks).Methods("GET")
	router.HandleFunc("/tasks/{id}", routes.GetTaskById).Methods("GET")
	router.HandleFunc("/tasks", routes.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", routes.DeleteTaskById).Methods("DELETE")

	//Iniciamos el servidor
	http.ListenAndServe(":3000", router)
}

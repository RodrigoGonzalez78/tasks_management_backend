package main

import (
	"net/http"

	"github.com/RodrigoGonzalez78/tasks_management_backend/db"
	"github.com/RodrigoGonzalez78/tasks_management_backend/middleware"
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

	//Manejo de la cuenta
	router.HandleFunc("/login", routes.Login).Methods("POST")
	router.HandleFunc("/signup", routes.SignUp).Methods("POST")
	router.HandleFunc("/delete-accout", middleware.CheckJwt(routes.DeleteUserAccout)).Methods("DELETE")
	router.HandleFunc("/update-user", middleware.CheckJwt(routes.UpdateUser)).Methods("PUT")
	router.HandleFunc("/update-password", middleware.CheckJwt(routes.UpdateUserPassword)).Methods("PUT")

	//Rutas para administrar tareas
	router.HandleFunc("/tasks", middleware.CheckJwt(routes.GetAllTaksByUser)).Methods("GET")
	router.HandleFunc("/tasks/{id}", middleware.CheckJwt(routes.GetTask)).Methods("GET")
	router.HandleFunc("/tasks", middleware.CheckJwt(routes.CreateTask)).Methods("POST")
	router.HandleFunc("/tasks/{id}", middleware.CheckJwt(routes.DeleteTask)).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", middleware.CheckJwt(routes.UpdateTask)).Methods("PUT")

	//Iniciamos el servidor
	http.ListenAndServe(":3000", router)
}

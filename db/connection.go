package db

import (
	"log"

	"github.com/RodrigoGonzalez78/tasks_management_backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "user=postgres password=12345678 dbname=taskM host=localhost port=5432 sslmode=disable"

// Conexion de a la base de datos
var dbConnection *gorm.DB

func DBConnection() {

	var dbError error
	dbConnection, dbError = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if dbError != nil {
		log.Fatal(dbError)
	} else {
		log.Println("Base de datos conectada!!")
	}

}

// Se crean las tablas users y task
func Migration() {
	dbConnection.AutoMigrate(models.User{})
	dbConnection.AutoMigrate(models.Task{})
}

package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "user=postgres password=12345678 dbname=taskM host=172.17.0.2 port=5432 sslmode=disable"
var DB *gorm.DB

func DBConnection() {

	var dbError error
	DB, dbError = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if dbError != nil {
		log.Fatal(dbError)
	} else {
		log.Println("Base de datos conectada!!")
	}

}

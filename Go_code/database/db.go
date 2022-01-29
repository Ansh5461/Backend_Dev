package database

import (
	"hello/Go_code/models"
	"log"
)

func Setup() {
	host := "localhost"
	port := "5432"
	dbName := "book"
	userName := "postgres"
	password := "postgres"
	db, err := gorm.open("postgres ", "host= "+host+" port= "+port+" dbName= "+dbName+" sslmode= disable userName= "+userName+" password= "+password)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(models.Book{})
	DB = db
}

package database

import (
	"log"

	"github.com/Ansh5461/Backend_Dev/tree/main/Go_code/models"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Setup() {
	host := "localhost"
	port := "5432"
	dbName := "book"
	userName := "postgres"
	password := "postgres"
	db, err := gorm.Open("postgres ", "host= "+host+" port= "+port+" dbName= "+dbName+" sslmode= disable userName= "+userName+" password= "+password)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(models.Book{})
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}

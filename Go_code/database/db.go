package database

import (
	"log"

	"github.com/Ansh5461/Backend_Dev/tree/main/Go_code/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Setup() {
	host := "localhost"
	port := "5432"
	dbName := "book"
	username := "postgres"
	passowrd := "postgres"
	args := "host=" + host + " port=" + port + " user=" + username + " dbname=" + dbName + " sslmode=disable password=" + passowrd
	db, err := gorm.Open("postgres", args)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(models.Book{})
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}

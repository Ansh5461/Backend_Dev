package databasee

import (
	"hello/Go_code/models"

	"github.com/jinzhu/gorm"
)

func GetDB() *gorm.DB {
	return DB
}

func GetBooks(db *gorm.DB) ([]models.Book, error) {
	books := []models.Book{}
	query := db.Select("books.*")
	err := query.Find(&books).Error

	if err != nil {
		return nil, err
	}
	return books, nil
}

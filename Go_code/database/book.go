package database

import (
	"github.com/Ansh5461/Backend_Dev/tree/main/Go_code/models"
	"github.com/jinzhu/gorm"
)

func GetBooks(db *gorm.DB) ([]models.Book, error) {
	books := []models.Book{}
	query := db.Select("books.*")
	err := query.Find(&books).Error

	if err != nil {
		return nil, err
	}
	return books, nil
}

func GetBookByID(db *gorm.DB, id string) (*models.Book, error) {
	book := models.Book{}
	err := db.Select("books.*").Group("books.id").Where("books.id=?", id).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}
func DeleteBookByID(db *gorm.DB, id string) error {
	return nil
}
func UpdateBookByID(db *gorm.DB, book *models.Book) error {
	return nil
}
func SaveBook(db *gorm.DB, book *models.Book) error {
	err := db.Save(book).Error

	if err != nil {
		return err
	}
	return nil
}

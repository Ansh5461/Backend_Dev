package handler

import (
	"log"
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"github.com/Ansh5461/Backend_Dev/tree/main/Go_code/database"
)

type Handler struct {
	DB *gorm.DB
}

type ( h *Handler) GetBooks (in *gin.Context) {
	books, err := database.GetBooks(h.DB)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, books)
}
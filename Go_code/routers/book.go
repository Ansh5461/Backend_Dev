package routers

import (
	"github.com/Ansh5461/Backend_Dev/tree/main/Go_code/database"
	"github.com/Ansh5461/Backend_Dev/tree/main/Go_code/handler"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	api := handler.Handler{
		DB: database.GetDB(),
	}
	router.GET("/books", api.GetBooks) //set the function for http://localhost:8080/books : Get request
	//while calling handler function, gin will pass *gin.Context in the handler function
	router.POST("/books", api.SaveBook)
	router.GET("/books/:id", api.GetBookByID)
	router.DELETE("/books/:id", api.DeleteBookByID)
	router.PUT("/books", api.UpdateBookByID)
	return router
}

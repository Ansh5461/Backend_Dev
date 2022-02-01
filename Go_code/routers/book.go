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
	router.GET("/books", api.GetBooks)
	router.POST("/books", api.SaveBook)
	return router
}

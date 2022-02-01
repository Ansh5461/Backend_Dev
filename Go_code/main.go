package main

import (
	"log"

	"github.com/Ansh5461/Backend_Dev/tree/main/Go_code/database"
	"github.com/Ansh5461/Backend_Dev/tree/main/Go_code/routers"
)

func main() {
	database.Setup()
	engine := routers.Router()
	err := engine.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
}

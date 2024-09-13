package main

import (
	"log"

	"github.com/gin-gonic/gin"
	database "github.com/veluvignesh027/worklog/Database"
	handlers "github.com/veluvignesh027/worklog/Internal/Handlers"
)

func init() {
	database.InitDataBase() //Intialize the database connection
}
func main() {
	router := gin.Default()
	err := handlers.InitHandelers(router)
	if err != nil {
		log.Fatal(err)
	}
	router.Run(":8080")
}

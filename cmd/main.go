package main

import (
	"log"
	"net/http"
	"redbook/cmd/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.StartDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	log.Printf("Successfully connected to database")

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	router.Run()
}

package main

import (
	"cmd/main/pkg/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	user.CreateUser()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	router.Run()
}

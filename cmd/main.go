package main

import (
	"redbook/pkg/controllers"
	"redbook/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.Cors())
	v1 := router.Group("/v1")
	{
		v1.POST("/signup", controllers.SignUp)
	}

	router.Run()
}

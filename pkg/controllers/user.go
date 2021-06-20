package controllers

import (
	"redbook/pkg/api"
	"redbook/pkg/models"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var user models.User
	c.Bind(&user)

	if user.GetUsername() != "" && user.GetPassword() != "" {
		_, err := api.CreateUser(&user)
		if err == nil {
			c.JSON(201, gin.H{"success": "Successfully made user"})
		} else {
			c.JSON(400, gin.H{"error": "Did not create user"})
		}
	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}
}

package middleware

import "github.com/gin-gonic/gin"

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Controll-Allow-Origin", "*")
		c.Next()
	}
}

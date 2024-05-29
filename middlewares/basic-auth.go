package middlewares

import "github.com/gin-gonic/gin"

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"juan": "1234",
	})
}

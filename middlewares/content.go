package middlewares

import "github.com/gin-gonic/gin"

func ContentJSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
	}
}

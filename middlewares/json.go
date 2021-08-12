package middlewares

import (
	"context"
	"fmt"
	"myapp/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type JsonCtx string

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerStr := "Bearer "
		auth := c.GetHeader("Authorization")

		if auth == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		auth = auth[len(bearerStr):]

		token, err := service.JwtValidateToken(auth)
		if err != nil || !token.Valid {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		tokenObj := token.Claims.(*service.JwtCustomClaim)
		fmt.Printf("%+v\n", time.Now().Unix()-tokenObj.ExpiresAt)

		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), JsonCtx("token_data"), &tokenObj))
	}
}

func TokenData(c *gin.Context) *service.JwtCustomClaim {
	raw, _ := c.Request.Context().Value(JsonCtx("token_data")).(*service.JwtCustomClaim)
	return raw
}

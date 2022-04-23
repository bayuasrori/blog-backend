package middleware

import (
	"goblog/authentication"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("accessToken")
		t, err := authentication.DecodeAcecessToken(token)
		if err != nil {

			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		claims, ok := t.Claims.(jwt.MapClaims)
		if ok && t.Valid {
			log.Println(claims["user_id"])
		}

		c.Next()
	}
}

package handler

import (
	"encoding/json"
	"goblog/authentication"
	"goblog/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshPayload struct {
	RefreshToken string `json:"refreshToken"`
}

func Login(c *gin.Context) {
	decoder := json.NewDecoder(c.Request.Body)
	var auth Auth

	if err := decoder.Decode(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.GetUserByEmail(auth.Email)
	if user.Password == auth.Password {
		accsessToken, _ := authentication.CreateToken(uint64(user.ID))
		refreshToken, _ := authentication.RefreshToken(uint64(user.ID))
		jwtToken := authentication.JWTToken{AccessToken: accsessToken, RefreshToken: refreshToken}
		models.CreateAuth(user, refreshToken)

		c.JSON(http.StatusOK, jwtToken)
		return
	}
}

func Refresh(c *gin.Context) {
	decoder := json.NewDecoder(c.Request.Body)
	var refresh RefreshPayload

	if err := decoder.Decode(&refresh); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, er := authentication.DecodeRefreshToken(refresh.RefreshToken)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId, ok := claims["user_id"].(uint64)
		if !ok {
			accsessToken, _ := authentication.CreateToken(userId)
			refreshToken, _ := authentication.RefreshToken(userId)

			jwtToken := authentication.JWTToken{AccessToken: accsessToken, RefreshToken: refreshToken}

			c.JSON(http.StatusOK, jwtToken)
		}
	}
}

func Register(c *gin.Context) {
	decoder := json.NewDecoder(c.Request.Body)

	var user_post models.User

	if err := decoder.Decode(&user_post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.CreateUser(user_post)

	c.JSON(http.StatusOK, gin.H{"status": "created"})
}

func Logout(c *gin.Context) {}

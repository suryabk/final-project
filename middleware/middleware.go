package middleware

import (
	"final-project/structs"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func MiddlewareAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.Header("WWW-Authenticate", `Bearer realm="Restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Mengambil token JWT dari header Authorization
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// Memverifikasi token JWT
		token, err := jwt.ParseWithClaims(tokenString, &structs.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return os.Getenv("JWTKEY"), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func GenerateToken(userID int, email string) string {
	expirationTime := time.Now().Add(24 * time.Hour) // Token berlaku selama 24 jam
	claims := structs.Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(os.Getenv("JWTKEY"))
	return tokenString
}

package structs

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	UserID    int       `json:"user_id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginReturn struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
}

type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

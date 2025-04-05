package dto

import "github.com/golang-jwt/jwt/v5"

type Credentials struct {
	Username string `json:"username" example:"username"`
	Password string `json:"password" example:"password"`
}

type Register struct {
	Username string `json:"username" example:"username"`
	Password string `json:"password" example:"password"`
	Email    string `json:"email" example:"email@gmail.com"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

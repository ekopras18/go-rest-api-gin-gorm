package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go-rest-api-gin-gorm/config"
	"go-rest-api-gin-gorm/internal/repository/auth"
	"go-rest-api-gin-gorm/pkg/utils"
	"time"
)

type Service struct {
	repo auth.Repository
}

func (s *Service) Login(creds utils.Credentials) (string, error) {
	user, err := s.repo.GetUserByUsername(creds.Username)
	if err != nil || user.ID == 0 {
		return "", errors.New("invalid credentials")
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &utils.Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JwtKey)
	if err != nil {
		return "", errors.New("could not generate token")
	}
	return tokenString, nil
}

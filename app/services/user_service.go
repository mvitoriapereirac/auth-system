package services

import (
	"auth-system/app/config"
	"auth-system/app/core/constants"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService interface {
	VerifyPassword(hashedPassword, password string) bool
	GenerateToken(userID uint) (string, error)
	SaveToken(ctx context.Context, userID uint, token string) error
}

type authService struct{}

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (s *authService) GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		constants.UserSessionKey: userID,
		"exp":                    time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret")) //Em uma aplicação real, deve ser usado um arquivo .ENV com variáveis de ambiente
}

func (s *authService) SaveToken(ctx context.Context, userID uint, token string) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return config.RDB.Set(ctx, getRedisKey(userID), token, time.Hour*24).Err()
}

func getRedisKey(userID uint) string {
	return fmt.Sprintf("%s%d", constants.RedisKey, userID)
}

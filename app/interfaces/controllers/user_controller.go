package controllers

import (
	"auth-system/app/config"
	"auth-system/app/core/constants"
	"auth-system/app/domain"
	"auth-system/app/usecases"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

type UserController struct {
	usecase usecases.UserUsecase
}

func NewUserController(usecase usecases.UserUsecase) *UserController {
	return &UserController{usecase: usecase}
}

func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	newUser, err := uc.usecase.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"user": newUser})
}

func (uc *UserController) Login(c *fiber.Ctx) error {
	var loginInput struct {
		CPF   string `json:"cpf"`
		Senha string `json:"senha"`
	}

	if err := c.BodyParser(&loginInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Dados inválidos"})
	}

	token, err := uc.usecase.Login(c.Context(), loginInput.CPF, loginInput.Senha)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}

func (uc *UserController) Logout(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "token ausente",
		})
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inválido")
		}
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "token inválido",
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "token inválido",
		})
	}

	userIDFloat, ok := claims[constants.UserSessionKey].(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "ID do usuário inválido no token",
		})
	}

	userID := uint(userIDFloat)

	err = config.RDB.Del(c.Context(), fmt.Sprintf("%s%d", constants.RedisKey, userID)).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "falha ao remover token",
		})
	}

	return c.JSON(fiber.Map{"message": "logout realizado com sucesso"})
}

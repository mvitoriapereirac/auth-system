package middleware

import (
    "auth-system/app/config"
    "context"
    "fmt"
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v4"
    "strings"
)

func AuthRequired() fiber.Handler {
    return func(c *fiber.Ctx) error {
        authHeader := c.Get("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token ausente ou mal formatado"})
        }

        tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
        token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
            return []byte("secret"), nil
        })

        if err != nil || !token.Valid {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token inválido"})
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token inválido"})
        }

        userID := fmt.Sprint(claims["user_id"])
        redisKey := "user_token:" + userID

        val, err := config.RDB.Get(context.Background(), redisKey).Result()
        if err != nil || val != tokenStr {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token expirado ou não autorizado"})
        }

        // Injeta ID no contexto
        c.Locals("user_id", userID)
        return c.Next()
    }
}

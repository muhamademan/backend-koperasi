package middleware

import (
	"backend-koperasi/utils"

	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("api-key")

	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "unauthenticated",
		})
	}

	// _, err := utils.VerifyToken(token)
	claims, err := utils.DecodeToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "invalid api-key or token",
		})
	}

	role := claims["role"].(string)

	if role != "admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   true,
			"message": "forbidden access",
		})
	}

	ctx.Locals("userInfo", claims)
	ctx.Locals("role", claims["role"])

	// if token != "secret" {
	// 	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error":   true,
	// 		"message": "Invalid api-key",
	// 	})
	// }

	return ctx.Next()
}

func PermissionMiddleware(ctx *fiber.Ctx) error {
	return ctx.Next()
}

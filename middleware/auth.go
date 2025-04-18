package middleware

import (
	"boilerplate/auth"

	"github.com/gofiber/fiber/v3"
)

// Checks incoming requests to validate the JWT token in the request
// Returns the next Middleware or Handler in the queue if JWT token is valid
// Otherwise will return StatusUnauthorized
func AuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		tokenString, err := auth.GetJWTToken(c)
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		
		err = auth.VerifyJWTToken(tokenString.(string))
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		return c.Next()
	}
}

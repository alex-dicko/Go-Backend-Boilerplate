package routes 

import (
	"boilerplate/helpers"
	"github.com/gofiber/fiber/v3"
	"boilerplate/dto"
	"boilerplate/models"
)

func RegisterUser(c fiber.Ctx) error {
	payload := struct {
		Username string `json:"username"`
		Email string `json:"email"`
		Password string `json:"password"`
	}{}

	err := c.Bind().JSON(&payload)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "incorrect payload",
			"user": nil,
		})
	}
	

	var user = models.User{
		Email: payload.Email,
		Username: payload.Username,
	}
	user.SetPassword(payload.Password)

	// Add checks to see if user with email/username already exists

	err = helpers.CreateModel(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "failed to create user",
			"user": nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Successfully created user",
		"user": dto.UserDTO{
			Email: user.Email,
			Username: user.Username,
		},
	})
}

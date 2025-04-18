package routes

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"boilerplate/auth"
	"boilerplate/config"
	"boilerplate/database"
	"boilerplate/dto"
	"boilerplate/helpers"
	"boilerplate/models"
	"boilerplate/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"boilerplate/logging"
)

func LoginUser(c fiber.Ctx) error {
	logger := logging.InitLogger("LOGIN")
	logger.Log(logging.Info, "Login Test")


	payload := struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}{}

	if err := c.Bind().JSON(&payload); err != nil {
		fmt.Println("WHAT")
		return err
	}

	var user models.User
	if err := database.Client.Where("username = ? OR email = ?", payload.User, strings.ToLower(payload.User)).
		Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "user doesn't exist",
				"user":    nil,
			})
		}
		return err
	}

	if !utils.CheckPasswordHash(payload.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "invalid credentials",
			"user":    nil,
		})
	}

	claims := jwt.MapClaims{
		"name":    user.Username,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
		"user_id": user.ID,
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := []byte(config.Vars.JWTSecret)
	tokenStr, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": nil,
		"token":   tokenStr,
		"user": dto.UserDTO{
			Username: user.Username,
			Email:    user.Email,
		},
	})
}

func CreateUser(c fiber.Ctx) error {
	var user = models.User{
		Username: "alex3",
		Email:    "alxdickens3@gmail.com",
	}
	user.SetPassword("alex")

	// Print the user object before creation for debugging
	fmt.Printf("Creating user: %+v\n", user)

	fmt.Println("ID")
	fmt.Println(user.ID)
	err := helpers.CreateModel(&user)
	if err != nil {
		// Return a more detailed error response
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to create user",
			"error":   err.Error(),
		})
	}
	fmt.Println(user.ID)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "User created successfully",
		"user": dto.UserDTO{
			Username: user.Username,
			Email:    user.Email,
		},
	})
}

func GetToken(c fiber.Ctx) error {
	token, err := auth.GetJWTToken(c)

	if err != nil {
		return err
	}

	claims, err := auth.GetJWTClaims(token.(string))
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"claims": claims,
	})
}

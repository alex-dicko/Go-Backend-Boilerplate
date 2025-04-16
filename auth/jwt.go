package auth

import (
	"errors"
	"boilerplate/config"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

// Verify a given JWT token
// If this does not produce any errors, we can trust the JWT token and its claims
func VerifyJWTToken(tokenString string) error {
	jwtSecret := []byte(config.Vars.JWTSecret)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}

type UserClaims struct {
	Name   string
	UserID float64
	Exp    float64
}

// Get claims from a given JWT token
// This does to see if the JWT token is valid
// so should only be used after making sure it is valid using auth.VerifyJWTToken
func GetJWTClaims(tokenString string) (UserClaims, error) {
	userClaims := UserClaims{}

	jwtSecret := []byte(config.Vars.JWTSecret)


	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return userClaims, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		userClaims.Name, _ = claims["name"].(string)
		userClaims.UserID, _ = claims["user_id"].(float64)
		userClaims.Exp, _ = claims["exp"].(float64)
	} else {
		return userClaims, errors.New("bad")
	}

	return userClaims, nil

}

// Takes in a request and returns the JWT token from that request.
// If it returns an error, no token exists in that request.
func GetJWTToken(c fiber.Ctx) (interface{}, error) {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return nil, errors.New("no token")
	}
	tokenString = tokenString[len("Bearer "):]

	return tokenString, nil
}

package main

import (
	"fmt"
	"boilerplate/config"
  "boilerplate/database"
	"boilerplate/middleware"
	"boilerplate/routes"
	"boilerplate/logging"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	config.Initialize() // Initialise .env variable (can be accessed at config.Vars.<key>)

	main_logger := logging.InitLogger("MAIN") // Initalise logger for main function

	main_logger.Log(logging.Error, "Logging Test")

	database.Connect() // Connect to database

	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{config.Vars.FrontendURL},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))

	// UNRESTRICTED VIEWS
	app.Post("/login", routes.LoginUser)
	app.Post("/create", routes.CreateUser)
	app.Post("/register", routes.RegisterUser)


	app.Post("/token", routes.GetToken)

	app.Use(middleware.AuthMiddleware())

	fmt.Println("Listening on port:", config.Vars.Port)
	fmt.Println("Accepting requests from:", config.Vars.FrontendURL)

	app.Listen(fmt.Sprintf("127.0.0.1:%s", config.Vars.Port))

}

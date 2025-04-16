package database

import (
	"boilerplate/models"
  "boilerplate/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Client *gorm.DB

// Takes no parameters. Connects to the database specified in the .env
func Connect() {
	var err error

	Client, err = gorm.Open(postgres.Open(config.Vars.PostgresURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = Client.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
}

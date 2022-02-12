package server

import (
	"log"
	"oliversavio/api/errors"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func InitFiberApplication() *fiber.App {
	loadApplicationConfig()
	config := fiber.Config{
		ReadTimeout:  2 * time.Second,
		ErrorHandler: errors.CustomErrorHandling,
	}
	return fiber.New(config)
}

func loadApplicationConfig() {
	loadEnv()
	log.Println("Running: ", os.Getenv("CURRENT_ENV"))
}

func loadEnv() {
	env := os.Getenv("FOO_ENV")
	if "" == env {
		env = "development"
	}
	godotenv.Load(".env." + env + ".local")
	if "test" != env {
		godotenv.Load(".env.local")
	}
	godotenv.Load(".env." + env)
	godotenv.Load()
}

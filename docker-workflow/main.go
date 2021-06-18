package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	fmt.Println("Starting Fiber API")
	app := fiber.New(fiber.Config{
		ErrorHandler: APIErrorHandler,
	})

	app.Use(recover.New())

	app.Get("/", rootHandler)
	// Connecting to an ES instance
	app.Get("/es", esInfoHandler)
	// Connect to Postgres
	app.Get("/pg", pgConnectHandler)
	// Start Server
	app.Listen(":3000")
}

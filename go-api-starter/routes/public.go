package routes

import (
	"oliversavio/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterPublicRoutes(a *fiber.App) {
	route := a.Group("/api")

	v1 := route.Group("/v1")
	v1.Get("/", handlers.RootHandler)
	v1.Get("/error", handlers.MockError)
}

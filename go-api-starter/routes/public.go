package routes

import (
	"oliversavio/api/handlers"
	"oliversavio/api/service"

	"github.com/gofiber/fiber/v2"
)

func RegisterPublicRoutes(a *fiber.App, quotter service.Quotter) {
	route := a.Group("/api")

	v1 := route.Group("/v1")
	v1.Get("/", handlers.RootHandler)
	v1.Get("/error", handlers.MockError)
	v1.Get("/quote", handlers.QuoteHandler(quotter))
}

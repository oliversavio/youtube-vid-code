package routes

import (
	"oliversavio/api/handlers"
	"oliversavio/api/middlewares"
	"oliversavio/api/service"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterPublicRoutes(a *fiber.App, quotter service.Quotter) {
	route := a.Group("/api", middlewares.RecordRequestLatency)

	v1 := route.Group("/v1")
	v1.Get("/", handlers.RootHandler)
	v1.Get("/error", handlers.MockError)
	v1.Get("/quote", handlers.QuoteHandler(quotter))
}

func RegisterPrometheusRoute(a *fiber.App) {
	a.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
}

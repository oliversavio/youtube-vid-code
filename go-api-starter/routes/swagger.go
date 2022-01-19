package routes

import (
	"log"
	"os"
	"strconv"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func RegisterSwaggerRoute(a *fiber.App) {
	enableSwagger := os.Getenv("ENABLE_SWAGGER")
	if enabled, _ := strconv.ParseBool(enableSwagger); enabled {
		route := a.Group("/swagger")
		route.Get("*", swagger.HandlerDefault)
		log.Println("Swagger Started")
	}

}

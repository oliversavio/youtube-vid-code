package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func FiberMiddlewares(a *fiber.App) {

	a.Use(
		recover.New(),
		logger.New(logger.Config{
			Format: "${status} - ${method}${path} [${latency}]\n",
		}),
	)
}

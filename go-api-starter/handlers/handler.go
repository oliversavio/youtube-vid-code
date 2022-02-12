package handlers

import (
	"oliversavio/api/errors"
	"oliversavio/api/service"

	"github.com/gofiber/fiber/v2"
)

// @Description Returns OK 200 and Hello World as a message
// @Produce json
// @Success 200
// @Router /v1/ [get]
func RootHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Hello World",
	})
}

// @Description Returns a 500 Error with message
// @Produce json
// @Failure 500 Mocked Internal Server Error
// @Router /v1/error/ [get]
func MockError(c *fiber.Ctx) error {
	return &errors.AppError{
		Msg:  "Mocked Error from Handler",
		Code: 500,
	}
}

func QuoteHandler(quotter service.Quotter) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		quote, err := quotter.GetQuote()
		if err != nil {
			return err
		}
		return c.Status(200).SendString(quote)
	}
}

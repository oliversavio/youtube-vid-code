package errors

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type AppError struct {
	Msg  string
	Code int
	Err  error
}

func (ae *AppError) Error() string {
	return fmt.Sprintf("%s with error %s", ae.Msg, ae.Err)
}

func CustomErrorHandling(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case *AppError:
		log.Println(e)
		return c.Status(e.Code).JSON(fiber.Map{
			"error": e.Msg,
		})
	default:
		log.Println(e)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "We've run into and issue, please try again later.",
		})
	}
}

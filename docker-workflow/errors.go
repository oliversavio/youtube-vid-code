package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var APIErrorHandler = func(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Send custom error page
	err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
	if err != nil {
		// In case the SendFile fails
		return ctx.Status(500).SendString("Internal Server Error")
	}

	// Return from handler
	return nil
}

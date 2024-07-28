package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/orhanince/go-wikiquote/src/model"
	"github.com/orhanince/go-wikiquote/src/router"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error

			if errors.As(err, &e) {
				code = e.Code
			}

			// Send custom error page
			err = ctx.Status(code).JSON(model.ErrorResponse{
				Status:  false,
				Message: e.Message,
			})
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			// Return from handler
			return nil
		},
	},
	)
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.Status(fiber.StatusOK).JSON(model.Response{
			Status:  true,
			Message: "Success",
			Data: map[string]interface{}{
				"name": "Go Wikiquote",
			}})

		return err
	})

	router.SetupRoutes(app)

	app.Listen(":9000")
}

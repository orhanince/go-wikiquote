package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/orhanince/go-wikiquote/src/model"
	"github.com/orhanince/go-wikiquote/src/router"
)

func main() {
	app := fiber.New()
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

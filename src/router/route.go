package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/orhanince/go-wikiquote/src/handler"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())

	// routes
	api.Get("/", handler.Home)
}

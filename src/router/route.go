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
	user := app.Group("/api/user", logger.New())
	scrape := app.Group("/api/scrape", logger.New())
	// Routes
	api.Get("/", handler.Home)

	// User routes
	user.Get("/", handler.Check)
	user.Get("/test", handler.Test)

	// Scrape
	scrape.Get("/", handler.Scrape)

}

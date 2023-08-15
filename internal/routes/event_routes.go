package routes

import (
	"kafka_events/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupEventRoutes(app *fiber.App) {
	app.Post("/visit", handlers.CreateEventHandler)
}

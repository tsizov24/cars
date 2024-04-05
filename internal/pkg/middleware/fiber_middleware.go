package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func FiberMiddleware(app *fiber.App) {
	app.Use(logger.New())
}

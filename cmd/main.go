package main

import (
	"cars/internal/pkg/configs"
	"cars/internal/pkg/middleware"
	"cars/internal/pkg/routes"

	"cars/internal/pkg/utils"

	"github.com/gofiber/fiber/v3"
)

func main() {
	conf := configs.FiberConfig()

	app := fiber.New(conf)

	middleware.FiberMiddleware(app)

	routes.Routes(app)

	utils.StartServer(app)
}

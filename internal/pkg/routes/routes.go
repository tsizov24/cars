package routes

import (
	"cars/internal/app/controllers"

	"github.com/gofiber/fiber/v3"
)

func Routes(app *fiber.App) {
	route := app.Group("/api/v1")

	route.Get("/info/:reg", controllers.GetCar)
	route.Get("/cars", controllers.GetCars)
	route.Post("/car", controllers.CreateCar)
	route.Put("/car", controllers.UpdateCar)
	route.Delete("/car/:reg", controllers.DeleteCar)
}

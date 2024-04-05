package utils

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

func StartServer(app *fiber.App) {
	url := fmt.Sprintf(
		"%s:%s",
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
	)

	logrus.Fatal(app.Listen(url))
}

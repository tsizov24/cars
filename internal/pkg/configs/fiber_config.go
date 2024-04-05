package configs

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
)

func FiberConfig() fiber.Config {
	timeout, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(timeout),
	}
}

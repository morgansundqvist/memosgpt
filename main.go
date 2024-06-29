package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/morgansundqvist/memosgpt/config"
	"github.com/morgansundqvist/memosgpt/handlers"
)

func main() {
	config.InitConfig()

	app := fiber.New()

	println("Server started on port 3214")

	app.Post("/wh", handlers.HandleWebHook)

	app.Listen(":3214")
}

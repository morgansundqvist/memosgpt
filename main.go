package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/morgansundqvist/memosgpt/handlers"
)

func main() {
	app := fiber.New()

	println("Server started on port 3214")

	app.Post("/wh", handlers.HandleWebHook)

	app.Listen(":3214")
}

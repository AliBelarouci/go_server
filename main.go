package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Define a route handler
	app.Get("/", func(c *fiber.Ctx) error {
		host, err := os.Hostname()
		if err != nil {
			return c.SendString("Error....")
		}
		return c.SendString("Hello, World! /" + host + " version 1.1.2")
	})
	// Start the server
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

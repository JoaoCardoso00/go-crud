package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

const SERVER_PORT = ":8000"

func main() {
	app := fiber.New()
	router := app.Group("/api")

	router.Get("/healthchecker", func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Server up and running",
		})
	})

	log.Fatal(app.Listen(":8000"))
}

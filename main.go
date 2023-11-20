package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "./public/")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("index.html")
	})

	app.Get("/home", func(c *fiber.Ctx) error {
		return c.SendFile("index.html")
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendFile("index.html")
	})

	app.Listen(":3000")
}

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type PageData struct {
	Title   string
	Message string
}

func main() {
	engine := html.New("./public/", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		pageDAta := PageData{
			Title:   "BROOOO",
			Message: "TAKE A SHOT",
		}

		return c.Render("index", pageDAta)
	})

	app.Listen(":3000")
}

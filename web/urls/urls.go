package urls

import (
	"github.com/gofiber/fiber/v2"
	"go-hello/web/login"
)

func Register(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	user := app.Group("/user")
	user.Post("/login", login.Login)

	app.Get("/info/:phone")
	app.Get("/info/:phone/json")
}

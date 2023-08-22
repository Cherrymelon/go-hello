package urls

import (
	"github.com/gofiber/fiber/v2"
	"go-hello/web/items"
	"go-hello/web/login"
)

func Register(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	user := app.Group("/user")
	user.Post("/login", login.Login)

	app.Get("/info/:phone", phone.Info)
	app.Get("/info/:phone/json", phone.Info_json)
}

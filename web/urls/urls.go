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

	info := app.Group("/info")
	info.Get("/:phone", phone.Info)
	info.Get("/:phone/json", phone.Info_json)
	info.Get("/phone_order/:order_id", phone.GetPhoneOrder)
	info.Post("/phone_order", phone.CreatePhoneOrder)
	info.Put("/phone_order/:order_id", phone.UpdatePhoneOrder)
}

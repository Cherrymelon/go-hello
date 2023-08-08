package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type phone struct {
	price int    `json:"price,omitempty"`
	brand string `json:"brand,omitempty"`
	ram   int    `json:"ram,omitempty"`
}

func main() {
	var ss = "this is sever"
	_ = ss
	hash := make(map[string]phone)
	hash["samsung"] = phone{price: 1000, brand: "samsung", ram: 4}
	hash["apple"] = phone{price: 2000, brand: "apple", ram: 8}
	hash["xiaomi"] = phone{price: 500, brand: "xiaomi", ram: 2}

	app := fiber.New()

	app.Use(logger.New())

	// Match any route
	app.Use(func(c *fiber.Ctx) error {
		log.Info("🥇 First handler")
		return c.Next()
	})

	// Match all routes starting with /api
	app.Use("/info", func(c *fiber.Ctx) error {
		log.Info("🥈 Second handler")
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/info/:phone", func(ctx *fiber.Ctx) error {
		msg := fmt.Sprintf("info about %+v", hash[ctx.Params("phone")])
		return ctx.SendString(msg)
	})
	app.Get("/info/:phone/json", func(ctx *fiber.Ctx) error {
		log.Info("params is ", hash[ctx.Params("phone")])
		msg := fiber.Map{"state": "ok", "data": hash[ctx.Params("phone")]}
		return ctx.JSON(msg)
	})

	log.Fatal(app.Listen(":3000"))

	//fmt.Println(ss, hash[], app
}

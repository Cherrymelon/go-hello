package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-hello/setting"
)

func main() {
	config := setting.Load_config()
	hash := make(map[string]phone)
	hash["samsung"] = phone{price: 1000, brand: "samsung", ram: 4}
	hash["apple"] = phone{price: 2000, brand: "apple", ram: 8}
	hash["xiaomi"] = phone{price: 500, brand: "xiaomi", ram: 2}
	app := fiber.New(fiber.Config{
		Prefork: config.WebServer.Prefork,
	})

	app.Use(logger.New())

	log.Fatal(app.Listen(config.WebServer.Host + ":" + fmt.Sprintf("%d", config.WebServer.Port)))

}

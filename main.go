package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	_ "go-hello/docs"
	"go-hello/setting"
	"go-hello/web/urls"
)

// swagger handler
// @title Fiber Example API of mine
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {

	config := setting.Load_config()
	app := fiber.New(fiber.Config{
		Prefork: config.WebServer.Prefork,
	})

	app.Use(logger.New())
	app.Get("/*", swagger.HandlerDefault)
	urls.Register(app)
	log.Fatal(app.Listen(config.WebServer.Host + ":" + fmt.Sprintf("%d", config.WebServer.Port)))

}

package items

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type phone struct {
	price int    `json:"price,omitempty"`
	brand string `json:"brand,omitempty"`
	ram   int    `json:"ram,omitempty"`
}

func (c *fiber.Ctx) error {
	msg := fmt.Sprintf("info about %+v", hash[c.Params("phone")])
	return c.SendString(msg)
}
func (c *fiber.Ctx) error {
	log.Info("params is ", hash[c.Params("phone")])
	msg := fiber.Map{"state": "ok", "data": hash[c.Params("phone")]}
	return c.JSON(msg)
}

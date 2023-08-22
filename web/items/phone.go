package phone

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

var hash = map[string]phone{
	"samsung": {price: 1000, brand: "samsung", ram: 4},
	"apple":   {price: 2000, brand: "apple", ram: 8},
	"xiaomi":  {price: 500, brand: "xiaomi", ram: 2},
}

func Info(c *fiber.Ctx) error {
	msg := fmt.Sprintf("info about %+v", hash[c.Params("phone")])
	return c.SendString(msg)
}
func Info_json(c *fiber.Ctx) error {
	log.Info("params is ", hash[c.Params("phone")])
	msg := fiber.Map{"state": "ok", "data": hash[c.Params("phone")]}
	return c.JSON(msg)
}

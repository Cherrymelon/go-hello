package phone

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go-hello/setting"
	"go-hello/web/models"
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

// Info  godoc
// @Summary      Show a phone info
// @Description  get phone info by str
// @Tags         Info
// @Accept       json
// @Produce      json
// @Param        phone   path      string  true  "Phone name"
// @Success      200  {object}  phone.phone
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /info/{phone} [get]
func Info(c *fiber.Ctx) error {
	msg := fmt.Sprintf("info about %+v", hash[c.Params("phone")])
	return c.SendString(msg)
}
func Info_json(c *fiber.Ctx) error {
	log.Info("params is ", hash[c.Params("phone")])
	msg := fiber.Map{"state": "ok", "data": hash[c.Params("phone")]}
	return c.JSON(msg)
}

func GetPhoneOrder(c *fiber.Ctx) error {
	var instance []models.Order
	id := c.Params("order_id")
	result := setting.Db.Find(&instance, id)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"state": "not found"})
	}
	return c.Status(200).JSON(instance)
}

func CreatePhoneOrder(c *fiber.Ctx) error {
	var instance []models.Order
	order := new(models.Order)
	if err := c.BodyParser(order); err != nil {
		log.Info("parser order failed")
		return c.Status(503).SendString(err.Error())
	}
	setting.Db.Create(&instance)
	return c.Status(201).JSON(instance)
}

func UpdatePhoneOrder(c *fiber.Ctx) error {
	instance := new(models.Order)
	id := c.Params("order_id")
	setting.Db.Where("id=?", id).Updates(&instance)
	return c.Status(200).JSON(instance)
}

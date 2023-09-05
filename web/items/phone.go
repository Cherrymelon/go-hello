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

// Info_json  godoc
// @Summary      Show a phone info
// @Description  get phone info by json
// @Tags         Info
// @Accept       json
// @Produce      json
// @Param        phone   path      string  true  "Phone name"
// @Success      200  {object}  phone.phone
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /info/{phone}/json [get]
func Info_json(c *fiber.Ctx) error {
	log.Info("params is ", hash[c.Params("phone")])
	msg := fiber.Map{"state": "ok", "data": hash[c.Params("phone")]}
	return c.JSON(msg)
}

// GetPhoneOrder  godoc
// @Summary      Show a phone order
// @Description  get phone order by id
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Order id"
// @Success      200  {object}  any
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /info/phone_order/{id} [get]
func GetPhoneOrder(c *fiber.Ctx) error {
	var instance []models.Order
	id := c.Params("id")
	result := setting.Db.Find(&instance, id)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"state": "not found"})
	}
	return c.Status(200).JSON(instance)
}

// CreatePhoneOrder  godoc
// @Summary      Create a phone order
// @Description  create phone order by json
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param        order   body      object  true  "Order"
// @Success      201  {object} any
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /info/phone_order [post]
func CreatePhoneOrder(c *fiber.Ctx) error {
	instance := new(models.Order)
	if err := c.BodyParser(instance); err != nil {
		log.Info("parser order failed")
		return c.Status(503).SendString(err.Error())
	}
	result := setting.Db.Create(&instance)
	log.Infof("create phone order %+v", result)
	return c.Status(201).JSON(result.RowsAffected)
}

// UpdatePhoneOrder  godoc
// @Summary      Update a phone order
// @Description  update phone order by json
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Order id"
// @Param        order   body      object  true  "Order"
// @Success      200  {object} any
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /info/phone_order/{id} [put]
func UpdatePhoneOrder(c *fiber.Ctx) error {
	instance := new(models.Order)
	id := c.Params("id")
	if err := c.BodyParser(instance); err != nil {
		log.Info("parser order failed")
		return c.Status(503).SendString(err.Error())
	}
	result := setting.Db.Where("id=?", id).Updates(&instance)
	log.Infof("result is %+v", result)
	return c.Status(200).JSON(instance)
}

package login

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type login_form struct {
	user   string `json:"user,omitempty"`
	passwd string `json:"passwd,omitempty"`
}

func Login(c *fiber.Ctx) error {
	form := new(login_form)
	if err := c.BodyParser(form); err != nil {
		log.Info("login failed")
		return err
	}
	if form.user != "admin" || form.passwd != "admin" {
		log.Info("login failed")
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	log.Info("login success")
	return c.SendStatus(fiber.StatusOK)

}

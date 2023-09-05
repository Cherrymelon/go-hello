package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go-hello/setting"
	"go-hello/web/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func login(c *fiber.Ctx) error {
	type LoginRequest struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}
	var loginForm LoginRequest
	if err := c.BodyParser(loginForm); err != nil {
		log.Error(err.Error())
		return c.Status(403).SendString("parse error")
	}
	found := models.User{}
	query := models.User{Name: loginForm.User}
	result := setting.Db.Where(query).First(&found)
	if result.Error == gorm.ErrRecordNotFound {
		return c.Status(403).SendString("user not found error")
	}
	if !comparePasswords(found.Password, []byte(loginForm.Password)) {
		return c.Status(403).SendString("password error")
	}

	return c.SendString("login")
}

func hashAndSalt(pwd []byte) string {
	hash, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	return err == nil
}

func register(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(user)
	if err != nil {
		log.Error(err.Error())
	}
	user.Password = hashAndSalt([]byte(user.Password))
	result := setting.Db.Where("name=?", user.Name).First(&user)
	if result.Error != nil {
		return c.Status(403).SendString("register error,name exist")
	}
	result = setting.Db.Create(&user)
	if result.Error != nil {
		return c.Status(403).SendString("register error")
	}
	return c.SendString("register success!")
}

func logout(c *fiber.Ctx) error {
	return c.SendString("logout")
}

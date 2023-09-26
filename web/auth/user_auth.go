package auth

import (
	"github.com/badoux/checkmail"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	guuid "github.com/google/uuid"
	"go-hello/setting"
	"go-hello/web/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type Session models.Session

var db = setting.Db

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
	result := db.Where(query).First(&found)
	if result.Error == gorm.ErrRecordNotFound {
		return c.Status(403).SendString("user not found error")
	}
	if !comparePasswords(found.Password, []byte(loginForm.Password)) {
		return c.Status(403).SendString("password error")
	}
	session := Session{UserRefer: found.ID, Expires: SessionExpires(), Sessionid: guuid.New()}
	db.Create(&session)
	c.Cookie(&fiber.Cookie{
		Name:     "sessionid",
		Expires:  SessionExpires(),
		Value:    session.Sessionid.String(),
		HTTPOnly: true,
	})

	log.Info("login success,user {}", found)
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
	type CreateUserRequest struct {
		Password string `json:"password"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	json := new(CreateUserRequest)
	err := c.BodyParser(json)
	if err != nil {
		log.Error(err.Error())
		return c.Status(403).SendString("parse error")
	}
	json.Password = hashAndSalt([]byte(json.Password))
	err = checkmail.ValidateFormat(json.Email)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid Email Address",
		})
	}
	user := models.User{
		Email:    json.Email,
		Password: json.Password,
		Name:     json.Username,
	}
	result := db.Where("name=?", json.Username).First(&user)
	if result.Error != gorm.ErrRecordNotFound {
		return c.Status(403).SendString("register error,name exist")
	}
	result = db.Create(&user)
	if result.Error != nil {
		return c.Status(403).SendString("register error")
	}
	session := Session{
		UserRefer: user.ID,
		Sessionid: guuid.New(),
	}
	if err := db.Create(&session).Error; err != nil {
		return c.Status(403).SendString("create session error")
	}

	c.Cookie(&fiber.Cookie{
		Name:     "sessionid",
		Expires:  time.Now().Add(5 * 24 * time.Hour),
		Value:    session.Sessionid.String(),
		HTTPOnly: true,
	})
	log.Info("register success,user {}", user)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
		"data":    session,
	})
}

func logout(c *fiber.Ctx) error {

	json := new(Session)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	session := Session{}
	query := Session{Sessionid: json.Sessionid}
	err := db.First(&session, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Session not found",
		})
	}
	db.Delete(&session)
	c.ClearCookie("sessionid")
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
	})

}

func SessionExpires() time.Time {
	return time.Now().Add(5 * 24 * time.Hour)
}

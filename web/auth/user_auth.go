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

// Login godoc
// @Summary      Login
// @Description  login by json
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user   body      object  true  "User"
// @Success      200  {object}  any
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /user/login [post]
func Login(c *fiber.Ctx) error {
	var db = setting.Db
	type LoginRequest struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}
	loginForm := new(LoginRequest)
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

// Register godoc
// @Summary      Register a user
// @Description  register a user by json
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user   body      object  true  "User"
// @Success      200  {object}  any
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /user/register [post]
func Register(c *fiber.Ctx) error {
	var db = setting.Db
	type CreateUserRequest struct {
		Password string `json:"password"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	log.Info("db is", db)
	json := new(CreateUserRequest)
	err := c.BodyParser(json)
	if err != nil {
		log.Error(err.Error())
		return c.Status(403).SendString("parse error")
	}
	json.Password = hashAndSalt([]byte(json.Password))
	err = checkmail.ValidateFormat(json.Email)
	if err != nil {
		log.Error("Invalid Email Addres")
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid Email Address",
		})
	}
	user := models.User{
		Password: json.Password,
		Email:    json.Email,
		Name:     json.Username,
	}
	find := models.User{}
	//result := db.Where("name=?", json.Username).Find(&user)
	result := db.First(&find, "name =?", json.Username)
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
		"message": "success",
		"data":    session,
	})
}

// Logout godoc
// @Summary      Logout
// @Description  logout by json
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        sessionid   body      object  true  "Session"
// @Success      200  {object}  any
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /user/logout [post]
func Logout(c *fiber.Ctx) error {
	var db = setting.Db
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

func hashAndSalt(pwd []byte) string {
	hash, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	return err == nil
}

package routes

import (
	"encoding/base64"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *fiber.Ctx) error {

	var LoginForm struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var user User

	//get request body
	c.BodyParser(&LoginForm)

	query := db.Where("password = crypt($1, password) AND email = $2", LoginForm.Password, LoginForm.Email).First(&user).Scan(&user)

	if query.RowsAffected < 1 {
		return c.Status(403).SendString("login failed")
	}

	//generete JWT token
	var tokenString string

	key, err := base64.StdEncoding.DecodeString(os.Getenv("JWT_KEY"))

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.String(),
		"name":    user.Name,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err = t.SignedString(key)

	if err != nil {
		log.Println(err)
	}

	type Resp struct {
		Token     string `json:"token"`
		UserName  string `json:"user_name"`
		UserEmail string `json:"user_email"`
		UserID    string `json:"user_id"`
	}

	var resp Resp

	resp.Token = tokenString
	resp.UserName = user.Name
	resp.UserEmail = user.Email
	resp.UserID = user.ID.String()

	return c.JSON(resp)
}

package routes

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c fiber.Ctx) error {

	var db = DbConnect()

	var LoginForm struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var user User

	c.Bind().Body(&LoginForm)

	//query := db.Where("password = crypt($1, password) AND email = $2", LoginForm.Password, LoginForm.Email).First(&user).Scan(&user)
	query := db.Where("email = ?", LoginForm.Email).First(&user).Scan(&user)

	pwCheck := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(LoginForm.Password))

	if query.RowsAffected < 1 {
		return c.Status(403).SendString(query.Error.Error())
	}

	if pwCheck == bcrypt.ErrMismatchedHashAndPassword {
		return c.Status(403).SendString(pwCheck.Error())
	}

	//generete JWT token
	var tokenString string

	key := os.Getenv("JWT_KEY")

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.String(),
		"name":    user.Name,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := t.SignedString([]byte(key))

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

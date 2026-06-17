package routes

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func PostUser(c fiber.Ctx) error {
	var db = DbConnect()

	var user User

	err := c.Bind().Body(&user)

	if err != nil {
		log.Println(err)
	}

	pwhash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 2)

	user.ID = uuid.New()
	user.Password = string(pwhash)

	rp := db.Create(&user)

	if rp.Error != nil {
		return c.SendStatus(400)
	}

	return c.SendStatus(200)
}

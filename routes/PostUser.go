package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PostUser(c *fiber.Ctx) error {
	var db = DbConnect()

	var user User

	err := c.BodyParser(&user)

	if err != nil {
		log.Println(err)
	}

	user.ID = uuid.New()

	statusCode := 200

	rp := db.Create(&user)

	if rp.Error != nil {
		statusCode = 400
	}

	return c.SendStatus(statusCode)
}

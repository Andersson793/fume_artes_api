package routes

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func PostProducts(c fiber.Ctx) error {
	var db = DbConnect()

	var product Products

	err := c.Bind().Body(&product)

	if err != nil {
		log.Println(err)
	}

	product.ID = uuid.New()

	statusCode := 200

	rp := db.Create(&product)

	if rp.Error != nil {
		statusCode = 400
	}

	return c.SendStatus(statusCode)
}

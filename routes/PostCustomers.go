package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PostCustomers(c *fiber.Ctx) error {

	var customer Customer

	err := c.BodyParser(&customer)

	if err != nil {
		log.Println(err)
	}

	customer.ID = uuid.New()

	statusCode := 200

	rp := db.Create(&customer)

	if rp.Error != nil {
		statusCode = 400
	}

	return c.SendStatus(statusCode)
}

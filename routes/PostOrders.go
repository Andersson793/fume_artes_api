package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PostOrders(c *fiber.Ctx) error {
	var db = DbConnect()

	var order Order

	//parse request body
	err := c.BodyParser(&order)

	if err != nil {
		log.Println(err)
	}

	order.ID = uuid.New()

	statusCode := 200

	rp := db.Create(&order)

	if rp.Error != nil {
		statusCode = 400
	}

	return c.SendStatus(statusCode)
}

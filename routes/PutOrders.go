package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PutOrders(c *fiber.Ctx) error {
	var db = DbConnect()
	var order Order

	statusCode := 200

	c.BodyParser(&order)

	//resp := db.Save(&order)
	resp := db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order)

	if resp.Error != nil {
		statusCode = 400
	}

	return c.SendStatus(statusCode)
}

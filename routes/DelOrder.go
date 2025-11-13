package routes

import "github.com/gofiber/fiber/v2"

func DelOrder(c *fiber.Ctx) error {

	var order Order

	db.Where("id = ?", c.Params("id")).Delete(&order)

	return c.SendString("item ? deleted", c.Params("id"))
}

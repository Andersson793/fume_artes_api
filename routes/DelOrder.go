package routes

import "github.com/gofiber/fiber/v3"

func DelOrder(c fiber.Ctx) error {

	var db = DbConnect()

	var order Order

	db.Where("id = ?", c.Params("id")).Delete(&order)

	return c.SendString("item " + c.Params("id") + " deleted")
}

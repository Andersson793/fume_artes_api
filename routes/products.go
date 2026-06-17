package routes

import "github.com/gofiber/fiber/v3"

func ProductsList(c fiber.Ctx) error {
	var db = DbConnect()

	var product []Products

	db.Find(&product)

	return c.JSON(product)
}

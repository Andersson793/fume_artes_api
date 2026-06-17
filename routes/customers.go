package routes

import (
	"github.com/gofiber/fiber/v3"
)

func Customers(c fiber.Ctx) error {

	var customers []Customer

	var db = DbConnect()

	db.Find(&customers)

	return c.JSON(customers)
}

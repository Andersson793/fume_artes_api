package routes

import "github.com/gofiber/fiber/v2"

func Users(c *fiber.Ctx) error {

	var users []User

	db.Find(&users)

	return c.JSON(users)
}

package routes

import "github.com/gofiber/fiber/v3"

func Users(c fiber.Ctx) error {
	var db = DbConnect()

	var users []User

	db.Find(&users)

	return c.JSON(users)
}

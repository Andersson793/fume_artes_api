package routes

import (
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	var db = DbConnect()

	var user User

	//can use cookies

	db.Find(&user, "id = ?", c.Params("id"))

	return c.JSON(user)
}

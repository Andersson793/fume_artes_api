package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func GetOrder(c *fiber.Ctx) error {

	var db = DbConnect()

	//var orders []Order
	var order Order

	db.Preload("order_Items").Preload(clause.Associations).Where("id = ?", c.Params("id")).Find(&order).Scan(&order)

	return c.JSON(order)
}

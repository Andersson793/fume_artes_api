package routes

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func Orders(c fiber.Ctx) error {
	var db = DbConnect()
	var orders []Order

	type Result struct {
		ID          uuid.UUID `json:"id"`
		Description string    `json:"description"`
		Payment     string    `json:"payment"`
		Customer    string    `json:"customer"`
		CreatedAt   time.Time `json:"created_at"`
		Total       string    `json:"total_items"`
	}

	var result []Result

	db.Model(&orders).Preload("OrderItems").Select("orders.id, orders.description,orders.payment, orders.customer, orders.created_at, SUM(order_items.price) as total").Joins("inner join order_items on order_items.order_id = orders.id").Group("orders.id").Order("orders.created_at DESC").Limit(100).Scan(&result)

	return c.JSON(&result)
}

package routes

import "github.com/gofiber/fiber/v3"

func Historical(c fiber.Ctx) error {

	var db = DbConnect()

	var resp []HistoricalData

	db.Table("historical_data").Limit(7).Scan(&resp)

	return c.JSON(resp)
}

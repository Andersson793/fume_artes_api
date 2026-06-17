package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"

	"github.com/Andersson793/fume_artes_api/m/routes"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env.local")

	app := fiber.New()

	app.Use(cors.New())

	//v1
	api := app.Group("/api", ValidateLogin)

	api.Get("/customers", routes.Customers)
	api.Get("/users/:id", routes.GetUser)
	api.Get("/users", routes.Users)
	api.Get("/orders_full/:id", routes.GetOrder)
	api.Get("/orders", routes.Orders)
	api.Get("/hgbrasil", routes.HgBrasil)
	api.Get("/historical_data", routes.Historical)
	api.Get("/products", routes.ProductsList)

	app.Post("/login", routes.Login)
	api.Post("/users", routes.PostUser)
	api.Post("/customers", routes.PostCustomers)
	api.Post("/orders", routes.PostOrders)
	api.Post("/products", routes.PostProducts)

	api.Put("/orders", routes.PutOrders)

	api.Delete("/orders/:id", routes.DelOrder)

	app.Listen(":3000")
}

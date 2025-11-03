package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/Andersson793/fume_artes_api/m/routes"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env.local")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://fume-artes-system.vercel.app, http://127.0.0.2:4173, http://localhost:5173, http://127.0.0.2, http://127.0.0.2:5173",
	}))

	api := app.Group("/api", func(c *fiber.Ctx) error {

		authorization := c.GetReqHeaders()

		var tokenString string

		if len(authorization["Authorization"]) > 0 {
			tokenString = authorization["Authorization"][0]
		} else {
			return fiber.NewError(404, "The Autorization header is missing")
		}

		var _, err = routes.JwtValidator(tokenString)

		if err != nil {

			c.SendStatus(404)

			return c.SendString(err.Error())
		}

		return c.Next()

	})

	api.Get("/customers", routes.Customers)
	api.Get("/users/:id", routes.GetUser)
	api.Get("/users", routes.Users)
	api.Get("/orders_full/:id", routes.GetOrder)
	api.Get("/orders", routes.Orders)
	api.Get("/hgbrasil", routes.HgBrasil)
	api.Get("historical_data", routes.Historical)
	//app.Get("/jwt_validate", routes.JwtValidator)

	//login
	app.Post("/login", routes.Login)
	api.Post("/users", routes.PostUser)
	api.Post("/customers", routes.PostCustomers)
	api.Post("/orders", routes.PostOrders)

	api.Put("/Orders", routes.PutOrders)

	api.Delete("/orders/:id", routes.DelOrder)

	app.Listen(":3000")
}

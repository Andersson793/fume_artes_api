package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func main() {

	//load .env file
	err := godotenv.Load()

	if err != nil {
		log.Println("Failed to load env file")
	}

	app := fiber.New()

	//connect database
	db, err := gorm.Open(postgres.Open(os.Getenv("PG_STRING")), &gorm.Config{})

	if err != nil {
		println("Can't connect database")
	}

	app.Get("/customers", func(c *fiber.Ctx) error {

		var customers []Customer

		db.Find(&customers)

		return c.JSON(customers)
	})

	app.Get("/users", func(c *fiber.Ctx) error {

		var users []User

		db.Find(&users)

		return c.JSON(users)
	})

	app.Get("/orders", func(c *fiber.Ctx) error {
		var orders []Order

		db.Find(&orders)

		return c.JSON(orders)
	})

	app.Get("/pending_services", func(c *fiber.Ctx) error {
		var pending_service []PendingService

		db.Find(&pending_service)

		return c.JSON(pending_service)
	})

	app.Post("/users", func(c *fiber.Ctx) error {

		var user User

		err := c.BodyParser(&user)

		if err != nil {
			log.Println(err)
		}

		user.ID = uuid.New()

		statusCode := 200

		rp := db.Create(&user)

		if rp.Error != nil {
			statusCode = 400
		}

		return c.SendStatus(statusCode)
	})

	app.Post("/customers", func(c *fiber.Ctx) error {

		var customer Customer

		err := c.BodyParser(&customer)

		if err != nil {
			log.Println(err)
		}

		customer.ID = uuid.New()

		statusCode := 200

		rp := db.Create(&customer)

		if rp.Error != nil {
			statusCode = 400
		}

		return c.SendStatus(statusCode)
	})

	app.Post("/orders", func(c *fiber.Ctx) error {

		var orders Order

		//parse request body
		err := c.BodyParser(&orders)

		if err != nil {
			log.Println(err)
		}

		orders.ID = uuid.New()

		statusCode := 200

		rp := db.Create(&orders)

		if rp.Error != nil {
			statusCode = 400
		}

		return c.SendStatus(statusCode)
	})

	app.Post("/pending_services", func(c *fiber.Ctx) error {

		var pendingServices PendingService

		err := c.BodyParser(&pendingServices)

		if err != nil {
			log.Println(err)
		}

		pendingServices.ID = uuid.New()

		statusCode := 200

		rp := db.Create(&pendingServices)

		if rp.Error != nil {
			statusCode = 400
		}

		return c.SendStatus(statusCode)
	})

	app.Listen(":3000")
}

package main

import (
	"encoding/base64"
	"errors"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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

	api := app.Group("/api", func(c *fiber.Ctx) error {

		authorization := c.GetReqHeaders()

		var tokenString string = authorization["Authorization"][0]

		key, err := base64.StdEncoding.DecodeString(os.Getenv("JWT_KEY"))

		//validate token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {

			return key, nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

		switch {
		case token.Valid:
			return c.Next()
		case errors.Is(err, jwt.ErrTokenMalformed):
			return c.SendString("That's not even a token")
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			// Invalid signature
			return c.SendString("Invalid signature")
		case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
			// Token is either expired or not active yet
			return c.SendString("Timing is everything")
		default:
			return c.SendString("Couldn't handle this token")
		}

	})

	//connect database
	db, err := gorm.Open(postgres.Open(os.Getenv("PG_STRING")), &gorm.Config{})

	if err != nil {
		println("Can't connect database")
	}

	api.Get("/customers", func(c *fiber.Ctx) error {

		var customers []Customer

		db.Find(&customers)

		return c.JSON(customers)
	})

	api.Get("/users/:id", func(c *fiber.Ctx) error {

		var user User

		//can use cookies

		db.Find(&user, "id = ?", c.Params("id"))

		return c.JSON(user)
	})

	api.Get("/users", func(c *fiber.Ctx) error {

		var users []User

		db.Find(&users)

		return c.JSON(users)
	})

	api.Get("/orders", func(c *fiber.Ctx) error {
		var orders []Order

		type Result struct {
			ID          uuid.UUID `json:"id"`
			Description string    `json:"description"`
			Customer    string    `json:"customer"`
			CreatedAt   time.Time `json:"created_at"`
			Total       string    `json:"total_items"`
		}

		var result []Result

		db.Model(&orders).Select("orders.id, orders.description, orders.customer, orders.created_at, SUM(order_items.price) as total").Joins("inner join order_items on order_items.order_id = orders.id").Group("orders.id").Scan(&result)

		return c.JSON(result)
	})

	//see Query params
	api.Get("/pending_services", func(c *fiber.Ctx) error {
		var pending_service []PendingService

		type Result struct {
			ID          uuid.UUID `json:"id"`
			Name        string    `json:"user_name"`
			Description string    `json:"description"`
			CreatedAt   time.Time `json:"created_at"`
		}

		var result []Result

		db.Model(&pending_service).Select("pending_services.id, users.name, pending_services.description, pending_services.created_at").Joins("inner join users on pending_services.user_id = users.id").Scan(&result)

		return c.JSON(result)
	})

	//login
	app.Get("/login", func(c *fiber.Ctx) error {

		var LoginForm struct {
			Email    string
			Password string
		}

		var user User

		//get request body
		c.BodyParser(LoginForm)

		er := db.First(&user).Where("password = crypt($1, password) and email = $2", LoginForm.Password, LoginForm.Email).Scan(&user)

		if er.Error != nil {
			c.SendStatus(403)
			c.SendString("login failed")
		}

		//generete JWT token
		var tokenString string

		key, err := base64.StdEncoding.DecodeString(os.Getenv("JWT_KEY"))

		t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.ID.String(),
			"exp":     time.Now().Add(24 * time.Hour).Unix(),
		})

		tokenString, err = t.SignedString(key)

		//set cookie
		var cookie fiber.Cookie
		var cookie2 fiber.Cookie
		var cookie3 fiber.Cookie

		cookie.Name = "user_id"
		cookie.Value = user.ID.String()
		cookie.Secure = true
		cookie.HTTPOnly = true
		cookie.SameSite = "Strict"
		cookie.Expires = time.Now().Add(24 * time.Hour)

		cookie2.Name = "name"
		cookie2.Value = user.Name
		cookie2.Secure = true
		cookie2.HTTPOnly = false
		cookie2.SameSite = "Strict"
		cookie2.Expires = time.Now().Add(24 * time.Hour)

		cookie3.Name = "test"
		cookie3.Value = user.Name
		cookie3.Secure = false
		cookie3.HTTPOnly = false
		cookie3.SameSite = "None"
		cookie3.Expires = time.Now().Add(24 * time.Hour)
		cookie3.SessionOnly = true
		cookie3.Domain = "localhost"

		c.Cookie(&cookie)
		c.Cookie(&cookie2)
		c.Cookie(&cookie3)

		if err != nil {
			log.Println(err)
		}

		return c.SendString(tokenString)
	})

	//validate jwt token
	app.Get("/jwt_validate", func(c *fiber.Ctx) error {

		authorization := c.GetReqHeaders()

		var tokenString string = authorization["Authorization"][0]

		key, err := base64.StdEncoding.DecodeString(os.Getenv("JWT_KEY"))

		//validate token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {

			return key, nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

		switch {
		case token.Valid:
			return c.SendString("You look nice today")
		case errors.Is(err, jwt.ErrTokenMalformed):
			return c.SendString("That's not even a token")
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			// Invalid signature
			return c.SendString("Invalid signature")
		case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
			// Token is either expired or not active yet
			return c.SendString("Timing is everything")
		default:
			return c.SendString("Couldn't handle this token")
		}

	})

	api.Post("/users", func(c *fiber.Ctx) error {

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

	api.Post("/customers", func(c *fiber.Ctx) error {

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

	api.Post("/orders", func(c *fiber.Ctx) error {

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

	api.Post("/pending_services", func(c *fiber.Ctx) error {

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

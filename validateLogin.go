package main

import (
	"github.com/Andersson793/fume_artes_api/m/routes"
	"github.com/gofiber/fiber/v3"
)

func ValidateLogin(c fiber.Ctx) error {

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

}

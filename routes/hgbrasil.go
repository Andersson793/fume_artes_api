package routes

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func HgBrasil(c *fiber.Ctx) error {

	//fiber cache (change !)
	c.Response().Header.Add("Cache-Control", "max-age=3600, private")

	agent := fiber.Get("https://api.hgbrasil.com/finance?key=" + os.Getenv("HG_KEY"))

	statusCode, body, errs := agent.Bytes()

	if len(errs) > 0 {
		log.Println(errs)
	}

	var something fiber.Map
	json.Unmarshal(body, &something)

	return c.Status(statusCode).JSON(something)
}

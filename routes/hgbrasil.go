package routes

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v3"
)

func HgBrasil(c fiber.Ctx) error {

	//fiber cache (change !)
	c.Response().Header.Add("Cache-Control", "max-age=3600, private")

	//agent := fiber.Get("https://api.hgbrasil.com/finance?key=" + os.Getenv("HG_KEY"))
	agent, err := http.Get("https://api.hgbrasil.com/finance?key=" + os.Getenv("HG_KEY"))

	if err != nil {
		c.SendString("fail to fetch hgbrasil api")
	}

	body := agent.Body

	return c.JSON(body)
}

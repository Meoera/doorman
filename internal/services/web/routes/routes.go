package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meoera/doorman/internal/models"
	"github.com/meoera/doorman/internal/services/config"
	"github.com/meoera/doorman/internal/services/web"
)

func Add(cfg *config.Web, redisCfg *config.Redis, devMode bool) {

	web.Server.Post("/login", func(c *fiber.Ctx) (err error) {
		c.Accepts("application/json")
		c.AcceptsCharsets("utf-8")

		body := &models.AuthBody{}
		err = c.BodyParser(body)
		if err != nil {
			if devMode {
				return c.SendString(err.Error())
			}
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Credentials!",
		})
	})

	web.Server.Post("/refresh", func(c *fiber.Ctx) (err error) {
		c.Accepts("application/json")
		c.AcceptsCharsets("utf-8")

		return
	})

}

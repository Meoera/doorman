package routes

import (
	"fmt"
	"strings"
	"time"

	"github.com/meoera/doorman/pkg/token"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/meoera/doorman/internal/services/config"
	"github.com/meoera/doorman/internal/services/database"
	"github.com/meoera/doorman/internal/services/web"
	"github.com/meoera/doorman/pkg/hasher"
	"github.com/meoera/doorman/pkg/models"
)

func Add(cfg *config.Web, db database.Database, cacheDb database.CacheDatabase, devMode bool) {

	if !devMode {
		web.Server.Use(limiter.New(limiter.Config{
			Max: int(cfg.AllowedRequestsPerMinute),
			Expiration: 1 * time.Minute,
			KeyGenerator: func(c *fiber.Ctx) string {
				return c.IP()
			},
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
					"message": "You send too many requests",
				})
			},
		}))
	}

	web.Server.Post("/login", func(c *fiber.Ctx) (err error) {
		c.Accepts("application/json")
		c.AcceptsCharsets("utf-8")

		body := &models.AuthRequestBody{}
		err = c.BodyParser(body)
		if err != nil {
			if devMode {
				return c.SendString(err.Error())
			} else {
				fmt.Println(err)
			}
		}

		dbRecord, err := db.UserByName(body.Username)
		if dbRecord == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Bad Credentials!",
			})
		}
		if err != nil {
			if devMode {
				panic(err)
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "An error occured! Try again later!",
			})
		}

		ok, err := hasher.ComparePasswords(dbRecord.PasswordHash, body.Password, dbRecord.Salt, nil)
		if err != nil {
			if devMode {
				panic(err)
			}
		}
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Bad Credentials!",
			})
		} else {
			accessToken, err := token.New(cfg.SingingSecret, "", dbRecord.Username, dbRecord.ID, 45)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "An error occured! Try again later!",
				})
			}
			refreshToken, err := token.New(cfg.SingingSecret, "", dbRecord.Username, dbRecord.ID, 180)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "An error occured! Try again later!",
				})
			}

			return c.Status(fiber.StatusCreated).JSON(fiber.Map{
				"access_token":  accessToken,
				"refresh_token": refreshToken,
			})
		}
	})

	web.Server.Post("/refresh", func(c *fiber.Ctx) (err error) {
		c.Accepts("application/json")
		c.AcceptsCharsets("utf-8")

		tokenHeader := c.Get("Authorization", "")
		if !strings.HasPrefix(tokenHeader, "Bearer ") {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Your Authorization Header is invalid!",
			})
		}

		return
	})

}

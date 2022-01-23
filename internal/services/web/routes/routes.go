package routes

import (
	"github.com/meoera/doorman/pkg/token"

	"github.com/gofiber/fiber/v2"
	"github.com/meoera/doorman/internal/services/config"
	"github.com/meoera/doorman/internal/services/database"
	"github.com/meoera/doorman/pkg/hasher"
	"github.com/meoera/doorman/pkg/models"
)

func Add(app *fiber.App, cfg *config.Web, db database.Database, cacheDb database.CacheDatabase) {

	app.Post("/login", func(c *fiber.Ctx) (err error) {
		c.Accepts("application/json")
		c.AcceptsCharsets("utf-8")

		body := &models.AuthenticationRequestBody{}
		err = c.BodyParser(body)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.BadCredentialsResponseBody)
		}

		dbRecord, err := db.UserByName(body.Username)
		if dbRecord == nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.BadCredentialsResponseBody)
		} else if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(models.InternalServerErrorResponseBody)
		}

		ok, err := hasher.ComparePasswords(dbRecord.PasswordHash, body.Password, dbRecord.Salt, nil)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(models.BadCredentialsResponseBody)
		} else {
			accessToken, err := token.New(cfg.SingingSecret, "", "", dbRecord.ID, cfg.AccessTokenExpiry)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(models.InternalServerErrorResponseBody)
			}
			refreshToken, err := token.New(cfg.SingingSecret, "", "", dbRecord.ID, cfg.RefreshTokenExpiry)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(models.InternalServerErrorResponseBody)
			}

			return c.Status(fiber.StatusCreated).JSON(fiber.Map{
				"access_token":  accessToken,
				"refresh_token": refreshToken,
			})
		}
	})

	app.Post("/refresh", func(c *fiber.Ctx) (err error) {
		c.Accepts("application/json")
		c.AcceptsCharsets("utf-8")

		body := &models.RefreshRequestBody{}
		err = c.BodyParser(body)

		claims, valid, err := token.Claims(cfg.SingingSecret, body.RefreshToken)
		if err != nil || !valid || claims == nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.BadCredentialsResponseBody)
		}

		uid, ok := claims["uid"].(int)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(models.BadCredentialsResponseBody)
		}

		dbRecord, err := db.UserByID(uid)
		if dbRecord == nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.BadCredentialsResponseBody)
		} else if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(models.InternalServerErrorResponseBody)
		}

		accessToken, err := token.New(cfg.SingingSecret, "", "", dbRecord.ID, cfg.AccessTokenExpiry)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(models.InternalServerErrorResponseBody)
		}
		refreshToken, err := token.New(cfg.SingingSecret, "", "", dbRecord.ID, cfg.RefreshTokenExpiry)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(models.InternalServerErrorResponseBody)
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
				"access_token":  accessToken,
				"refresh_token": refreshToken,
			})
	})

}

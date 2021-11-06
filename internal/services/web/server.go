package web

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var Server *fiber.App = fiber.New(fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, e error) error {
		fmt.Println(e)
		return nil
	},
})


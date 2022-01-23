package inits

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/meoera/doorman/internal/services/config"
	"github.com/meoera/doorman/internal/services/database"
	"github.com/meoera/doorman/internal/services/web/routes"
)

func Web(cfg *config.Web, db database.Database, cacheDb database.CacheDatabase) (err error) {

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			return fmt.Errorf("An error occured: %v\n The context: %v", e, c)
		},
	})
	routes.Add(app, cfg, db, cacheDb)
	return app.Listen(cfg.Host)
}

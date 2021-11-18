package inits

import (
	"github.com/meoera/doorman/internal/services/config"
	"github.com/meoera/doorman/internal/services/database"
	"github.com/meoera/doorman/internal/services/web"
	"github.com/meoera/doorman/internal/services/web/routes"
)

func Web(cfg *config.Web, db database.Database, cacheDb database.CacheDatabase, devMode bool) (err error) {
	routes.Add(cfg, db, cacheDb, devMode)
	return web.Server.Listen(cfg.Host)
}

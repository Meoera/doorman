package inits

import (
	"github.com/meoera/doorman/internal/services/config"
	"github.com/meoera/doorman/internal/services/web"
	"github.com/meoera/doorman/internal/services/web/routes"
)

func InitializeWeb(cfg *config.Web, redisCfg *config.Redis, devMode bool) (err error) {
	routes.Add(cfg, redisCfg, devMode)
	return web.Server.Listen(cfg.Host)
}

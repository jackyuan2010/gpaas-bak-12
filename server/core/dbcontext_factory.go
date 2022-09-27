package core

import (
	"github.com/jackyuan2010/gpaas/server/config"
	gpaasgorm "github.com/jackyuan2010/gpaas/server/gorm"
	gpaaspostgres "github.com/jackyuan2010/gpaas/server/gorm/postgres"
)

func CreateDbContext(cfg config.ServerConfig) gpaasgorm.DbContext {
	var dbcontext gpaasgorm.DbContext
	if cfg.DbType == "postgres" {
		postgresDbCtx := gpaaspostgres.NewDbContext(&cfg.DbConfig)
		dbcontext = &postgresDbCtx
	}
	return dbcontext
}

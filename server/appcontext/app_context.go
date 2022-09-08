package appcontext

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"github.com/jackyuan2010/gpaas/server/config"
	gpaasgorm "github.com/jackyuan2010/gpaas/server/gorm"
	gpaaspostgres "github.com/jackyuan2010/gpaas/server/gorm/postgres"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"github.com/songzhibin97/gkit/cache/local_cache"
)

var (
	APP_VP     *viper.Viper
	APP_CONFIG config.ServerConfig
	APP_DbContext gpaasgorm.DbContext
	APP_REDIS  *redis.Client
	APP_Concurrency_Controller = &singleflight.Group{}
	APP_JWTCache local_cache.Cache
)

func InitAppContext() {
	APP_VP = Viper()
	fmt.Println(APP_CONFIG)
	initDbContext()
	// initRedis()
}

func initDbContext() {
	if(APP_CONFIG.DbType == "postgres") {
		dbcontext := gpaaspostgres.NewDbContext(&APP_CONFIG.DbConfig)
		APP_DbContext = &dbcontext
	}
}

func initRedis() {

}
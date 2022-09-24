package appcontext

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jackyuan2010/gpaas/server/config"
	gpaasgorm "github.com/jackyuan2010/gpaas/server/gorm"
	gpaaspostgres "github.com/jackyuan2010/gpaas/server/gorm/postgres"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"golang.org/x/sync/singleflight"
)

var (
	APP_VP                     *viper.Viper
	APP_CONFIG                 config.ServerConfig
	APP_DbContext              gpaasgorm.DbContext
	APP_REDIS                  *redis.Client
	APP_Concurrency_Controller = &singleflight.Group{}
	APP_JWTCache               local_cache.Cache
)

func InitAppContext() {
	APP_VP = Viper()
	fmt.Println(APP_CONFIG)
	initDbContext()
	initRedis()
}

func initDbContext() {
	if APP_CONFIG.DbType == "postgres" {
		dbcontext := gpaaspostgres.NewDbContext(&APP_CONFIG.DbConfig)
		APP_DbContext = &dbcontext
	}
}

func initRedis() {
	redisCfg := APP_CONFIG.RedisConfig
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pong)
		APP_REDIS = client
	}
}

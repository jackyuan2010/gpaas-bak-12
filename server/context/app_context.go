package context

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v8"
	"github.com/jackyuan2010/gpaas/server/config"
	"github.com/jackyuan2010/gpaas/server/core"
	gpaasgorm "github.com/jackyuan2010/gpaas/server/gorm"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"golang.org/x/sync/singleflight"
	"time"
)

var AppContext appContext

func InitAppContext() {
	AppContext = appContext{
		APP_Concurrency_Controller: &singleflight.Group{},
	}
	// AppContext = appContext{}
	AppContext.initViper()
	AppContext.initDbContext()
	AppContext.initRedis()
	AppContext.initJWTCache()
}

type appContext struct {
	APP_VP                     *viper.Viper
	APP_CONFIG                 config.ServerConfig
	APP_DbContext              gpaasgorm.DbContext
	APP_REDIS                  *redis.Client
	APP_Concurrency_Controller *singleflight.Group
	APP_JWTCache               local_cache.Cache
}

func (ctx appContext) initViper() {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("./config/")
	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error read config file: %s", err))
	}

	vp.WatchConfig()

	vp.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := vp.Unmarshal(&ctx.APP_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := vp.Unmarshal(&ctx.APP_CONFIG); err != nil {
		fmt.Println(err)
	}
	ctx.APP_VP = vp
}

func (ctx appContext) initDbContext() {
	core.CreateDbContext(ctx.APP_CONFIG)
}

func (ctx appContext) initRedis() {
	redisCfg := ctx.APP_CONFIG.RedisConfig
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
		ctx.APP_REDIS = client
	}
}

func (ctx appContext) initJWTCache() {
	ctx.APP_JWTCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(time.Second * time.Duration(ctx.APP_CONFIG.JWTConfig.ExpiresTime)),
	)
}

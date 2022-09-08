package appcontext

import (
	"fmt"
	"time"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
	"github.com/songzhibin97/gkit/cache/local_cache"
)

func Viper() *viper.Viper {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("./config/")
	err := vp.ReadInConfig()
	if(err != nil) {
		panic(fmt.Errorf("Fatal error read config file: %s", err))
	}

	vp.WatchConfig()

	vp.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := vp.Unmarshal(&APP_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := vp.Unmarshal(&APP_CONFIG); err != nil {
		fmt.Println(err)
	}

	APP_JWTCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(time.Second * time.Duration(APP_CONFIG.JWTConfig.ExpiresTime)),
	)
	return vp
} 
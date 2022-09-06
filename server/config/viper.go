package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
)


func Viper() *viper.Viper {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("./server/config/")
	err := vp.ReadInConfig()
	if(err != nil) {
		panic(fmt.Errorf("Fatal error read config file: %s", err))
	}
	serverConfig := &ServerConfig{}

	vp.WatchConfig()

	vp.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := vp.Unmarshal(serverConfig); err != nil {
				fmt.Println(err)
		}
	})


	if err := vp.Unmarshal(serverConfig); err != nil {
		fmt.Println(err)
	}
	return vp
} 
package config

import (
	gpaasgorm "github.com/jackyuan2010/gpaas/server/gorm"
)

type ServerConfig struct {
	DbType      string             `mapstructure:"dbtype" json:"dbtype" yaml:"dbtype"`
	DbConfig    gpaasgorm.DbConfig `mapstructure:"dbconfig" json:"dbconfig" yaml:"dbconfig"`
	JWTConfig   JWTConfig          `mapstructure:"jwtconfig" json:"jwtconfig" yaml:"jwtconfig"`
	RedisConfig RedisConfig        `mapstructure:"redisconfig" json:"redisconfig" yaml:"redisconfig"`
}

package gorm

import (
	"gorm.io/gorm"
)

type DbContext interface {
	GetDb() *gorm.DB
	DSN() string 
}
package gorm

import (
	model "github.com/jackyuan2010/gpaas/server/gorm/model"
)

type Repository interface {
	QueryById(id string) *model.User
	QueryList() *[]model.User
}
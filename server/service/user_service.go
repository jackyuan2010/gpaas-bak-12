package service

import (
	gpaasgorm "github.com/jackyuan2010/gpaas/server/gorm"
)

type UserService struct {
	userRepository gpaasgorm.UserRepository
}

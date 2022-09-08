package model

import (
	uuid "github.com/satori/go.uuid"
	"github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	jwt.StandardClaims
	UUID       uuid.UUID
	Mobile     string
	UserName   string
	BufferTime int64
}
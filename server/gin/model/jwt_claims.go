package model

import (
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

type JWTClaims struct {
	jwt.StandardClaims
	UUID       uuid.UUID
	Mobile     string
	UserName   string
	BufferTime int64
}

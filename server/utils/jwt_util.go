package utils

import (
	"errors"
	"fmt"
	"time"
	ginmodel "github.com/jackyuan2010/gpaas/server/gin/model"
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/jackyuan2010/gpaas/server/appcontext"
)

type JWTUtil struct {
	SecretKey []byte
	ExpiresTime int64
	BufferTime int64
	Issuer     string
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValid    = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWTUtil() *JWTUtil {
	return &JWTUtil{
		SecretKey: []byte(appcontext.APP_CONFIG.JWTConfig.SecretKey),
		ExpiresTime: appcontext.APP_CONFIG.JWTConfig.ExpiresTime,
		BufferTime: appcontext.APP_CONFIG.JWTConfig.BufferTime,
		Issuer: appcontext.APP_CONFIG.JWTConfig.Issuer,
	}
}

func (jwtUtil *JWTUtil) CreateClaims(mobile string, userName string) ginmodel.JWTClaims {
	claims := ginmodel.JWTClaims{
		UUID: uuid.NewV4(),
		Mobile: mobile,
		UserName: userName,
		BufferTime: jwtUtil.BufferTime, //缓冲时间内会获得新的token刷新令牌
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                 // 签名生效时间
			ExpiresAt: time.Now().Unix() + jwtUtil.ExpiresTime, // 过期时间
			Issuer:    jwtUtil.Issuer,                          // 签名的发行者
		},
	}
	return claims
}

func (jwtUtil *JWTUtil) GernerateToken(claims ginmodel.JWTClaims) string {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtUtil.SecretKey)
	if err != nil {
		fmt.Println(err)
	}
	return token
}

func (jwtUtil *JWTUtil) Refresh(oldToken string, claims ginmodel.JWTClaims) (string, error) {
	v, err, _ := appcontext.APP_Concurrency_Controller.Do("JWT:"+oldToken, func() (interface{}, error) {
		return jwtUtil.GernerateToken(claims), nil
	})
	return v.(string), err
}

func (jwtUtil *JWTUtil) ParseToken(tokenString string) (*ginmodel.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &ginmodel.JWTClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtUtil.SecretKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValid
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*ginmodel.JWTClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
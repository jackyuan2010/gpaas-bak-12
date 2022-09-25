package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jackyuan2010/gpaas/server/gin/controller/v1"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	userController := v1.UserController{}
	baseRouter.POST("login", userController.Login)
	baseRouter.GET("captcha", userController.Captcha)

	return baseRouter
}

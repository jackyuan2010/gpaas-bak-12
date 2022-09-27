package core

import (
	"github.com/gin-gonic/gin"
	ginrouter "github.com/jackyuan2010/gpaas/server/gin/router"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	// {
	// 	systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	// 	systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	// }
	// baseRouter := PublicGroup.Group("base")

	sysBaseRouter := ginrouter.BaseRouter{}
	sysBaseRouter.InitBaseRouter(PublicGroup)
	// baseRouter.GET("/Login", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Login",
	// 	})
	// })
	return Router
}

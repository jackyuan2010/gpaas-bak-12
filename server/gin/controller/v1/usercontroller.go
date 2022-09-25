package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	dto "github.com/jackyuan2010/gpaas/server/gin/dto"
	model "github.com/jackyuan2010/gpaas/server/gin/model"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

type UserController struct {
}

func (t *UserController) Login(c *gin.Context) {
	var loginDto dto.LoginDto
	_ = c.ShouldBindJSON(&loginDto)
	fmt.Println(loginDto)
}

func (t *UserController) Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		model.Fail(c, "验证码获取失败", map[string]interface{}{})
	} else {
		model.Ok(c, "验证码获取成功", model.CaptchaResponse{
			CaptchaId:     id,
			PicPath:       b64s,
			CaptchaLength: 4,
		})
	}
}

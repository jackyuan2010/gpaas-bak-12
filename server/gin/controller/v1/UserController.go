package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
	dto "github.com/jackyuan2010/gpaas/server/gin/dto"
)

type UserController struct {

}

func (controller *UserController) Login(ctx *gin.Context) {
	var loginDto dto.LoginDto
	_ = c.ShouldBindJSON(&loginDto)
}
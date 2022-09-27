package context

import (
	"github.com/gin-gonic/gin"
	// dto "github.com/jackyuan2010/gpaas/server/gin/dto"
	"github.com/satori/go.uuid"
)

type RequestContext struct {
	RequestUUID uuid.UUID
	User        User
	ClientInfo  ClientInfo
}

type User struct {
	EnterpriseId string
	UserId       string
	Mobile       string
}

type ClientInfo struct {
	ClientIP string
	DeviceId string
	DeviceOS string
}

func NewRequestContext(c *gin.Context) RequestContext {
	ctx := RequestContext{RequestUUID: uuid.NewV4()}

	clientInfo := ClientInfo{
		ClientIP: c.ClientIP(),
		DeviceId: c.Request.Header.Get("x-device-id"),
		DeviceOS: c.Request.Header.Get("x-device-os"),
	}
	ctx.ClientInfo = clientInfo

	return ctx
}

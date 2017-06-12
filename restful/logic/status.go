package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/songtianyi/go-aida/restful/manager"
	"net/http"
)

type Plugin struct {
	name    string
	enabled bool
}

func Status(c *gin.Context) {
	uuid := c.Query("uuid")
	session := manager.GlobalSessionManager.Get(uuid)
	if session == nil {
		c.String(http.StatusNotFound, fmt.Sprintf("no session for %s", uuid))
		return
	}
	if session.Cookies == nil {
		c.JSON(200, gin.H{
			"status":    "CREATED",
			"qrcode":    session.QrcodePath,
			"startTime": session.CreateTime,
		})
		return
	}

	handlerWrappers := session.HandlerRegister.GetAll()
	plugins := make([]gin.H, 0)
	for _, v := range handlerWrappers {
		plugin := gin.H{
			"name":    v.GetName(),
			"enabled": v.GetEnabled(),
		}
		plugins = append(plugins, plugin)
	}

	c.JSON(200, gin.H{
		"status":    "SERVING",
		"startTime": session.CreateTime,
		"plugins":   plugins,
	})
	return
}

package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/songtianyi/go-aida/restful/manager"
	"net/http"
)

func Disable(c *gin.Context) {
	uuid := c.Query("uuid")
	name := c.Query("name")
	session := manager.GlobalSessionManager.Get(uuid)
	if err := session.HandlerRegister.DisableByName(name); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(200, "OK")
	return
}

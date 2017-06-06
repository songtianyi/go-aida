package logic
import (
	"github.com/gin-gonic/gin"
	"github.com/songtianyi/wechat-go/wxweb"
	"net/http"
	"github.com/songtianyi/go-aida/restful/manager"
	"github.com/songtianyi/rrframework/logs"
	"time"
)
func Create(c *gin.Context) {

	// create session
	session, err := wxweb.CreateSession(nil, nil, wxweb.WEB_MODE)
	if err != nil {
		logs.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	LoadAllPlugins(session)
	sessionId := manager.GlobalSessionManager.Add(session);

	go func() {
		for {
			if err := session.LoginAndServe(false); err != nil {
				logs.Error("session exit, %s", err)
				for i := 0; i < 3; i++ {
					logs.Info("trying re-login with cache")
					if err := session.LoginAndServe(true); err != nil {
						logs.Error("re-login error, %s", err)
					}
					time.Sleep(3 * time.Second)
				}
				if session, err = wxweb.CreateSession(nil, session.HandlerRegister, wxweb.WEB_MODE); err != nil {
					logs.Error("create new sesion failed, %s", err)
					break
				}
			} else {
				logs.Info("closed by user")
				break
			}
		}
	}()

	c.String(http.StatusOK, sessionId)
}

package main

import (
	"github.com/songtianyi/rrframework/logs"
	"github.com/songtianyi/wechat-go/plugins/faceplusplus"
	"github.com/songtianyi/wechat-go/wxweb"
)

func main() {
	session, err := wxweb.CreateSession(nil, wxweb.TERMINAL_MODE)
	if err != nil {
		logs.Error(err)
		return
	}
	// add plugins
	faceplusplus.Register(session)

	go func() {
		for {
			select {
			case <-session.RefreshFlag:
				old := session
				session, err = wxweb.CreateSession(nil, wxweb.TERMINAL_MODE)
				if err != nil {
					logs.Error(err)
				} else {
					old.Close()
				}
			}
		}
	}()
	for {
		if err := session.LoginAndServe(); err != nil {
			logs.Error("session exit, %s", err)
		}
	}
}

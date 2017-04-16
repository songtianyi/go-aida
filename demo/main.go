package main

import (
	"github.com/songtianyi/rrframework/logs"
	"github.com/songtianyi/wechat-go/plugins/wxweb/faceplusplus"
	"github.com/songtianyi/wechat-go/plugins/wxweb/replier"
	"github.com/songtianyi/wechat-go/wxweb"
)

func main() {
	// create session
	session, err := wxweb.CreateSession(nil, wxweb.TERMINAL_MODE)
	if err != nil {
		logs.Error(err)
		return
	}
	// add plugins for this session
	faceplusplus.Register(session)
	replier.Register(session)

	// watch refresh flag
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
		// watch session
		if err := session.LoginAndServe(); err != nil {
			logs.Error("session exit, %s", err)
		}
	}
}

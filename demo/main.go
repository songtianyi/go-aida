package main

import (
	"github.com/songtianyi/rrframework/logs"
	"github.com/songtianyi/wechat-go/plugins/wxweb/faceplusplus"
	"github.com/songtianyi/wechat-go/plugins/wxweb/gifer"
	"github.com/songtianyi/wechat-go/plugins/wxweb/replier"
	"github.com/songtianyi/wechat-go/plugins/wxweb/switcher"
	"github.com/songtianyi/wechat-go/wxweb"
)

func main() {
	// create session
	session, err := wxweb.CreateSession(nil, wxweb.TERMINAL_MODE)
	if err != nil {
		logs.Error(err)
		return
	}
	// load plugins for this session
	faceplusplus.Register(session)
	replier.Register(session)
	switcher.Register(session)
	gifer.Register(session)

	// enable plugin
	session.HandlerRegister.EnableByName("switcher")
	session.HandlerRegister.EnableByName("faceplusplus")

	// disable by type example
	//if err := session.HandlerRegister.EnableByType(wxweb.MSG_TEXT); err != nil {
	//	logs.Error(err)
	//}

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

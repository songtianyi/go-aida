package main

import (
	"github.com/songtianyi/rrframework/logs"
	"github.com/songtianyi/wechat-go/plugins/wxweb/cleaner"
	"github.com/songtianyi/wechat-go/plugins/wxweb/demo"
	"github.com/songtianyi/wechat-go/plugins/wxweb/faceplusplus"
	"github.com/songtianyi/wechat-go/plugins/wxweb/gifer"
	"github.com/songtianyi/wechat-go/plugins/wxweb/laosj"
	"github.com/songtianyi/wechat-go/plugins/wxweb/replier"
	"github.com/songtianyi/wechat-go/plugins/wxweb/switcher"
	"github.com/songtianyi/wechat-go/wxweb"
)

func main() {
	// create session
	session, err := wxweb.CreateSession(nil, nil, wxweb.TERMINAL_MODE)
	if err != nil {
		logs.Error(err)
		return
	}
	// load plugins for this session
	faceplusplus.Register(session)
	replier.Register(session)
	switcher.Register(session)
	gifer.Register(session)
	cleaner.Register(session)
	laosj.Register(session)
	demo.Register(session)

	// enable plugin
	session.HandlerRegister.EnableByName("switcher")
	session.HandlerRegister.EnableByName("faceplusplus")
	session.HandlerRegister.EnableByName("cleaner")
	session.HandlerRegister.EnableByName("laosj")

	// disable by type example
	//if err := session.HandlerRegister.EnableByType(wxweb.MSG_TEXT); err != nil {
	//	logs.Error(err)
	//}

	for {
		// watch session
		if err := session.LoginAndServe(); err != nil {
			logs.Error("session exit, %s", err)
			if session, err = wxweb.CreateSession(nil, session.HandlerRegister, wxweb.TERMINAL_MODE); err != nil {
				logs.Error("create new sesion failed, %s", err)
				break
			}
		}
	}
}

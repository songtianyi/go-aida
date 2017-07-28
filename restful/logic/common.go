package logic

import (
	"github.com/songtianyi/wechat-go/plugins/wxweb/faceplusplus"
	"github.com/songtianyi/wechat-go/plugins/wxweb/gifer"
	"github.com/songtianyi/wechat-go/plugins/wxweb/joker"
	"github.com/songtianyi/wechat-go/plugins/wxweb/laosj"
	"github.com/songtianyi/wechat-go/plugins/wxweb/replier"
	"github.com/songtianyi/wechat-go/plugins/wxweb/revoker"
	"github.com/songtianyi/wechat-go/plugins/wxweb/verify"
	"github.com/songtianyi/wechat-go/plugins/wxweb/youdao"
	"github.com/songtianyi/wechat-go/wxweb"
)

func LoadAllPlugins(session *wxweb.Session) {
	// load plugins for this session
	faceplusplus.Register(session)
	replier.Register(session)
	gifer.Register(session)
	laosj.Register(session)
	joker.Register(session)
	revoker.Register(session)
	youdao.Register(session)
	verify.Register(session)
}

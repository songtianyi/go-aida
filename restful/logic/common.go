package logic
import (
	"github.com/songtianyi/wechat-go/wxweb"
	"github.com/songtianyi/rrframework/logs"
	"github.com/songtianyi/wechat-go/plugins/wxweb/faceplusplus"
	"github.com/songtianyi/wechat-go/plugins/wxweb/forwarder"
	"github.com/songtianyi/wechat-go/plugins/wxweb/gifer"
	"github.com/songtianyi/wechat-go/plugins/wxweb/joker"
	"github.com/songtianyi/wechat-go/plugins/wxweb/laosj"
	"github.com/songtianyi/wechat-go/plugins/wxweb/replier"
	"github.com/songtianyi/wechat-go/plugins/wxweb/revoker"
	"github.com/songtianyi/wechat-go/plugins/wxweb/switcher"
	"github.com/songtianyi/wechat-go/plugins/wxweb/system"
	"github.com/songtianyi/wechat-go/plugins/wxweb/youdao"
)

func LoadAllPlugins(session *wxweb.Session) {
	// load plugins for this session
	faceplusplus.Register(session)
	replier.Register(session)
	switcher.Register(session)
	gifer.Register(session)
	laosj.Register(session)
	joker.Register(session)
	revoker.Register(session)
	forwarder.Register(session)
	system.Register(session)
	youdao.Register(session)

	// enable plugin
	session.HandlerRegister.EnableByName("switcher")
	session.HandlerRegister.EnableByName("faceplusplus")
	session.HandlerRegister.EnableByName("laosj")
	session.HandlerRegister.EnableByName("joker")
	//session.HandlerRegister.EnableByName("youdao")

	// enable by type example
	if err := session.HandlerRegister.EnableByType(wxweb.MSG_SYS); err != nil {
		logs.Error(err)
		return
	}
}

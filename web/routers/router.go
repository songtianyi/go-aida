package routers

import (
	"github.com/astaxie/beego"
	"github.com/songtianyi/go-aida/web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}

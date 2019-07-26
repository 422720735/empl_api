package routers

import (
	"empl_api/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 注册路由
    beego.Router("/emp/reg", &controllers.RegControllers{},"post:AddUser")
    beego.Router("/emp/login", &controllers.LoginControllers{}, "post:Verify")
}

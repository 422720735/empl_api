package routers

import (
	"empl_api/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 注册路由
    beego.Router("/reg", &controllers.UserControllers{},"post:AddUser")
}

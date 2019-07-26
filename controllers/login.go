package controllers

import (
	"empl_api/common"
	"empl_api/global"
	"empl_api/models"
	"empl_api/transform"
	"github.com/astaxie/beego"
)

type LoginControllers struct {
	beego.Controller
}

func (c *LoginControllers) Verify() {
	msg, err := common.Unmarshal(&c.Controller)
	if err != nil {
		common.Echo(&c.Controller, 201, "获取数据失败")
		return
	}
	username, err := transform.InterToString(msg["username"])
	if err != nil {
		common.Echo(&c.Controller, 201, "用户名不合法")
		return
	}
	password, err := transform.InterToString(msg["password"])
	if err != nil {
		common.Echo(&c.Controller, 201, "密码不合法")
		return
	}
	// 密码加密
	newMd5Pwd := common.Md5(password, beego.AppConfig.String("MD5PASSWORD"))
	user, err := models.VerifyUser(username, newMd5Pwd)
	if err != nil {
		common.Echo(&c.Controller, 201, "查询失败")
		return
	}
	if user.Username == "" || user.Password == "" || user.Id == 0 {
		common.Echo(&c.Controller, 201, "用户名或密码不正确")
		return
	}
	// 这里要去注册一个token返给前端
	//common.Echo(&c.Controller,201, "登陆成功")
	token, err := common.NewToken(user, global.G_Sign, global.G_AesKey, 60*60*24*15)
	if err != nil {
		common.Echo(&c.Controller, 201, "初始化token失败，请重新登录")
		return
	}
	data := make(map[string]interface{})
	data["token"] = token
	data["position"] = user.Position
	common.Echo(&c.Controller, 200, data)
}

package controllers

import (
	"empl_api/common"
	"empl_api/transform"
	"github.com/astaxie/beego"
)

type UserControllers struct {
	beego.Controller
}

type User struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

func (c *UserControllers) AddUser() {
	msg, err := common.Unmarshal(&c.Controller)
	if err != nil {
		return
	}
	username, err := transform.InterToString(msg["username"])
	if err != nil {
		common.Echo(&c.Controller, 201, "用户名或密码不合法")
	}
	password, err := transform.InterToString(msg["password"])
	if err != nil {
		common.Echo(&c.Controller, 201, "用户名或密码不合法")
	}
	beego.Info("----------------", msg)
	common.Echo(&c.Controller, 200, &User{Name: username, Pwd: password})
}

package controllers

import (
	"empl_api/common"
	"empl_api/models"
	"empl_api/transform"
	"github.com/astaxie/beego"
)

type RegControllers struct {
	beego.Controller
}

const (
	MD5PASSWORD = "1@vfgyh[;][;hiueugu%_fuhuuh}.fe;efdf"
)

func (c *RegControllers) AddUser() {
	msg, err := common.Unmarshal(&c.Controller)
	if err != nil {
		common.Echo(&c.Controller, 201, "获取数据失败")
		return
	}
	username, err := transform.InterToString(msg["username"])
	if err != nil {
		common.Echo(&c.Controller, 201, "用户名或密码不合法")
		return
	}
	password, err := transform.InterToString(msg["password"])
	if err != nil {
		common.Echo(&c.Controller, 201, "用户名或密码不合法")
		return
	}
	if username == "" || password == "" {
		common.Echo(&c.Controller, 201, "用户名或密码不能为空")
		return
	}
	if len(username) < 6 || len(password) < 6 {
		common.Echo(&c.Controller, 201, "用户命或密码不能少于6位")
		return
	}
	phone, err := transform.InterToString(msg["phone"])

	ok := common.VerifyMobileFormat(phone)
	if !ok {
		common.Echo(&c.Controller, 201, "手机号码不正确")
		return
	}
	// 职位
	position, err := transform.InterToInt(msg["position"])
	if err != nil {
		common.Echo(&c.Controller, 201, "职位不合法")
		return
	}
	// 验证用户名是不是在数据库存在
	usercount,err := models.UserCount(username)
	if err != nil {
		common.Echo(&c.Controller, 201, "参数不合法，请重新注册")
		return
	}
	if usercount > 0 {
		common.Echo(&c.Controller, 201, "该用户名已被占用")
		return
	}
	// 密码加密
	newMd5Pwd := common.Md5(password, beego.AppConfig.String("MD5PASSWORD"))
	err = models.AddUser(username, newMd5Pwd, phone, position)
	if err != nil {
		common.Echo(&c.Controller, 201, "新增用户失败，请重新数据")
		return
	}
	common.Echo(&c.Controller, 200, "新增用户成功")
}

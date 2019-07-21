package common

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"regexp"
	"time"
)

//返回消息
func Echo(c *beego.Controller, code int, body interface{}) {
	c.Data["json"] = map[string]interface{}{
		"code": code,
		"msg":  body,
		"time": time.Now().Unix(),
	}
	c.ServeJSON()
}

// 获取body数据
func Unmarshal(c *beego.Controller) (map[string]interface{}, error) {
	msg := make(map[string]interface{})
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

// 验证手机号
func VerifyMobileFormat(mobileNum int) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(string(mobileNum))
}

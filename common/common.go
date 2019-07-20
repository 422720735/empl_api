package common

import (
	"encoding/json"
	"github.com/astaxie/beego"
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

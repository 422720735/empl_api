package global

import "github.com/astaxie/beego"

var (
	G_AesKey     string
	G_Sign       string
)

func init()  {
	G_AesKey = beego.AppConfig.String("aes_key")
	G_Sign = beego.AppConfig.String("sign")
}
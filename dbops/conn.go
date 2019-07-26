package dbops

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DbConn *sql.DB
	err    error
)

func init() {
	address := beego.AppConfig.String("address")
	port := beego.AppConfig.String("port")
	DbConn, err = sql.Open("mysql", "root:123456@tcp("+address+":"+port+")/empl?charset=utf8")
	if err != nil {
		beego.Info("数据库异常")
		panic(err.Error())
	}
}

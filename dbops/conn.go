package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DbConn *sql.DB
	err    error
)

func init() {
	DbConn, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/empl?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}

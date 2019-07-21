package models

import (
	"empl_api/dbops"
	"fmt"
)

func AddUser(username, password string, phone, position int) error {
	stmtIns, err := dbops.DbConn.Prepare("insert into users (user_name, pwd, phone, `position`)values (?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(username, password, phone, position)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	fmt.Println("保存成功")
	return nil
}

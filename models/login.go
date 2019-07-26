package models

import (
	"empl_api/dbops"
	"empl_api/defs"
	"log"
)

/*数据库进行密码比对*/
func VerifyUser(username, password string) (*defs.Users, error) {
	stmtOut, err := dbops.DbConn.Prepare("select id,username,password,phone,`position` from users where username = ? and password = ?")
	if err != nil {
		log.Panicf("%s", err)
		return nil, err
	}
	var name, pwd, phone string
	var id, position int
	err = stmtOut.QueryRow(username, password).Scan(&id, &name, &pwd, &phone, &position)
	if err != nil {
		return nil, err
	}
	defer stmtOut.Close()
	res := &defs.Users{Id: id, Username: username, Password: pwd, Phone: phone, Position: position}
	return res, nil
}

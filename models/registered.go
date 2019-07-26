package models

import (
	"empl_api/dbops"
	"log"
)

func AddUser(username, password, phone string, position int) error {
	stmtIns, err := dbops.DbConn.Prepare("insert into users (username, password, phone, `position`)values (?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(username, password, phone, position)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func UserCount(username string) (int, error) {
	// select pwd from users where login_name =?
	// select count(user_name) from users where user_name='123456';
	stmtOut, err := dbops.DbConn.Prepare("select count(username) from users where username =?")
	if err != nil {
		log.Panicf("%s", err)
		return -1, err
	}
	var count int
	stmtOut.QueryRow(username).Scan(&count)
	defer stmtOut.Close()
	return count, nil
}

package models

import (
	"empl_api/dbops"
	"testing"
)

func clearTables() {
	dbops.DbConn.Exec("truncate users")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	err:= AddUser("bob","liu", 13880821491, 1)
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

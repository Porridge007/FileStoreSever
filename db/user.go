package db

import (
	mydb "FileStoreSever/db/mysql"
	"fmt"
)

// sign up with username and password
func UserSignUp(username string, password string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_user (`user_name`, `user_pwd`) values (?,?)")
	if err != nil{
		fmt.Println("Failed to insert, err: ", err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(username, password)
	if err!=nil{
		fmt.Println("Failed to insert,err:", err.Error())
		return false
	}
	if rowsAffected, err := ret.RowsAffected(); nil == err&& rowsAffected>0{
		return true
	}
	return false
}

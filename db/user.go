package db

import (
	mydb "FileStoreSever/db/mysql"
	"fmt"
)

type User struct {
	Username     string
	Email        string
	Phone        string
	SignupAt     string
	LastActiveAt string
	Status       int
}

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

// check the password consistence
func UserSignIn(username string, encpwd string) bool{
	stmt, err := mydb.DBConn().Prepare("select * from tbl_user where user_name=? limit 1")
	if err != nil{
		fmt.Println(err.Error())
		return false
	}
	rows,err:=stmt.Query(username)
	if err!=nil{
		fmt.Println(err.Error())
		return false
	}else if rows == nil{
		fmt.Println("username not found:"+username)
		return false
	}

	pRows := mydb.ParseRows(rows)
	if len(pRows) >0 && string(pRows[0]["user_pwd"].([]byte))==encpwd{
		return true
	}
	return false
}

// update user's token
func UpdateToken(username string, token string) bool{
	stmt,err := mydb.DBConn().Prepare(
		"replace into tbl_user_token (`user_name`,`user_token`) values (?,?)")
	if err != nil{
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, token)
	if err != nil{
		fmt.Println(err.Error())
		return false
	}
	return true
}

// GetUserInfo : Query user info
func GetUserInfo(username string) (User, error) {
	user := User{}

	stmt, err := mydb.DBConn().Prepare(
		"select user_name,signup_at from tbl_user where user_name=? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}
	defer stmt.Close()

	// exec the query
	// todo: query more fields
	err = stmt.QueryRow(username).Scan(&user.Username, &user.SignupAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

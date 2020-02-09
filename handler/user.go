package handler

import (
	"FileStoreSever/util"
	"io/ioutil"
	"net/http"
	dblayer "FileStoreSever/db"
	mydb "FileStoreSever/db/mysql"
	"fmt"
	"time"
	"strconv"
)

const (
	pwdSalt = "*#1989"
)

// handle user signing up
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}
	r.ParseForm()
	username := r.Form.Get("username")
	passwd := r.Form.Get("password")

	if len(username) < 3 || len(passwd) < 5 {
		w.Write([]byte("Invalid parameter"))
		return
	}
	encPasswd := util.Sha1([]byte(passwd + pwdSalt))
	suc := dblayer.UserSignUp(username, encPasswd)
	if suc {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("FAIL"))
	}
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// data, err := ioutil.ReadFile("./static/view/signin.html")
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }
		// w.Write(data)
		http.Redirect(w, r, "/static/view/signin.html", http.StatusFound)
		return
	}

	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	encPasswd := util.Sha1([]byte(password + pwdSalt))

	// 1. 校验用户名及密码
	pwdChecked := dblayer.UserSignIn(username, encPasswd)
	if !pwdChecked {
		w.Write([]byte("FAILED"))
		return
	}

	// 2. 生成访问凭证(token)
	token := GenToken(username)
	upRes := dblayer.UpdateToken(username, token)
	if !upRes {
		w.Write([]byte("FAILED"))
		return
	}

	// 3. 登录成功后重定向到首页
	//w.Write([]byte("http://" + r.Host + "/static/view/home.html"))
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: struct {
			Location string
			Username string
			Token    string
		}{
			Location: "http://" + r.Host + "/static/view/home.html",
			Username: username,
			Token:    token,
		},
	}
	w.Write(resp.JSONBytes())
}

// Query user information
func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	// 1. parse request params
	r.ParseForm()
	username := r.Form.Get("username")
	token := r.Form.Get("token")
	// 2. check the toke
	if !IsTokenValid(token) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// 3. query user info
	user, err := dblayer.GetUserInfo(username)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// 4. form user info and response
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: user,
	}
	w.Write(resp.JSONBytes())
}

func GenToken(username string) string {
	// 40位字符 md5(username+timestamp+token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}

func IsTokenValid(token string) bool {
	if len(token) != 40 {
		return false
	}
	//  判断token的时效性，是否过期
	tsNow := time.Now().Unix()
	tokenTime,_ := strconv.ParseInt(token[31:],16,64)
	if tsNow - tokenTime>86400{
		fmt.Println(tsNow - tokenTime)
		return false
	}
	// 从数据库表tbl_user_token查询username对应的token信息
	var tokenDb string
	err := mydb.DBConn().QueryRow("select user_token from tbl_user_token").Scan(&tokenDb)
	if err != nil {
		fmt.Println(err)
	}

	//  对比两个token是否一致
	if token != tokenDb{
		return  false
	}
	return true
}

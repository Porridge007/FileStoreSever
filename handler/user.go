package handler

import (
	"FileStoreSever/util"
	"io/ioutil"
	"net/http"
	dblayer "FileStoreSever/db"
)

const(
	pwdSalt ="*#1989"
)

// handle user signing up
func SignUpHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodGet{
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}
	r.ParseForm()
	username := r.Form.Get("username")
	passwd := r.Form.Get("password")

	if len(username) <3 || len(passwd) <5{
		w.Write([]byte("Invalid parameter"))
		return
	}
	enc_passwd := util.Sha1([]byte(passwd+pwdSalt))
	suc := dblayer.UserSignUp(username, enc_passwd)
	if suc{
		w.Write([]byte("SUCCESS"))
	}else {
		w.Write([]byte("FAIL"))
	}
}
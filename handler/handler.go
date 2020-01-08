package handler

import (
	"io"
	"io/ioutil"
	"net/http"
)

// handle upload file
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//return upload html
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internel server error")
			return
		}
		io.WriteString(w,string(data))
	} else if r.Method == "POST" {
		//restore stream-file to local storage
	}
}

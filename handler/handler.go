package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		//restore stream-file to local storage
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Failed to get data, err:", err.Error())
			return
		}
		defer file.Close()
		newFile, err := os.Create("../../../Storage/" + head.Filename)
		if err != nil {
			fmt.Println("Failed to create file, err:", err.Error())
			return
		}
		defer newFile.Close()
		_,err = io.Copy(newFile, file)
		if  err!= nil{
			fmt.Println("Failed to save data into file, err:",err.Error())
			return
		}
		http.Redirect(w,r,"/file/upload/suc", http.StatusFound)
	}
}


func UploadSucHandler(w http.ResponseWriter, r *http.Request){
	io.WriteString(w,"Upload finished!")
}
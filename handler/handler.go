package handler

import (
	"FileStoreSever/meta"
	"FileStoreSever/util"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
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
		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: "../../../Storage/" + head.Filename,
			UploadAt: time.Now().Format("2016-01-02 15:06:07"),
		}

		newFile, err := os.Create("../../../Storage/" + fileMeta.Location)
		if err != nil {
			fmt.Println("Failed to create file, err:", err.Error())
			return
		}
		defer newFile.Close()
		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Println("Failed to save data into file, err:", err.Error())
			return
		}
		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdateFileMeta(fileMeta)
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload finished!")
}

// Get file meta info
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	filehash := r.Form["filehash"][0]
	fMeta := meta.GetFileMeta(filehash)
	data, err := json.Marshal(fMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

// Query multiple file meta information
func FileQueryHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	limitCnt, _ := strconv.Atoi(r.Form.Get("limit"))
	fileMetas := meta.GetLastFileMetas(limitCnt)
	data, err := json.Marshal(fileMetas)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

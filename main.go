package main

import (
	"FileStoreSever/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/query", handler.FileQueryHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Fail to start server, err:%s", err.Error())
	}
}

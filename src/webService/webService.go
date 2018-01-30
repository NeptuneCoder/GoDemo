package main

import (
	"log"
	"net/http"
	"upload"
)

func main() {
	http.HandleFunc("/upload",upload.UploadFileAndText)
	err := http.ListenAndServe(":9999", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

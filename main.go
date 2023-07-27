package main

import (
	"DNetDisk/hander"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", hander.UploadHander)
	http.HandleFunc("/file/upload/suc", hander.UploadSuHander)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("服务启动失败:%s", err.Error())
	}
}

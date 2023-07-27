package hander

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func UploadHander(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		file, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "index server error")
			return
		}
		io.WriteString(w, string(file))
	} else if r.Method == "POST" {
		file, head, err := r.FormFile("file")
		if err != nil {
			io.WriteString(w, "文件接受失败")
			return
		}
		defer file.Close()
		newfile, err := os.Create("/tmp/" + head.Filename)
		if err != nil {
			io.WriteString(w, "文件接受失败")
			return
		}
		defer newfile.Close()
		_, err = io.Copy(newfile, file)
		if err != nil {
			io.WriteString(w, "文件保存失败")
			return
		}
		http.Redirect(w, r, "file/upload/suc", http.StatusFound)
	}
}

func UploadSuHander(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "上传完成")

}

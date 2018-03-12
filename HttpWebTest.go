package main

import (
	. "fmt"
	"net/http"
	"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数
	Println("Form", r.Form)
	Println("path", r.URL.Path)
	Println("schema", r.URL.Scheme)
	Println(r.Form["url_long"])
	for k, v := range r.Form {		Println("key:", k)
		Println("val:", strings.Join(v, ","))
	}
	Fprintf(w, "Hello YueXiong!")
}

func main() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":9090", nil) //handler 设置成nil，默认获取DefaultServerMux路由
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
	}

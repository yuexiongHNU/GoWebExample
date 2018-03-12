package main

import (
	. "fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"net"
	"os"
)

func sayhelloName3(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm() // 解析Form
	Println("Form ", r.Form)
	Println("path ", r.URL.Path)
	Println("schema ", r.URL.Scheme)
	Println("url_long ", r.Form["url_long"])

	for k, v := range r.Form {
		Println("key:", k)
		Println("value:", strings.Join(v, ","))
	}
	Fprintf(w,"Welcome to this site!")
}

func login(w http.ResponseWriter, r *http.Request) {
	Println("Method", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix() // 返回现在的Unix时间戳
		Println("Time unix:", crutime)
		h := md5.New()
		Println("Md5 new: %s \n", h)
		// strconv.FormatInt 转换int成指定进制的字符串
		// io.WriteString 将strconv.FormatInt(crutime, 10)返回的字符串写入h中
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		Println("Curtime after format:", strconv.FormatInt(crutime, 10))
		Println("After write string: ", h)
		// %x 小写十六进制 %X 大写十六进制
		// Sum(nil) 返回h的MD5 校验码
		token := Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm() //需要显示调用，解析Form才能从From里面读取数据
		token := r.Form.Get("token")
		if token != "" {
			Println("Token:", token)
		} else {
			log.Fatal("Get nil token!")
		}
		Println("Username Length:", len(r.Form["username"][0]))
		Println("Username:", template.HTMLEscapeString(r.Form.Get("username")))
		Println("Password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
}

func LocalIP() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		log.Fatal("Error:", err)
		os.Exit(1)
	}
	for _, address := range addrs {
		Println("addrs:", address)
		// 判断address中是否存在IP网络数据-IPNet
		// 判断IPNet中的IP是不是回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			Println("Ipnet:", ipnet)
			// ipnet.IP.To4()判断是不是IPV4地址,是直接返回；不是返回nil
			if ipnet.IP.To4() != nil {
				Println("IP To4(): ", ipnet.IP.To4())
				return ipnet.IP.String()
			}
		}
	}
	return "Can't get local ip!"
}

func main()  {
	http.HandleFunc("/", sayhelloName3)
	http.HandleFunc("/login", login)
	ip := LocalIP()
	listenPort := 9090
	log.Printf("Server running at http://%s:%d", ip, listenPort)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
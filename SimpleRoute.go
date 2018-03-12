package main

import (
	. "fmt"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path == "/" {
		sayhelloName2(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayhelloName2(w http.ResponseWriter, r *http.Request)  {
	Fprintf(w, "Hello YueXiong")
	Println(r.URL.Path)
	Println(r.Method)
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}

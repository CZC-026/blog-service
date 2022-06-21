package main

import (
	"fmt"
	"log"
	"net/http"
)

//手写handler，定义一个空的路由 重写servehttp方法即可
//根据不同的url可以有不同的跳转

//殊途同归，看源码可以发现，都是重新写了一个handler或者是controller代替原来的default（主要就是重写了serverhttp方法）
//然后通过调用listenandserve函数将handler传递给http库下的server
// Engine is the uni handler for all requests
type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	//engine := new(Engine)
	engine := &Engine{}
	log.Fatal(http.ListenAndServe(":9999", engine))
}

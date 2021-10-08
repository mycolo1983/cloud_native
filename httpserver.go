package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//接收客户端request，并将request中带的header写入response header
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Method: %s ,URL: %s ,Proto: %s \n", r.Method, r.URL, r.Proto)
	//遍历所有header字段
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr= %q\n", r.RemoteAddr)
}

//读取当前系统的环境变量中的Version配置，并写入response header
//Server段记录访问日志包括客户端ip，HTTP返回码，输出到server端的标准输出
//当访问localhost/healthz时，应返回200

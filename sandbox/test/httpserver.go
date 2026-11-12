package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// 根路径处理器
func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("访问路径：%s，来源：%s\n", r.URL.Path, r.RemoteAddr)
	fmt.Fprintf(w, "Hello, World! 👋")
}

// /time 路径处理器
func timeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("访问路径：%s，来源：%s\n", r.URL.Path, r.RemoteAddr)
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(w, "当前服务器时间：%s", currentTime)
}
func main() {
	// 注册路由处理器
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/time", timeHandler)
	// 启动 HTTP 服务
	port := ":8080"
	log.Printf("启动 HTTP 服务，监听端口 %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("服务器启动失败：%v", err)
	}
}

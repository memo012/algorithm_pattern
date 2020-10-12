package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	ctValue := r.Header.Get("Content-Type")
	fmt.Println(ctValue)                  // 请求的内容类型
	fmt.Println(r.ContentLength)          // 请求提长度
	data := make([]byte, r.ContentLength) // 切片的长度是请求体内容的长度
	r.Body.Read(data)                     // 从请求体读数据
	fmt.Printf("%#v\n", data)
	// json反序列化
	ul := new(User)
	json.Unmarshal(data, ul)
	fmt.Printf("%#v\n", ul)
	fmt.Fprintf(w, ctValue)
	//w.WriteHeader(http.StatusOK)
}

func indexHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "This is index page")
	//w.Header().Set("Content-Type", "application/json")
	//w.Header().Get("Content-Type")
}

func main() {
	// 创建Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/json", jsonHandler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux, // 注册处理器
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

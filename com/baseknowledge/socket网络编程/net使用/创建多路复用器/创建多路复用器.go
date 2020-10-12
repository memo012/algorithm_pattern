package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello Web!")
	
}

func main()  {
	// 创建Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	server := &http.Server{
		Addr: ":8080",
		Handler: mux, // 注册处理器
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

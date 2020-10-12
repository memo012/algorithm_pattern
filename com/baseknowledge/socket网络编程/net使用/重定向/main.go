package main

import (
	"log"
	"net/http"
)

func headerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://www.baidu.com")
	w.WriteHeader(http.StatusMovedPermanently) // 301
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/header", headerHandler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux, // 注册处理器
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

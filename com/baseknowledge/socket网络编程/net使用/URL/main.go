package main

import (
	"fmt"
	"log"
	"net/http"
)

func urlHandler(w http.ResponseWriter, r *http.Request) {
	// r *http.Request 所有跟请求相关的都在这里
	// w http.ResponseWriter 所有跟响应相关的都在这里
	URL := r.URL
	fmt.Fprintf(w, "Scheme: %s\n", URL.Scheme)
	fmt.Fprintf(w, "Host: %s\n", URL.Host)
	fmt.Fprintf(w, "Path: %s\n", URL.Path)
	fmt.Fprintf(w, "RawPath: %s\n", URL.RawPath)
	fmt.Fprintf(w, "RawQuery: %s\n", URL.RawQuery)
	fmt.Fprintf(w, "Fragment: %s\n", URL.Fragment)
}

func headerHandler(w http.ResponseWriter, r *http.Request) {
	for key, value := range r.Header {
		fmt.Fprintf(w, "%s:%v\n", key, value)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/url", urlHandler)
	mux.HandleFunc("/header", headerHandler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux, // 注册处理器
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

type GreetingHandler struct {
	Language string
}

func (g GreetingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "%s", g.Language)
}

func main() {
	mux := http.NewServeMux()

	// mux.HandleFunc()
	// mux.Handle(URL, Handler接口:实现了ServerHTTP方法)
	mux.Handle("/chinese", GreetingHandler{
		Language: "你好",
	})
	mux.Handle("/english", GreetingHandler{
		Language: "Hello",
	})
	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

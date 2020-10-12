package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello Web!")
}

func main() {

	// HandleFunc 将hello函数注册到根路径/ 上 hello函数也叫处理器
	http.HandleFunc("/", hello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}


}

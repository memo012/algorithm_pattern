package main

import (
	"fmt"
	"log"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := &http.Cookie{
		Name:     "name",
		Value:    "memolei",
		HttpOnly: true,
	}
	c2 := &http.Cookie{
		Name:     "age",
		Value:    "18",
		HttpOnly: true,
	}

	// 将cookie 写入响应头部
	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())
	//http.SetCookie(w, c1)
	//http.SetCookie(w, c2)

}

func getCookie(w http.ResponseWriter, r *http.Request) {
	name, err := r.Cookie("name")
	if err != nil {
		fmt.Fprintln(w, "cannot get cookie of name")
		return
	}
	cookie := r.Cookies()
	fmt.Println(name.Value)
	fmt.Fprintln(w, name)
	fmt.Fprintln(w, cookie)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cookie", setCookie)
	mux.HandleFunc("/getcookie", getCookie)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}

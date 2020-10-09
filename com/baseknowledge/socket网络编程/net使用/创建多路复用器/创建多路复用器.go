package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello Web!")
	
}

func main()  {
	
}

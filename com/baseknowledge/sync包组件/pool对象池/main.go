package main

import (
	"fmt"
	"sync"
)

func main() {
	var pool sync.Pool
	var val interface{}
	pool.Put(1)
	pool.Put("memolei")
	pool.Put("腾讯")

	for {
		val = pool.Get()
		if val == nil {
			break
		}
		fmt.Println(val)
	}
}

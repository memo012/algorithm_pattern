package main

import (
	"fmt"
	"sync/atomic"
)

var mgr []interface{}
var count uint32

func init() {
	mgr = append(mgr, "8080")
	mgr = append(mgr, "8081")
}

func getSelectStrategy() {
	value := atomic.AddUint32(&count, 1)
	index := value % uint32(len(mgr))
	fmt.Printf("端口号:%s\n", mgr[index])
}

func main() {
	for i := 0; i < 10; i++ {
		getSelectStrategy()
	}
}

package main

import (
	"algorithm_pattern/com/baseknowledge/负载均衡/加权随机/server"
	"fmt"
)

func main() {
	lb := server.AddHttpServers()
	for i := 0; i < 11; i++ {
		fmt.Println(lb.GetHttpServerByRandomWithWeight())
	}
}

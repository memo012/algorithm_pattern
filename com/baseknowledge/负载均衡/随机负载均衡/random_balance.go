package main

import (
	"algorithm_pattern/com/baseknowledge/负载均衡/随机负载均衡/server"
	"fmt"
)

func main() {
	lb := server.AddHttpServers()
	for i := 0; i < 5; i++ {
		fmt.Println(lb.GetHttpServerByRandom())
	}

}

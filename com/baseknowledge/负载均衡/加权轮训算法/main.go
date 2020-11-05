package main

import (
	"algorithm_pattern/com/baseknowledge/负载均衡/加权轮训算法/server"
	"fmt"
)

func main()  {
	s := server.AddHttpServers()
	for i := 0; i < 11; i++ {
		fmt.Println(s.GetHttpServerByRoundRobinWithWeight())
	}
}

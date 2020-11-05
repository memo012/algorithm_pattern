package server

import (
	"math/rand"
	"time"
)

type HttpServer struct {
	Host   string
	Weight int
}

type LoadBalance struct {
	Servers []*HttpServer
}

func NewLoadBalance() *LoadBalance {
	return &LoadBalance{make([]*HttpServer, 0)}
}

func NewHttpServer(host string, weight int) *HttpServer {
	return &HttpServer{Host: host, Weight: weight}
}

func (l *LoadBalance) Add(server *HttpServer) {
	l.Servers = append(l.Servers, server)
}

func (l *LoadBalance) GetHttpServerByRandomWithWeight() string {
	rand.Seed(time.Now().UnixNano())

	// 计算所有节点的权重之和
	weightSum := 0
	for i := 0; i < len(l.Servers); i++ {
		weightSum += l.Servers[i].Weight
	}
	// 随机数获取
	random := rand.Intn(weightSum)

	sum := 0
	for i := 0; i < len(l.Servers); i++ {
		sum += l.Servers[i].Weight
		// 因为区间是[ ) ，左闭右开，故随机数小于当前权重sum值，则代表落在该区间，返回当前的index
		if random < sum {
			return l.Servers[i].Host
		}
	}
	return l.Servers[0].Host
}

func AddHttpServers() *LoadBalance {
	lb := NewLoadBalance()
	lb.Add(NewHttpServer("8080", 1))
	lb.Add(NewHttpServer("8081", 2))
	lb.Add(NewHttpServer("8082", 3))
	lb.Add(NewHttpServer("8083", 4))
	lb.Add(NewHttpServer("8084", 1))
	return lb
}

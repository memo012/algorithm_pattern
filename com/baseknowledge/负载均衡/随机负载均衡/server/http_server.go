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

func (lb *LoadBalance) GetHttpServerByRandom() string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(lb.Servers))
	return lb.Servers[index].Host
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

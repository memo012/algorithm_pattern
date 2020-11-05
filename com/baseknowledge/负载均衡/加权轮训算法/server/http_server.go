package server

type HttpServer struct {
	Host   string
	Weight int
}

type LoadBalance struct {
	Index   int
	Servers []*HttpServer
}

func NewLoadBalance() *LoadBalance {
	return &LoadBalance{0, make([]*HttpServer, 0)}
}

func NewHttpServer(host string, weight int) *HttpServer {
	return &HttpServer{Host: host, Weight: weight}
}

func (l *LoadBalance) Add(server *HttpServer) {
	l.Servers = append(l.Servers, server)
}

func (l *LoadBalance) GetHttpServerByRoundRobinWithWeight() string {
	server := l.Servers[0]
	sum := 0
	for i := 0; i < len(l.Servers); i++ {
		sum += l.Servers[i].Weight
		if l.Index < sum {
			server = l.Servers[i]
			// 落到某个区间的位置为末位置-1 && 当前区间为末区间
			if l.Index == sum-1 && i != len(l.Servers)-1 {
				l.Index++
			} else {
				l.Index = (l.Index + 1) % sum
			}
			break
		}
	}

	return server.Host
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

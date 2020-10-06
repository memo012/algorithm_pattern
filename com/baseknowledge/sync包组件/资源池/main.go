package main

import (
	"fmt"
	"io"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines = 25
	pooledResources = 2
)

// 实现接口类型 资源类型
type dbConnection struct {
	ID int32
}

// 实现接口方法
func (conn *dbConnection) Close() error {
	fmt.Printf("资源关闭,ID:%d\n", conn.ID)
	return nil
}

//给每个连接资源给 id
var idCounter int32

// 创建新资源
func CreateConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	fmt.Printf("创建新资源,id:%d\n", id)
	return &dbConnection{ID: id}, nil
}

// 测试资源池
func performQueries(query int, p *Pool)  {
	conn, err := p.Acquired(10 * time.Second)
	if err != nil {
		fmt.Println("获取资源超时")
		fmt.Println(err)
		return
	}
	// 方法结束后将资源放回资源池
	defer p.Release(conn)

	// 模拟使用资源
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("查询goroutine id:%d,资源ID：%d\n", query, conn.(*dbConnection).ID)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p, err := New(CreateConnection, pooledResources)

	if err != nil {
		fmt.Println(err)
	}

	// 每个goroutine一个查询 每个查询从资源池中获取
	for i := 0; i < maxGoroutines; i++ {
			go func(q int) {
				performQueries(i, p)
				wg.Done()
			}(i)
	}

	wg.Wait()

	fmt.Println("程序结束")

	p.Close()


}

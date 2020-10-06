package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	cond := sync.NewCond(new(sync.Mutex))

	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Println("协程", i, "启动。。。")
			wg.Add(1)
			defer wg.Done()
			cond.L.Lock()
			fmt.Println("协程", i, "加锁。。。")
			cond.Wait()
			fmt.Println("协程", i, "解锁。。。")
			cond.L.Unlock()
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("主协程发送信号量。。。")
	cond.Signal()

	time.Sleep(time.Second)
	fmt.Println("主协程发送信号量。。。")
	cond.Signal()

	time.Sleep(time.Second)
	fmt.Println("主协程发送信号量。。。")
	cond.Signal()

	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	var wg sync.WaitGroup
	onceFunc := func() {
		fmt.Println("法师爱你们哟~")
	}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			once.Do(onceFunc) // 多次调用只执行一次
		}()
	}
	wg.Wait()
}

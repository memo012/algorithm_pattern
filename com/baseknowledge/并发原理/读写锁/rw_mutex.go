package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.RWMutex
	count int32
}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Count() int32 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func main() {
	var count Counter
	for i := 0; i < 10; i++ {
		go func() {
			for {
				ct := count.Count()
				fmt.Println(ct)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	for {
		count.Incr()
		time.Sleep(time.Second)
	}
}

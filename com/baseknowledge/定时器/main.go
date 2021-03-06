package main

import (
	"fmt"
	"time"
)

func timer() {
	time1 := time.NewTimer(time.Second * 2)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)

	t2 := <-time1.C
	fmt.Printf("t1:%v\n", t2)

	timer2 := time.NewTimer(time.Second * 2)
	t2 = <-timer2.C
	fmt.Println("2s后:", t2)

	// 如果只是想单纯的等待的话 可以使用time.Sleep来实现
	time.Sleep(time.Second * 2)
	fmt.Println("再一次2s后", time.Now())

	t2 = <-time.After(time.Second * 2) // time.After函数的返回值是chan Time
	fmt.Println("再再一次2s后:", t2)

	timer3 := time.NewTimer(time.Second)
	go func() {
		t2 = <-timer3.C
		fmt.Println("Timer 3 expired:", t2)
	}()

	stop := timer3.Stop() // 停止定时器

	// 阻止timer事件发生 当该函数执行后 timer计时器停止 相应的事件不再执行
	if stop {
		fmt.Println("Timer 3 stopped")
	}

	fmt.Println("before:", time.Now())
	timer4 := time.NewTimer(time.Second * 5) //原来设置5s
	timer4.Reset(time.Second * 1)            //重新设置时间 即修改NewTimer的时间
	<-timer4.C
	fmt.Println("after")
}
func ticker() {
	tick := time.NewTicker(time.Second)
	go func() {
		for i := range tick.C {
			fmt.Println("current time:", i)
		}
	}()
	time.Sleep(time.Second * 10)
	tick.Stop()
	fmt.Println("Ticker stopped")
}
func main() {
	timer()
	ticker()
}

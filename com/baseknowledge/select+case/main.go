package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan bool, 1)

	//ch <- true

	fmt.Println(time.Now())
	idleDuration := 5 * time.Second
	idleTimeout := time.NewTimer(idleDuration)
	defer idleTimeout.Stop()
	select {
	case t := <-ch:
		fmt.Println(t)
	case t := <-idleTimeout.C:
		fmt.Println(t, time.Now(),"---------------")
	}
	fmt.Println("end over")
	time.Sleep(time.Second * 10)
}

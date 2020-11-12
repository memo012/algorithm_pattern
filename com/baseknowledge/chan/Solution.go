package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
		case v := <-ch:
			fmt.Println(v)
		}
	}

	//go goRoutineB(ch)
	//go goRoutineA(ch)
	//ch <- 3
	//ch <- 3
	//time.Sleep(time.Second)
}

func goRoutineA(a <-chan int) {
	val := <-a
	fmt.Println("goRoutineA:", val)
}

func goRoutineB(b chan int) {
	val := <-b
	fmt.Println("goRoutineB:", val)
}

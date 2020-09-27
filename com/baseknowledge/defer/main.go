package main

import "fmt"

func s() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func m()  {
	for i := 0; i < 3; i++ {
		i := i
		defer func() {
			fmt.Println(i)
		}()
	}
}

func n() {
	for i := 0; i < 3; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func main() {
	s()
	fmt.Println("-------")
	m()
	fmt.Println("-------")
	n()
}

package main

import (
	"algorithm_pattern/com/sorts/tbsort"
	"fmt"
)

func main() {
	var arr []int = []int{4, 6, 2, 76, 13, 97, 24, 567, 0, 4}
	tbsort.BubbleSort(arr)
	fmt.Println(arr)
}

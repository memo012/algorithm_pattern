package main

import (
	"algorithm_pattern/com/sorts/tbsort"
)

func main() {
	var arr []int = []int{4, 6, 2, 76, 13, 97, 24, 567, 0, 4}
	//tbsort.BubbleSort(arr)
	//tbsort.InsertSort(arr)
	//tbsort.SelectSort(arr)
	tbsort.QuickSort(arr)
	//tbsort.HeadSort(arr)
	//tbsort.MergeSort(arr)
}

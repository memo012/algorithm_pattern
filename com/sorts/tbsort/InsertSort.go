package tbsort

import "fmt"

// 插入排序
func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && arr[insertIndex] > insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		arr[insertIndex+1] = insertVal
	}
	fmt.Println(arr)
}

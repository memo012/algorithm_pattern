package tbsort

import "fmt"

// 快排

func QuickSort(arr []int) {
	QuickSorts(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func QuickSorts(arr []int, left int, right int) {
	if left > right {
		return
	}
	i := left
	j := right
	temp := arr[left]
	for i != j {
		for i < j && arr[j] >= temp {
			j--
		}
		for i < j && arr[i] <= temp {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[left] = arr[i]
	arr[i] = temp
	QuickSorts(arr, left, i-1)
	QuickSorts(arr, i+1, right)
}

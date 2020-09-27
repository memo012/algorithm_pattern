package tbsort

import "fmt"

// 归并排序

func MergeSort(arr []int)  {
	slice := make([]int, len(arr))
	MergeSorts(arr, 0, len(arr)-1, slice)
	fmt.Println(arr)
}

func MergeSorts(arr []int, left int, right int, temp []int) {
	if left < right {
		mid := (left + right) / 2
		MergeSorts(arr, left, mid, temp)
		MergeSorts(arr, mid+1, right, temp)
		merge(arr, left, mid, right, temp)
	}
}

func merge(arr []int, left int, mid int, right int, temp []int) {
	i := left
	j := mid + 1
	t := 0
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[t] = arr[i]
			t++
			i++
		} else {
			temp[t] = arr[j]
			t++
			j++
		}
	}

	for i <= mid {
		temp[t] = arr[i]
		t++
		i++
	}
	for j <= right {
		temp[t] = arr[j]
		t++
		j++
	}

	index := 0
	newTemp := left
	for newTemp <= right {
		arr[newTemp] = temp[index]
		newTemp++
		index++
	}
}

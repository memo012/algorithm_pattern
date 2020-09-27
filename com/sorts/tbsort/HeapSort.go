package tbsort

import "fmt"

// 堆排序
func HeadSort(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjustHeap(arr, i, len(arr))
	}
	for i := len(arr) - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		adjustHeap(arr, 0, i)
	}
	fmt.Println(arr)
}

func adjustHeap(arr []int, i, length int) {
	// 先取出当前元素的值 存储到临时变量
	temp := arr[i]
	for j := 2*i + 1; j < length; j = 2*i + 1 {
		// 说明左子节点的值小于右子节点的值
		if j+1 < length && arr[j] < arr[j+1] {
			j++
		}
		// 如果子节点大于父节点
		if arr[j] > temp {
			arr[i] = arr[j]
			i = j
		} else {
			break
		}
	}
	// 循环结束后 我们已经将以i为父节点的树的最大值 放在了最顶(局部)
	// 将temp值放到调整后的位置
	arr[i] = temp
}

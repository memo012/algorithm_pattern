package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	left, top, index := 0, 0, 0
	right, bottom := len(matrix[0])-1, len(matrix)-1
	sum := len(matrix) * len(matrix[0])
	res := make([]int, sum)
	for index < sum {
		for i := left; i <= right; i++ {
			res[index] = matrix[top][i]
			index++
		}
		top++
		if top > bottom {
			break
		}
		for i := top; i <= bottom; i++ {
			res[index] = matrix[i][right]
			index++
		}
		right--
		if right < left {
			break
		}
		for i := right; i >= left; i-- {
			res[index] = matrix[bottom][i]
			index++
		}
		bottom--
		if bottom < top {
			break
		}
		for i := bottom; i >= top; i-- {
			res[index] = matrix[i][left]
			index++
		}
		left++
		if left > right {
			break
		}
	}
	return res
}

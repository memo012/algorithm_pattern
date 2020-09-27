package 调整数组顺序使奇数位于偶数前面

func exchange(nums []int) []int {
	l := 0
	r := len(nums) - 1
	for l < r {
		if nums[l]%2 == 0 && nums[r]%2 != 0 {
			nums[l], nums[r] = nums[r], nums[l]
		}
		for l < len(nums) && nums[l]%2 != 0 {
			l++
		}
		for r >= 0 && nums[r]%2 == 0 {
			r--
		}
	}
	return nums
}

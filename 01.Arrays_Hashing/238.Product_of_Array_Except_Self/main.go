package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	left := 1
	for i := 0; i < n; i++ {
		result[i] = left
		left *= nums[i]
	}

	right := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}
	return result
}

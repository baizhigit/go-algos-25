package twoPointers

func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	left, right := 0, n-1
	pos := n - 1

	for left <= right {
		lv := nums[left] * nums[left]
		rv := nums[right] * nums[right]

		if lv > rv {
			res[pos] = lv
			left++
		} else {
			res[pos] = rv
			right--
		}
		pos--
	}

	return res
}

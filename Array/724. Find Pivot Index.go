package array

func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	left := 0
	for i, num := range nums {
		if left == total-left-num {
			return i
		}
		left += num
	}
	return -1
}

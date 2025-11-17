package array

func subarraySum(nums []int, k int) int {
	sum, count, prefixCount := 0, 0, map[int]int{0: 1}

	for _, num := range nums {
		sum += num

		if freq, ok := prefixCount[sum-k]; ok {
			count += freq
		}

		prefixCount[sum]++
	}

	return count
}

package main

import (
	"slices"
)

// O(n2) time, O(1) space
func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	n := len(nums)
	result := make([][]int, 0)

	for i := 0; i < n-2; i++ {
		// Skip duplicate values for the first number
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1
		target := -nums[i] // Pre-compute for clarity & efficiency

		for left < right {
			sum := nums[left] + nums[right]

			switch {
			case sum == target:
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicates for left and right pointers
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--

			case sum < target:
				left++
			default:
				right--
			}
		}
	}
	return result
}

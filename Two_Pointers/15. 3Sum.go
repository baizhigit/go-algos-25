package twoPointers

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := make([][]int, 0)

	for i := 0; i < n-2; i++ {
		// Skip duplicate values for i
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1
		target := -nums[i]

		for left < right {
			sum := nums[left] + nums[right]

			switch {
			case sum == target:
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				// Skip duplicate values for left
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				// Skip duplicate values for right
				for left < right && nums[right] == nums[right+1] {
					right--
				}

			case sum < target:
				left++

			default:
				right--
			}
		}
	}

	return res
}

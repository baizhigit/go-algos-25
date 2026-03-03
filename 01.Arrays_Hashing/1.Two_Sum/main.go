package main

// O(n) time | O(n) space
func twoSum(nums []int, target int) []int {
	lookup := make(map[int]int)
	for idx, num := range nums {
		diff := target - num
		if foundIdx, found := lookup[diff]; found {
			return []int{foundIdx, idx}
		}
		lookup[num] = idx
	}
	return []int{}
}

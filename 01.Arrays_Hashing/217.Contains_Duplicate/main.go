package main

// O(n) time | O(n) space
func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	lookup := make(map[int]struct{})
	for _, num := range nums {
		if _, found := lookup[num]; found {
			return true
		}
		lookup[num] = struct{}{}
	}
	return false
}

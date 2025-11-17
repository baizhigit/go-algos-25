package twoPointers

func twoSum(numbers []int, target int) []int {
	for left, right := 0, len(numbers)-1; left < right; {
		sum := numbers[left] + numbers[right]

		switch {
		case sum == target:
			return []int{left + 1, right + 1}
		case sum > target:
			right--
		default:
			left++
		}
	}
	return nil
}

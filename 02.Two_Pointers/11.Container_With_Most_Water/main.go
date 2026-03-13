package main

// O(n) time, O(1) space
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		leftHeight := height[left]
		rightHeight := height[right]
		minHeight := min(leftHeight, rightHeight)
		area := minHeight * (right - left)
		maxArea = max(maxArea, area)

		if leftHeight <= rightHeight {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

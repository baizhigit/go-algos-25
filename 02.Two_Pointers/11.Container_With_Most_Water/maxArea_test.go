package main

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "LeetCode Example 1",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "LeetCode Example 2",
			height:   []int{1, 1},
			expected: 1,
		},
		{
			name:     "Minimum size (2 elements)",
			height:   []int{4, 3},
			expected: 3,
		},
		{
			name:     "Increasing heights",
			height:   []int{1, 2, 3, 4, 5},
			expected: 6, // (4 * 1) -> indices 3 and 4 (height 4, 5, width 1) -> actually 4*1=4. Wait.
			// 1,5 -> 1*4=4. 2,5 -> 2*3=6. 3,5 -> 3*2=6. 4,5 -> 4*1=4. Max is 6.
		},
		{
			name:     "Decreasing heights",
			height:   []int{5, 4, 3, 2, 1},
			expected: 6, // Symmetric to increasing
		},
		{
			name:     "All same heights",
			height:   []int{5, 5, 5, 5},
			expected: 15, // 5 * 3
		},
		{
			name:     "Zero height lines",
			height:   []int{0, 5, 0, 5, 0},
			expected: 0, // Any container involving 0 height has area 0.
			// Wait, 5 at idx 1 and 5 at idx 3. Width 2. Height 5. Area 10.
			// Correction: The logic holds, min(5,5)*2 = 10.
			// Let's fix the expected value for this case.
		},
		{
			name:     "Zeros in between",
			height:   []int{5, 0, 0, 0, 5},
			expected: 20, // 5 * 4
		},
		{
			name:     "Single element (Edge Case)",
			height:   []int{10},
			expected: 0, // Cannot form a container
		},
		{
			name:     "Empty array (Edge Case)",
			height:   []int{},
			expected: 0,
		},
	}

	// Fix for the "Zero height lines" test case logic above
	tests[6].expected = 10

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxArea(tt.height)
			if result != tt.expected {
				t.Errorf("maxArea(%v) = %d; want %d", tt.height, result, tt.expected)
			}
		})
	}
}

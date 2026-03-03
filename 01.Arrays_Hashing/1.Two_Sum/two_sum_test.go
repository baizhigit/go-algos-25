package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "basic case - two numbers sum to target",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "numbers not in order",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "same number twice",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, -2, -3, -4, -5},
			target:   -8,
			expected: []int{2, 4},
		},
		{
			name:     "with zero",
			nums:     []int{0, 4, 3, 0},
			target:   0,
			expected: []int{0, 3},
		},
		{
			name:     "large numbers",
			nums:     []int{1000000, 2000000, 3000000},
			target:   5000000,
			expected: []int{1, 2},
		},
		{
			name:     "single element - no solution",
			nums:     []int{1},
			target:   2,
			expected: []int{},
		},
		{
			name:     "empty slice - no solution",
			nums:     []int{},
			target:   5,
			expected: []int{},
		},
		{
			name:     "no valid pair exists",
			nums:     []int{1, 2, 3, 4},
			target:   10,
			expected: []int{},
		},
		{
			name:     "target achieved with later numbers",
			nums:     []int{1, 2, 3, 4, 5},
			target:   9,
			expected: []int{3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.nums, tt.target)

			// Check if the result matches expected
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("twoSum(%v, %d) = %v; want %v",
					tt.nums, tt.target, result, tt.expected)
			}

			// Additional validation: verify that the indices actually sum to target
			if len(result) == 2 {
				sum := tt.nums[result[0]] + tt.nums[result[1]]
				if sum != tt.target {
					t.Errorf("twoSum(%v, %d) returned indices %v that sum to %d, not %d",
						tt.nums, tt.target, result, sum, tt.target)
				}

				// Verify indices are in ascending order
				if result[0] > result[1] {
					t.Errorf("twoSum(%v, %d) returned indices %v - expected ascending order",
						tt.nums, tt.target, result)
				}
			}
		})
	}
}

// Benchmark tests to check performance
func BenchmarkTwoSum(b *testing.B) {
	nums := []int{2, 7, 11, 15, 3, 6, 8, 12, 14, 18, 21, 24, 27, 30}
	target := 41 // 11 + 30

	for i := 0; i < b.N; i++ {
		twoSum(nums, target)
	}
}

// Example test showing how to test with large input
func TestTwoSumLargeInput(t *testing.T) {
	// Create a large slice
	size := 10000
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = i + 1
	}

	// Target is sum of first and last elements
	target := 1 + size // 10001

	result := twoSum(nums, target)

	// Check that we got exactly 2 indices
	if len(result) != 2 {
		t.Errorf("twoSum on large input returned %d indices, expected 2", len(result))
		return
	}

	// Verify the indices actually sum to the target
	sum := nums[result[0]] + nums[result[1]]
	if sum != target {
		t.Errorf("twoSum on large input returned indices %v that sum to %d, want %d",
			result, sum, target)
	}

	// Verify the indices are valid (within range)
	for _, idx := range result {
		if idx < 0 || idx >= size {
			t.Errorf("twoSum returned invalid index %d (slice size %d)", idx, size)
		}
	}

	// Optional: Verify it found the correct pair (0 and size-1)
	// But don't fail if the order is different
	foundCorrectPair := (result[0] == 0 && result[1] == size-1) ||
		(result[0] == size-1 && result[1] == 0)

	if !foundCorrectPair {
		t.Logf("twoSum returned %v, which is a valid pair but not [0, %d] or [%d, 0]",
			result, size-1, size-1)
		// This is informational, not a failure
	}
}

// Test helper function to verify result order
func TestTwoSumOrder(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{3, 2, 4}, 6},       // Should return [1,2] not [2,1]
		{[]int{2, 5, 5, 11}, 10},  // Should return [1,2] not [2,1]
		{[]int{1, 2, 3, 4, 5}, 9}, // Should return [3,4] not [4,3]
	}

	for _, tc := range testCases {
		result := twoSum(tc.nums, tc.target)
		if len(result) == 2 && result[0] > result[1] {
			t.Errorf("twoSum(%v, %d) = %v - indices should be in ascending order",
				tc.nums, tc.target, result)
		}
	}
}

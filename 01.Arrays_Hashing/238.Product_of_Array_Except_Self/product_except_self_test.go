package main

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Basic example 1",
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			name:     "Basic example 2",
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
		{
			name:     "Single element",
			nums:     []int{5},
			expected: []int{1},
		},
		{
			name:     "Two elements",
			nums:     []int{2, 3},
			expected: []int{3, 2},
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		{
			name:     "Single zero",
			nums:     []int{0, 1, 2, 3},
			expected: []int{6, 0, 0, 0},
		},
		{
			name:     "Multiple zeros",
			nums:     []int{0, 0, 1, 2},
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -2, -3},
			expected: []int{6, 3, 2},
		},
		{
			name:     "Mixed positive and negative",
			nums:     []int{-1, 2, -3, 4},
			expected: []int{-24, 12, -8, 6},
		},
		{
			name:     "Large numbers",
			nums:     []int{10, 20, 30},
			expected: []int{600, 300, 200},
		},
		{
			name:     "Contains one and zero",
			nums:     []int{1, 0},
			expected: []int{0, 1},
		},
		{
			name:     "All ones",
			nums:     []int{1, 1, 1, 1},
			expected: []int{1, 1, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := productExceptSelf(tt.nums)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("productExceptSelf(%v) = %v; expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestProductExceptSelfDoesNotModifyInput(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	original := make([]int, len(nums))
	copy(original, nums)

	productExceptSelf(nums)

	if !reflect.DeepEqual(nums, original) {
		t.Errorf("Input array was modified: got %v, expected %v", nums, original)
	}
}

func TestProductExceptSelfOutputLength(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
	}{
		{"Empty array", []int{}},
		{"Single element", []int{5}},
		{"Multiple elements", []int{1, 2, 3, 4, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := productExceptSelf(tc.nums)
			if len(result) != len(tc.nums) {
				t.Errorf("Output length mismatch: got %d, expected %d", len(result), len(tc.nums))
			}
		})
	}
}

// Benchmark test for performance
func BenchmarkProductExceptSelf(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < b.N; i++ {
		productExceptSelf(nums)
	}
}

func BenchmarkProductExceptSelfLarge(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i + 1
	}

	for i := 0; i < b.N; i++ {
		productExceptSelf(nums)
	}
}

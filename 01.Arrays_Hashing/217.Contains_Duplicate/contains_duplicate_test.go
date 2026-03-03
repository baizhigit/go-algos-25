package main

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "basic case - contains duplicate",
			nums:     []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "no duplicates",
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "empty slice",
			nums:     []int{},
			expected: false,
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: false,
		},
		{
			name:     "multiple duplicates",
			nums:     []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
		{
			name:     "all same numbers",
			nums:     []int{5, 5, 5, 5, 5},
			expected: true,
		},
		{
			name:     "negative numbers with duplicates",
			nums:     []int{-1, -2, -3, -1},
			expected: true,
		},
		{
			name:     "negative numbers without duplicates",
			nums:     []int{-1, -2, -3, -4},
			expected: false,
		},
		{
			name:     "zero and positive numbers with duplicates",
			nums:     []int{0, 1, 2, 0, 3},
			expected: true,
		},
		{
			name:     "large numbers with duplicates",
			nums:     []int{1000000, 2000000, 1000000, 3000000},
			expected: true,
		},
		{
			name:     "duplicate at the end",
			nums:     []int{1, 2, 3, 4, 5, 1},
			expected: true,
		},
		{
			name:     "duplicate at the beginning",
			nums:     []int{5, 1, 2, 3, 4, 5},
			expected: true,
		},
		{
			name:     "two elements with duplicate",
			nums:     []int{1, 1},
			expected: true,
		},
		{
			name:     "two elements without duplicate",
			nums:     []int{1, 2},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := containsDuplicate(tt.nums)
			if result != tt.expected {
				t.Errorf("containsDuplicate(%v) = %v; want %v",
					tt.nums, result, tt.expected)
			}
		})
	}
}

// Benchmark tests to check performance
func BenchmarkContainsDuplicate(b *testing.B) {
	// Test with a slice that has no duplicates (worst case - must check all elements)
	nums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		nums[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		containsDuplicate(nums)
	}
}

func BenchmarkContainsDuplicateWithEarlyExit(b *testing.B) {
	// Test with a slice that has a duplicate early (best case)
	nums := make([]int, 10000)
	for i := 1; i < 10000; i++ {
		nums[i] = i
	}
	nums[5000] = 0 // create duplicate at position 5000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		containsDuplicate(nums)
	}
}

// Test with very large input to ensure no performance issues
func TestContainsDuplicateLargeInput(t *testing.T) {
	// Create a large slice without duplicates
	size := 100000
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = i
	}

	// Should return false (no duplicates)
	result := containsDuplicate(nums)
	if result != false {
		t.Errorf("containsDuplicate on large slice without duplicates = %v; want false", result)
	}

	// Add a duplicate at the end
	nums = append(nums, 42)
	result = containsDuplicate(nums)
	if result != true {
		t.Errorf("containsDuplicate on large slice with duplicate = %v; want true", result)
	}
}

// Test with mixed types (through the lens of ints)
func TestContainsDuplicateEdgeCases(t *testing.T) {
	// Test with maximum int values
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1

	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "max int values with duplicate",
			nums:     []int{maxInt, maxInt - 1, maxInt},
			expected: true,
		},
		{
			name:     "min int values with duplicate",
			nums:     []int{minInt, minInt + 1, minInt},
			expected: true,
		},
		{
			name:     "mixed min and max without duplicate",
			nums:     []int{minInt, 0, maxInt},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := containsDuplicate(tt.nums)
			if result != tt.expected {
				t.Errorf("containsDuplicate(%v) = %v; want %v",
					tt.nums, result, tt.expected)
			}
		})
	}
}

// Test helper to verify the function doesn't modify the input slice
func TestContainsDuplicateImmutability(t *testing.T) {
	original := []int{3, 1, 4, 1, 5}
	originalCopy := make([]int, len(original))
	copy(originalCopy, original)

	_ = containsDuplicate(original)

	// Verify original slice wasn't modified
	for i := range original {
		if original[i] != originalCopy[i] {
			t.Errorf("containsDuplicate modified the input slice at index %d: original %v, now %v",
				i, originalCopy, original)
			break
		}
	}
}

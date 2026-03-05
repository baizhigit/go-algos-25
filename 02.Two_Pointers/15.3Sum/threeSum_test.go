package main

import (
	"slices"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{"empty", []int{}, [][]int{}},
		{"single element", []int{1}, [][]int{}},
		{"no triplet", []int{1, 2, -2}, [][]int{}},
		{"basic triplet", []int{-1, 0, 1}, [][]int{{-1, 0, 1}}},
		{"LeetCode example", []int{-1, 0, 1, 2, -1, -4},
			[][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"all zeros", []int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{"duplicates", []int{-2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := threeSum(tt.input)
			if !equal2D(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// Helper: compare 2D slices ignoring order
func equal2D(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	// Sort both for order-insensitive comparison
	sort2D := func(s [][]int) {
		slices.SortFunc(s, func(a, b []int) int {
			for i := 0; i < 3; i++ {
				if a[i] != b[i] {
					return a[i] - b[i]
				}
			}
			return 0
		})
	}
	sort2D(a)
	sort2D(b)
	for i := range a {
		if !slices.Equal(a[i], b[i]) {
			return false
		}
	}
	return true
}

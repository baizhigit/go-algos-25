package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		// Edge cases
		{"empty string", "", true},
		{"single character", "a", true},
		{"two same chars", "aa", true},
		{"two different chars", "ab", false},

		// Basic palindromes
		{"simple palindrome", "racecar", true},
		{"not palindrome", "hello", false},

		// Case insensitivity
		{"mixed case palindrome", "RaceCar", true},
		{"mixed case not palindrome", "RaceCat", false},

		// With spaces
		{"palindrome with spaces", "a b a", true},
		{"not palindrome with spaces", "a b c", false},

		// With punctuation
		{"LeetCode example 1", "A man, a plan, a canal: Panama", true},
		{"LeetCode example 2", "race a car", false},
		{"No x in Nixon", "No 'x' in Nixon", true},

		// Only non-alphanumeric
		{"all special chars", "!!!", true},
		{"all spaces", "   ", true},

		// Numbers
		{"numeric palindrome", "12321", true},
		{"alphanumeric palindrome", "a1b2b1a", true},
		{"alphanumeric not palindrome", "0P", false},

		// Mixed edge cases
		{"single alphanumeric in specials", "!!a!!", true},
		{"two alphanumeric in specials", "!!ab!!", false},
		{"palindrome with mixed special chars", "A.B,C:BA", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("isPalindrome(%q) = %v; expected %v",
					tt.input, result, tt.expected)
			}
		})
	}
}

// Benchmark for performance validation
func BenchmarkIsPalindrome(b *testing.B) {
	input := "A man, a plan, a canal: Panama"
	for i := 0; i < b.N; i++ {
		isPalindrome(input)
	}
}

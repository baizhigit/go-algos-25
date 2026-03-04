package main

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		// Basic anagram cases
		{"valid anagram 1", "anagram", "nagaram", true},
		{"valid anagram 2", "rat", "tar", true},
		{"valid anagram 3", "listen", "silent", true},
		{"valid anagram 4", "abc", "cba", true},

		// Non-anagram cases
		{"not anagram 1", "rat", "car", false},
		{"not anagram 2", "hello", "world", false},
		{"not anagram 3", "abc", "def", false},

		// Length mismatch cases
		{"different length 1", "ab", "abc", false},
		{"different length 2", "a", "", false},
		{"different length 3", "abc", "abcd", false},

		// Edge cases
		{"empty strings", "", "", true},
		{"single char same", "a", "a", true},
		{"single char different", "a", "b", false},

		// Character frequency cases
		{"same chars different freq", "aabb", "abbb", false},
		{"all same characters", "aaa", "aaa", true},
		{"all same chars different count", "aaa", "aa", false},

		// Mixed cases
		{"complex anagram", "algorithm", "logarithm", true},
		{"complex not anagram", "algorithm", "logarithms", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isAnagram(tt.s, tt.t)
			if result != tt.expected {
				t.Errorf("isAnagram(%q, %q) = %v; expected %v",
					tt.s, tt.t, result, tt.expected)
			}
		})
	}
}

// Benchmark test to measure performance
func BenchmarkIsAnagram(b *testing.B) {
	s := "anagram"
	t := "nagaram"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isAnagram(s, t)
	}
}

// Benchmark with longer strings
func BenchmarkIsAnagramLong(b *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyz"
	t := "zyxwvutsrqponmlkjihgfedcba"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isAnagram(s, t)
	}
}

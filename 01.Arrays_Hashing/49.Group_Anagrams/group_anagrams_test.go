package main

import (
	"sort"
	"testing"
)

// compareAnagramGroups compares two slices of anagram groups
// Order of groups and order within groups doesn't matter
func compareAnagramGroups(got, want [][]string) bool {
	if len(got) != len(want) {
		return false
	}

	// Normalize both: sort each group internally, then sort all groups
	normalize := func(groups [][]string) [][]string {
		normalized := make([][]string, len(groups))
		for i, group := range groups {
			normalized[i] = make([]string, len(group))
			copy(normalized[i], group)
			sort.Strings(normalized[i])
		}
		sort.Slice(normalized, func(i, j int) bool {
			if len(normalized[i]) != len(normalized[j]) {
				return len(normalized[i]) < len(normalized[j])
			}
			for k := 0; k < len(normalized[i]); k++ {
				if normalized[i][k] != normalized[j][k] {
					return normalized[i][k] < normalized[j][k]
				}
			}
			return false
		})
		return normalized
	}

	gotNorm := normalize(got)
	wantNorm := normalize(want)

	for i := range gotNorm {
		if len(gotNorm[i]) != len(wantNorm[i]) {
			return false
		}
		for j := range gotNorm[i] {
			if gotNorm[i][j] != wantNorm[i][j] {
				return false
			}
		}
	}

	return true
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			name:     "example from leetcode",
			input:    []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			name:     "single string",
			input:    []string{"hello"},
			expected: [][]string{{"hello"}},
		},
		{
			name:     "empty input",
			input:    []string{},
			expected: [][]string{},
		},
		{
			name:     "empty string",
			input:    []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "all same anagrams",
			input:    []string{"abc", "bca", "cab", "acb"},
			expected: [][]string{{"abc", "bca", "cab", "acb"}},
		},
		{
			name:     "no anagrams",
			input:    []string{"abc", "def", "ghi"},
			expected: [][]string{{"abc"}, {"def"}, {"ghi"}},
		},
		{
			name:     "single character strings",
			input:    []string{"a", "b", "a", "c", "b"},
			expected: [][]string{{"a", "a"}, {"b", "b"}, {"c"}},
		},
		{
			name:     "different lengths",
			input:    []string{"ab", "ba", "abc", "bca", "abcd"},
			expected: [][]string{{"ab", "ba"}, {"abc", "bca"}, {"abcd"}},
		},
		{
			name:     "collision test - freq key separator",
			input:    []string{"ab", "aab"}, // Tests that "1#1#" != "2#1#"
			expected: [][]string{{"ab"}, {"aab"}},
		},
		{
			name:     "repeated anagrams",
			input:    []string{"test", "test", "test", "sett"},
			expected: [][]string{{"test", "test", "test", "sett"}},
		},
		{
			name:     "different words same chars",
			input:    []string{"abc", "abd", "bca"},
			expected: [][]string{{"abc", "bca"}, {"abd"}},
		},
		{
			name:     "longer strings",
			input:    []string{"listen", "silent", "enlist", "hello", "world"},
			expected: [][]string{{"listen", "silent", "enlist"}, {"hello"}, {"world"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := groupAnagrams(tt.input)
			if !compareAnagramGroups(result, tt.expected) {
				t.Errorf("groupAnagrams(%v)\n got: %v\nwant: %v",
					tt.input, result, tt.expected)
			}
		})
	}
}

// Test that freqKey produces consistent keys for anagrams
func TestFreqKey(t *testing.T) {
	tests := []struct {
		str1    string
		str2    string
		sameKey bool
	}{
		{"eat", "tea", true},
		{"abc", "bca", true},
		{"listen", "silent", true},
		{"abc", "def", false},
		{"ab", "aab", false},
		{"", "", true},
		{"a", "a", true},
		{"a", "b", false},
	}

	for _, tt := range tests {
		t.Run(tt.str1+"_"+tt.str2, func(t *testing.T) {
			key1 := freqKey(tt.str1)
			key2 := freqKey(tt.str2)
			if tt.sameKey && key1 != key2 {
				t.Errorf("freqKey(%q) = %q, freqKey(%q) = %q; expected same key",
					tt.str1, key1, tt.str2, key2)
			}
			if !tt.sameKey && key1 == key2 {
				t.Errorf("freqKey(%q) = %q, freqKey(%q) = %q; expected different keys",
					tt.str1, key1, tt.str2, key2)
			}
		})
	}
}

// Test edge case: ensure separator prevents collisions
func TestFreqKeyNoCollisions(t *testing.T) {
	// These should produce different keys despite similar frequency patterns
	key1 := freqKey("ab")  // a:1, b:1
	key2 := freqKey("aab") // a:2, b:1

	if key1 == key2 {
		t.Errorf("freqKey collision detected: %q == %q", key1, key2)
	}
}

// Benchmark tests
func BenchmarkGroupAnagrams(b *testing.B) {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		groupAnagrams(input)
	}
}

func BenchmarkGroupAnagramsLarge(b *testing.B) {
	input := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		input[i] = "abcdefghijklmnopqrstuvwxyz"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		groupAnagrams(input)
	}
}

func BenchmarkGroupAnagramsManyGroups(b *testing.B) {
	input := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		input[i] = string([]byte{byte('a' + i%26)})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		groupAnagrams(input)
	}
}

func BenchmarkFreqKey(b *testing.B) {
	str := "anagram"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		freqKey(str)
	}
}

package main

import (
	"bytes"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "hello world",
			input:    []byte("hello"),
			expected: []byte("olleh"),
		},
		{
			name:     "empty slice",
			input:    []byte(""),
			expected: []byte(""),
		},
		{
			name:     "single character",
			input:    []byte("a"),
			expected: []byte("a"),
		},
		{
			name:     "two characters",
			input:    []byte("ab"),
			expected: []byte("ba"),
		},
		{
			name:     "odd length",
			input:    []byte("abcde"),
			expected: []byte("edcba"),
		},
		{
			name:     "even length",
			input:    []byte("abcdef"),
			expected: []byte("fedcba"),
		},
		{
			name:     "all same characters",
			input:    []byte("aaaa"),
			expected: []byte("aaaa"),
		},
		{
			name:     "palindrome",
			input:    []byte("racecar"),
			expected: []byte("racecar"),
		},
		{
			name:     "with spaces",
			input:    []byte("h e l l o"),
			expected: []byte("o l l e h"),
		},
		{
			name:     "with numbers",
			input:    []byte("12345"),
			expected: []byte("54321"),
		},
		{
			name:     "mixed alphanumeric",
			input:    []byte("abc123"),
			expected: []byte("321cba"),
		},
		{
			name:     "special characters",
			input:    []byte("!@#$%"),
			expected: []byte("%$#@!"),
		},
		{
			name:     "unicode bytes (UTF-8)",
			input:    []byte("café"), // Note: 'é' is 2 bytes in UTF-8
			expected: []byte("éfac"), // Reversed by byte, not rune
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy to avoid modifying the test input
			inputCopy := make([]byte, len(tt.input))
			copy(inputCopy, tt.input)

			reverseString(inputCopy)

			if !bytes.Equal(inputCopy, tt.expected) {
				t.Errorf("reverseString(%q) = %q; expected %q",
					tt.input, inputCopy, tt.expected)
			}
		})
	}
}

// Test that the function modifies in-place (no new allocation)
func TestReverseStringInPlace(t *testing.T) {
	input := []byte("testing")
	originalPtr := &input[0]

	reverseString(input)

	// Verify the same underlying array was modified
	newPtr := &input[0]
	if originalPtr != newPtr {
		t.Error("reverseString did not modify in-place")
	}

	expected := []byte("gnitset")
	if !bytes.Equal(input, expected) {
		t.Errorf("Expected %q, got %q", expected, input)
	}
}

// Test that input is actually modified (not read-only)
func TestReverseStringModifiesInput(t *testing.T) {
	input := []byte("before")
	reverseString(input)

	if string(input) == "before" {
		t.Error("reverseString did not modify the input slice")
	}

	if string(input) != "erofeb" {
		t.Errorf("Expected 'erofeb', got %q", input)
	}
}

// Benchmark tests
func BenchmarkReverseString(b *testing.B) {
	input := []byte("hello world")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		inputCopy := make([]byte, len(input))
		copy(inputCopy, input)
		reverseString(inputCopy)
	}
}

func BenchmarkReverseStringShort(b *testing.B) {
	input := []byte("ab")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		inputCopy := make([]byte, len(input))
		copy(inputCopy, input)
		reverseString(inputCopy)
	}
}

func BenchmarkReverseStringLong(b *testing.B) {
	input := make([]byte, 10000)
	for i := range input {
		input[i] = byte('a' + (i % 26))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		inputCopy := make([]byte, len(input))
		copy(inputCopy, input)
		reverseString(inputCopy)
	}
}

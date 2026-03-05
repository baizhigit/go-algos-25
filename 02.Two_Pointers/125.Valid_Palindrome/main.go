package main

func isAlphanum(c byte) bool { // unicode.IsNumber / unicode.IsLetter
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9')
}

func toLower(c byte) byte { // unicode.ToLower
	if 'A' <= c && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

// O(n) time, O(1) space
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if !isAlphanum(s[left]) {
			left++
		} else if !isAlphanum(s[right]) {
			right--
		} else if toLower(s[left]) == toLower(s[right]) {
			left++
			right--
		} else {
			return false
		}
	}

	return true
}

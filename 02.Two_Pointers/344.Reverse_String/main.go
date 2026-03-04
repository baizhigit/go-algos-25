package main

// O(n) time | O(1) space
func reverseString(s []byte) {
	left, right := 0, len(s)-1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func reverseStringByRune(s []byte) {
	runes := []rune(string(s))
	left, right := 0, len(runes)-1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	copy(s, []byte(string(runes)))
}

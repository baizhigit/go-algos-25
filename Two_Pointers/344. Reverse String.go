package twoPointers

func reverseString(s []byte) {
	for left, right := 0, len(s)-1; left < right; {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

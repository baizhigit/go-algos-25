package main

// O(n) time | O(1) space
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var freqs [26]int
	for idx := range s {
		freqs[s[idx]-'a']++
		freqs[t[idx]-'a']--
	}

	for _, freq := range freqs {
		if freq != 0 {
			return false
		}
	}

	return true
}

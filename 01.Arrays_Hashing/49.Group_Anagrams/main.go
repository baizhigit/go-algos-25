package main

import (
	"fmt"
	"strconv"
	"strings"
)

func freqKey(str string) string {
	var freqs [26]int

	for _, symbol := range str {
		freqs[symbol-'a']++
	}

	var sb strings.Builder

	for _, freq := range freqs {
		sb.WriteByte('#')
		sb.WriteString(strconv.Itoa(freq))
	}

	return sb.String()
}

// O(n*k) time | O(n*k) space
func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, str := range strs {
		key := freqKey(str)
		groups[key] = append(groups[key], str)
	}
	fmt.Println(groups)

	result := make([][]string, 0, len(groups))

	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

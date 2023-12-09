package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {

	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}

	m1 := make(map[string]int)
	m2 := make(map[string]int)

	for i := 0; i < len(pattern); i++ {
		index1, ok1 := m1[string(pattern[i])]

		index2, ok2 := m2[words[i]]
		if ok1 != ok2 || index1 != index2 {
			return false
		}

		if !ok1 {
			m1[string(pattern[i])] = i
		}
		if !ok2 {
			m2[words[i]] = i
		}
	}

	return true
}
func main() {
	//pattern := "abba"
	//s := "dog cat cat dog"
	//pattern := "abba"
	//s := "dog cat cat fish"
	pattern := "aaaa"
	s := "dog cat cat dog"
	//pattern := "abab"
	//s := "dog cat cat dog"

	isValid := wordPattern(pattern, s)

	fmt.Printf("isValid: %v\n", isValid)
}

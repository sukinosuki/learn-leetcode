package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	m := make(map[string]string)
	m2 := make(map[string]string)
	wordStack := []string{}
	patternStack := strings.Split(pattern, "")

	for index, char := range patternStack {
		m[char] = words[index]
		m2[words[index]] = m[char]
		wordStack = append(wordStack, words[index])
	}
	if len(m) != len(m2) {
		return false
	}

	for i := 0; i < len(patternStack); i++ {
		char := patternStack[i]
		if m[char] != wordStack[i] {
			return false
		}
	}

	return true
}

func main() {
	pattern := "abba"
	s := "dog cat cat dog"
	//pattern := "abba"
	//s := "dog cat cat fish"
	//pattern := "aaaa"
	//s := "dog cat cat dog"
	//pattern := "abab"
	//s := "dog cat cat dog"

	isValid := wordPattern(pattern, s)

	fmt.Printf("isValid: %v\n", isValid)
}

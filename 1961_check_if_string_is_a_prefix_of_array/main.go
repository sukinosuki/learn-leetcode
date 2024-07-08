package main

import "fmt"

func isPrefixString(s string, words []string) bool {

	n := len(s)

	l := 0
	for i := range words {

		for j := range words[i] {
			if l == n {
				return false
			}
			if words[i][j] != s[l] {
				return false
			}
			l++
		}

		if l == n {
			return true
		}
	}

	return false
}

func main() {
	// true
	//s := "iloveleetcode"
	//words := []string{"i", "love", "leetcode", "apples"}

	// false
	//s := "a"
	//words := []string{"aa", "aaaa", "banana"}
	s := "ccccccccc"
	words := []string{"c", "cc"}
	res := isPrefixString(s, words)

	fmt.Printf("res: %v\n", res)
}

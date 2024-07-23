package main

import "fmt"

func countPrefixes(words []string, s string) int {

	ans := 0
	n := len(s)
	for _, word := range words {

		m := len(word)
		if m > n {
			continue
		}

		if word == s[:m] {
			ans++
		}
	}

	return ans
}

func main() {
	words := []string{"a", "b", "c", "ab", "bc", "abc"}
	s := "abc"
	res := countPrefixes(words, s)

	fmt.Printf("res: %v\n", res)
}

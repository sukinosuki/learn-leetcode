package main

import "fmt"

func firstPalindrome(words []string) string {

	for _, word := range words {
		l := 0
		r := len(word) - 1

		for l < r {
			if word[l] != word[r] {
				break
			}
			l++
			r--
		}

		if l == r || r == l-1 {
			return word
		}

	}

	return ""
}

func main() {
	words := []string{"abc", "car", "ada", "racecar", "cool"}
	res := firstPalindrome(words)

	fmt.Printf("res: %v\n", res)
}

package main

import "fmt"

func longestPalindrome(s string) int {

	count := len(s)

	m := make(map[int32]int)
	for _, char := range s {
		m[char]++
	}
	singleCount := 0
	for _, value := range m {
		if value%2 == 1 {
			singleCount++
			count--
		}
	}
	if singleCount > 0 {
		count++
	}

	return count
}

func main() {
	//s := "abccccdd"
	s := "aaaAaaaa"
	count := longestPalindrome(s)

	fmt.Printf("count: %v\n", count)
}

package main

import "fmt"

func longestPalindrome(s string) int {
	length := len(s)

	arr := make([]int, 52)
	for _, char := range s {
		if 'a' <= char && char <= 'z' {
			arr[char-'a']++
		} else {
			arr[char-'A'+26]++
		}
	}
	for _, value := range arr {
		if value%2 == 1 {
			length--
		}
	}

	if length < len(s) {
		length++
	}

	return length
}

func main() {
	//s := "abccccdd"
	s := "aaaAaaaa"
	count := longestPalindrome(s)

	fmt.Printf("count: %v\n", count)
}

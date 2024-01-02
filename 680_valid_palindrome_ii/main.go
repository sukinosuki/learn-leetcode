package main

import (
	"fmt"
)

func validPalindrome(s string) bool {

	n := 1
	return isPalindrome(s, n)
}

func isPalindrome(s string, n int) bool {
	l1 := 0
	l2 := len(s) - 1

	for l1 < l2 {
		if s[l1] != s[l2] {
			if n == 0 {
				return false
			}

			return isPalindrome(s[l1:l2], n-1) || isPalindrome(s[l1+1:l2+1], n-1)
		}

		l1++
		l2--
	}
	return true
}

func main() {
	s := "abc"
	//s := "abca"
	//s := "abcdcbab"
	//s := "deeee"
	//s := "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"
	res := validPalindrome(s)

	fmt.Printf("res: %v\n", res)
}

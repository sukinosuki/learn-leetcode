package main

import "fmt"

func longestPalindrome(s string) string {

	if len(s) < 2 {
		return s
	}
	n := len(s)
	ans := ""
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if j-i+1 > len(ans) && isPalindrome(s[i:j]) {
				if j-i+1 > len(ans) {
					ans = s[i:j]
				}
			}
		}
	}
	return ans
}

func isPalindrome(s string) bool {

	n := len(s)

	l := 0
	r := n - 1

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	//bab
	s := "babad"
	// a
	//s := "a"

	// bb
	//s := "bb"

	// a
	//s := "ac"
	res := longestPalindrome(s)

	fmt.Printf("res: %v\n", res)
}

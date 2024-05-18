package main

import "fmt"

func removePalindromeSub(s string) int {

	l := 0
	r := len(s)

	count := 0
	for l < r {

		//sub1:=  s[l:r]
		isPalindrome := check(s[l:r])
		if isPalindrome {
			l = r
			r = len(s)
			count++
		} else {
			r--
		}
	}

	return count
}

func check(s string) bool {

	l := 0
	r := len(s)
	for l < r {
		if s[l] != s[r-1] {
			return false
		}

		l++
		r--
	}

	return true

}

func main() {
	// 1
	//s := "ababa"

	// 2
	//s := "abb"

	// 2
	//s := "baabb"
	s := "bbaabaaa"
	res := removePalindromeSub(s)

	fmt.Printf("res: %v\n", res)
}

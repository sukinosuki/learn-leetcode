package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {

	l := 0
	r := len(s) - 1

	for l <= r {

		flag := false
		if s[l] < 48 || s[l] > 122 || (s[l] > 57 && s[l] < 65) || (s[l] > 90 && s[l] < 97) {
			flag = true
			l++
		}
		if s[r] < 48 || s[r] > 122 || (s[r] > 57 && s[r] < 65) || (s[r] > 90 && s[r] < 97) {
			flag = true
			r--
		}
		if flag {
			continue
		}
		if strings.ToLower(string(s[l])) != strings.ToLower(string(s[r])) {
			return false
		}

		l++
		r--
	}

	return true
}

func main() {

	s := "A man, a plan, a canal: Panama"

	isValid := isPalindrome(s)
	fmt.Printf("isValid: %v\n", isValid)
	//s := "0123456789azAZ"
	//for _, char := range s {
	//	fmt.Println(char)
	//}

	//fmt.Println(s[0])
	//fmt.Println(s[1])
	//fmt.Println(s[2])
	//fmt.Println(s[3])
	//fmt.Println(strings.ToLower(string(s[0])))
	//fmt.Println(strings.ToLower(string(s[1])))
	//fmt.Println(s[0] == s[1])
}

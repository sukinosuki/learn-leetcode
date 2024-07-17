package main

import "fmt"

func checkString(s string) bool {

	l := 0

	n := len(s)

	for l < n && s[l] == 'a' {
		l++
	}

	for l < n {
		if s[l] == 'a' {
			return false
		}

		l++
	}

	return true
}

func main() {
	s := "aaabbb"
	res := checkString(s)

	fmt.Printf("res: %v\n", res)
}

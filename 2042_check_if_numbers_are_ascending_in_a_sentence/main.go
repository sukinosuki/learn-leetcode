package main

import (
	"fmt"
	"strconv"
)

func areNumbersAscending(s string) bool {

	l := 0
	n := len(s)

	prevNum := -1
	for l < n {

		for l < n && (s[l]-'0' < 0 || s[l]-'0' > 9) {
			l++
		}

		if l == n {
			return true
		}

		l2 := l + 1
		for l2 < n && (s[l2]-'0' >= 0 && s[l2]-'0' <= 9) {
			l2++
		}
		num, _ := strconv.Atoi(s[l:l2])
		if prevNum != -1 && num <= prevNum {
			return false
		}
		prevNum = num

		l = l2

	}
	return true
}

func main() {
	// true
	//s := "1 box has 3 blue 4 red 6 green and 12 yellow marbles"
	// false
	//s := "hello world 5 x 5"
	s := "sunset is at 7 51 pm overnight lows will be in the low 50 and 60 s"
	res := areNumbersAscending(s)

	fmt.Printf("res: %v\n", res)
}

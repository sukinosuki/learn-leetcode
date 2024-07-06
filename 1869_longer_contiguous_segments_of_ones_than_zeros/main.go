package main

import "fmt"

func checkZeroOnes(s string) bool {

	l := 0

	n := len(s)
	maxOneLength := 0
	maxZeroLength := 0

	for l < n {
		l2 := l
		for l2 < n && s[l2] == '1' {
			l2++
		}

		maxOneLength = max(maxOneLength, l2-l)

		l = l2
		for l2 < n && s[l2] == '0' {
			l2++
		}

		maxZeroLength = max(maxZeroLength, l2-l)

		l = l2
	}

	fmt.Printf("one: %v, zero: %v \n", maxOneLength, maxZeroLength)

	return maxOneLength > maxZeroLength
}

func main() {
	// true
	//s := "1101"
	// true
	s := "111000"
	// 3 2 true
	//s := "0111010011"
	res := checkZeroOnes(s)

	fmt.Printf("res: %v\n", res)
}

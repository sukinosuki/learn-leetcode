package main

import "fmt"

func maxScore(s string) int {

	ans := 0

	for i := 1; i <= len(s)-1; i++ {

		cnt := count(s[:i], '0') + count(s[i:], '1')
		if cnt > ans {
			ans = cnt
		}
	}

	return ans
}

func count(s string, c uint8) int {
	cnt := 0
	for i := range s {
		if s[i] == c {
			cnt++
		}
	}
	return cnt
}

func main() {
	// 5
	//s := "011101"

	// 3
	s := "1111"

	// 1
	//s := "00"
	res := maxScore(s)

	fmt.Printf("res: %v\n", res)
}

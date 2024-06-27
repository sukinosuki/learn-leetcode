package main

import "fmt"

func minOperations(s string) int {

	ans := 0
	prev := s[0]

	for i := 1; i < len(s); i++ {

		if s[i] == prev {
			ans++
			if prev == '1' {
				prev = '0'
			} else {
				prev = '1'
			}
		} else {
			prev = s[i]
		}
	}

	return min(ans, len(s)-ans)
}

func main() {
	// 1
	//s := "0100"
	// 0
	//s := "10"
	s := "1111"
	res := minOperations(s)

	fmt.Printf("res: %v\n", res)
}

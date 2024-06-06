package main

import "fmt"

func makeGood(s string) string {

	var stack []uint8

	for i := range s {

		n := len(stack)
		if n == 0 {
			stack = append(stack, s[i])
			continue
		}

		if s[i]-32 == stack[n-1] || s[i]+32 == stack[n-1] {
			stack = stack[:n-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}

func main() {

	// leetcode
	//s := "leEeetcode"
	s := "abBAcC"

	res := makeGood(s)

	fmt.Printf("res: %v\n", res)
}

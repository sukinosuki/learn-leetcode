package main

import "fmt"

func maxDepth(s string) int {

	var stack []byte

	ans := 0
	for i := range s {

		if s[i] == '(' {
			stack = append(stack, s[i])
			if len(stack) > ans {
				ans = len(stack)
			}
			continue
		}
		if s[i] == ')' {
			stack = stack[:len(stack)-1]
		}
	}

	return ans
}

func main() {
	s := "(1+(2*3)+((8)/4))+1"
	res := maxDepth(s)

	fmt.Printf("res: %v\n", res)
}

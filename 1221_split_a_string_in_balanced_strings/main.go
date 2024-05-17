package main

import "fmt"

func balancedStringSplit(s string) int {

	stack := []uint8{s[0]}
	ans := 0
	for i := 1; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] != s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}

		if len(stack) == 0 {
			ans++
		}
	}

	return ans
}

func main() {
	// 4
	//s := "RLRRLLRLRL"

	// 2
	//s := "RLRRRLLRLL"

	s := "LLLLRRRR"
	res := balancedStringSplit(s)

	fmt.Printf("res: %v\n", res)
}

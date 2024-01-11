package main

import "fmt"

func shortestToChar(s string, c byte) []int {

	m := make(map[int]bool)
	var stack []int

	for i, v := range s {
		if byte(v) == c {
			stack = append(stack, i)
			m[i] = true
		}
	}

	ans := make([]int, len(s))
	loopLevel := 0
	for len(stack) > 0 {
		arr := stack[:]
		stack = nil

		for _, index := range arr {
			if !m[index] {
				m[index] = true
				ans[index] = loopLevel
			}
			if !m[index+1] && index+1 < len(s) {
				stack = append(stack, index+1)
			}
			if !m[index-1] && index-1 >= 0 {
				stack = append(stack, index-1)
			}
		}

		loopLevel++
	}

	return ans
}

func main() {
	s := "loveleetcode"
	c := 'e'
	res := shortestToChar(s, byte(c))

	fmt.Printf("res: %v\n", res)
}

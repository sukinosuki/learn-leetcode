package main

import "fmt"

func removeOuterParentheses(s string) string {

	stack := []uint8{}

	res := []byte{}
	l := 0

	for i := 0; i < len(s); i++ {

		if s[i] == '(' {
			stack = append(stack, '(')
		} else {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			for j := l + 1; j < i; j++ {
				res = append(res, s[j])
			}
			l = i + 1
		}
	}

	res2 := string(res)

	return res2
}

func main() {
	//s := "(()())(())" // ()()()
	//s := "(()())(())(()(()))" // ()()()()(())
	s := "()()" // ()()()()(())
	res := removeOuterParentheses(s)
	fmt.Printf("res: %s\n", res)
}

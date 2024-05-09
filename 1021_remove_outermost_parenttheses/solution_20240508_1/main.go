package main

import "fmt"

func removeOuterParentheses(s string) string {
	var ans, st []rune
	for _, c := range s {
		if c == ')' {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans = append(ans, c)
		}
		if c == '(' {
			st = append(st, c)
		}
	}
	return string(ans)
}

func main() {
	s := "(()())(())" // ()()()
	//s := "(()())(())(()(()))" // ()()()()(())
	//s := "()()" // ()()()()(())
	res := removeOuterParentheses(s)
	fmt.Printf("res: %s\n", res)
}

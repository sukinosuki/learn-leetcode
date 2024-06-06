package main

import "fmt"

func backspaceCompare(s string, t string) bool {

	n1 := len(s)
	n2 := len(t)
	r1 := n1 - 1
	r2 := n2 - 1

	for r1 >= 0 && r2 >= 0 {

		for r1 >= 0 && s[r1] == '#' {
			r1 -= 2
		}
		for r2 >= 0 && t[r2] == '#' {
			r2 -= 2
		}

		if r1 >= 0 && r2 >= 0 {
			if s[r1] != t[r2] {
				return false
			}
		}

		r1--
		r2--
	}

	return true
}

func main() {

	//s := "ab#c"
	//t := "ad#c"

	s := "ab##"
	t := "c#d#"
	res := backspaceCompare(s, t)

	fmt.Printf("res: %v\n", res)
}

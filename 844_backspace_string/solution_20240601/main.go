package main

import "fmt"

func backspaceCompare(s string, t string) bool {

	st1 := print(s)
	st2 := print(t)

	return string(st1) == string(st2)
}

func print(s string) []uint8 {
	var st1 []uint8

	for i := range s {
		if s[i] == '#' {
			if len(st1) > 0 {
				st1 = st1[:len(st1)-1]
			}
		} else {
			st1 = append(st1, s[i])
		}
	}
	return st1
}

func main() {
	s := "a##c"
	t := "#a#c"

	res := backspaceCompare(s, t)

	fmt.Printf("res: %v\n", res)
}

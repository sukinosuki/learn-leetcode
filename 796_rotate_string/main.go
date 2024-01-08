package main

import "fmt"

func rotateString(s string, goal string) bool {
	if s == goal {
		return true
	}
	if len(s) != len(goal) {
		return false
	}
	n := len(s)

	s1L1 := 0
	s2L1 := 0

	s2L2start := 1
	s2L2 := s2L2start
	for s1L1 < n {

		if s2L2 == n {
			if s[s1L1] == goal[s2L1] {
				s1L1++
				s2L1++
			} else {
				return false
			}
		} else {

			if s[s1L1] == goal[s2L2] {
				s1L1++
			} else {
				s1L1 = 0
				s2L2start++
				s2L2 = s2L2start
			}

			s2L2++
		}

	}

	return true
}

func main() {
	//s := "abcde"
	//goal := "cdeab"

	s := "bbbacddceeb"
	goal := "ceebbbbacdd"
	res := rotateString(s, goal)

	fmt.Printf("res: %v\n", res)
}

package main

import "fmt"

func rotateString(s, goal string) bool {
	n := len(s)
	if n != len(goal) {
		return false
	}
next:
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[(i+j)%n] != goal[j] {
				continue next
			}
		}
		return true
	}
	return false
}

func main() {

	//s := "bbbacddceeb"
	//goal := "ceebbbbacdd"
	// 		  01234
	s := "abcde"
	// 		    01234
	goal := "cdeab"
	res := rotateString(s, goal)

	fmt.Printf("res: %v\n", res)
}

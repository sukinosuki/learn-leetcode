package main

import (
	"fmt"
)

func truncateSentence(s string, k int) string {

	//arr := strings.Split(s, " ")
	//
	//if k >= len(arr) {
	//	return s
	//}
	//
	//return strings.Join(arr[:k], " ")
	count := 0
	for i := 0; i < len(s); i++ {

		if s[i] == ' ' {
			count++
		}
		if count == k {
			return s[:i]
		}
	}

	return s
}

func main() {
	//s := "Hello how are you Contestant"
	//k := 4

	//s := "What is the solution to this problem"
	//k := 4
	s := "chopper is not a tanuki"
	k := 5
	res := truncateSentence(s, k)

	fmt.Printf("res: %v\n", res)
}

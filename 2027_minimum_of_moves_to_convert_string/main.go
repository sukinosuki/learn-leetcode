package main

import "fmt"

func minimumMoves(s string) int {

	ans := 0

	for i := 0; i < len(s); i += 3 {

		for i < len(s) && s[i] == 'O' {
			i++
		}

		if i < len(s) {
			ans++
		}
	}

	return ans
}

func main() {

	// 1
	//s := "XXX"
	s := "0XXX"
	res := minimumMoves(s)

	fmt.Printf("res: %v\n", res)
}

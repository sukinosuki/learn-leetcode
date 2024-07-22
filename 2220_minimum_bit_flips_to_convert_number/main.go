package main

import "fmt"

func minBitFlips(start int, goal int) int {

	ans := 0

	s := -1
	g := -1
	for start > 0 || goal > 0 {
		s = 0
		g = 0
		// 10: 1010
		// 7: 0111
		// 10: 0001010
		// 82: 1010010
		if start > 0 {
			s = start & 1
			start >>= 1
		}
		if goal > 0 {
			g = goal & 1
			goal >>= 1
		}

		if s != g {
			ans++
		}
	}

	return ans
}

func main() {

	// 3
	start := 10
	goal := 7
	// 3
	//start := 3
	//goal := 4
	// 3
	//start := 10
	//goal := 82
	res := minBitFlips(start, goal)

	fmt.Printf("res: %v\n", res)
}

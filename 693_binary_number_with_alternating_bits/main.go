package main

import "fmt"

func hasAlternatingBits(n int) bool {

	prev := n & 1
	n >>= 1
	for n > 0 {
		if n&1 == prev {
			return false
		}

		prev = n & 1
		n >>= 1
	}

	return true
}

func main() {
	//n := 5
	//n := 7
	n := 11
	res := hasAlternatingBits(n)

	fmt.Printf("res %v\n", res)
}

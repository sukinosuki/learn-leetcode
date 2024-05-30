package main

import "fmt"

func bitwiseComplement(n int) int {

	mask := 1
	for n > mask {
		// 构造出 11111111...
		mask = mask<<1 + 1
	}

	// 1001 ^ 1111
	return n ^ mask
}

func main() {
	// 2
	n := 5

	// 0
	//n := 7

	// 5
	//n := 10
	res := bitwiseComplement(n)

	fmt.Printf("res: %v\n", res)
}

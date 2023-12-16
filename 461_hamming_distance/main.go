package main

import "fmt"

func hammingDistance(x int, y int) int {

	n := x ^ y

	ans := 0
	for n > 0 {
		if n&1 == 1 {
			ans++
		}

		n = n >> 1
	}

	return ans
}

func main() {

	x := 1
	y := 3
	res := hammingDistance(x, y)

	fmt.Printf("res: %v\n", res)
}

package main

import "fmt"

func arrangeCoins(n int) int {

	res := helper(1, n)

	return res
}

func helper(stagger int, coins int) int {

	if coins < stagger {
		return 0
	}

	return helper(stagger+1, coins-stagger) + 1
}

func main() {
	n := 5
	res := arrangeCoins(n)

	fmt.Printf("n: %v\n", res)
}

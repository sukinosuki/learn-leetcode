package main

import "fmt"

func arrangeCoins(n int) int {

	stagger := 0
	acc := 0
	for acc+stagger < n {

		stagger++
		acc += stagger

	}

	return stagger
}

func main() {
	n := 5
	res := arrangeCoins(n)

	fmt.Printf("n: %v\n", res)
}

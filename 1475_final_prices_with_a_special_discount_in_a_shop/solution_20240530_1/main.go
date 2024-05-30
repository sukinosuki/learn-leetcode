package main

import "fmt"

func finalPrices(prices []int) []int {

	n := len(prices)
	ans := make([]int, n)
	stack := []int{0}

	for i := n - 1; i >= 0; i-- {

		for len(stack) > 1 && stack[len(stack)-1] > prices[i] {
			stack = stack[:len(stack)-1]
		}
		ans[i] = prices[i] - stack[len(stack)-1]
		stack = append(stack, prices[i])
	}

	return ans
}

func main() {
	prices := []int{8, 4, 6, 2, 3}
	res := finalPrices(prices)

	fmt.Printf("res: %v\n", res)
}

package main

import "fmt"

func finalPrices(prices []int) []int {

	n := len(prices)
	ans := make([]int, n)

	for i := range prices {
		ans[i] = prices[i]

		for j := i + 1; j < n; j++ {
			if prices[j] <= prices[i] {
				ans[i] -= prices[j]
				break
			}
		}
	}

	return ans
}

func main() {

	prices := []int{8, 4, 6, 2, 3}
	res := finalPrices(prices)

	fmt.Printf("res: %v\n", res)
}

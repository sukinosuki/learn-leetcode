package main

import "fmt"

func maxProfit(prices []int) int {
	minPrice := prices[0]
	ans := 0
	for _, price := range prices {
		ans = max(ans, price-minPrice)
		minPrice = min(minPrice, price)
	}

	return ans
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	prices := []int{7, 6, 4, 3, 1}
	//prices := []int{1, 2}
	index := maxProfit(prices)

	fmt.Printf("index: %d\n", index)
}

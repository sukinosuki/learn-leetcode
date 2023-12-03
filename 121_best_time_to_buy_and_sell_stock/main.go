package main

import "fmt"

func maxProfit(prices []int) (ans int) {

	minPrice := prices[0]
	for _, p := range prices {
		ans = max(ans, p-minPrice)
		minPrice = min(minPrice, p)
	}
	return
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

	//prices := []int{7, 1, 5, 3, 6, 4}
	//prices := []int{1, 4, 2}
	prices := []int{3, 2, 6, 5, 0, 3}
	//prices := []int{1, 2}
	index := maxProfit(prices)

	fmt.Printf("index: %d\n", index)
}

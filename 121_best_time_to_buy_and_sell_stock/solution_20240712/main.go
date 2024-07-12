package main

import "fmt"

func maxProfit(prices []int) int {

	prevMin := prices[0]

	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prevMin {
			ans = max(ans, prices[i]-prevMin)
		} else {
			prevMin = prices[i]
		}
	}

	return ans
}

func main() {

	prices := []int{7, 1, 5, 3, 6, 4}

	res := maxProfit(prices)

	fmt.Printf("res: %v\n", res)
}

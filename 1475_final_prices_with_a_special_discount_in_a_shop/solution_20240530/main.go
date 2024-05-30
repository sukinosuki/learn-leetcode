package main

import "fmt"

func finalPrices(prices []int) []int {
	n := len(prices)
	ans := make([]int, n)
	st := []int{0}

	for i := n - 1; i >= 0; i-- {
		p := prices[i]
		for len(st) > 1 && st[len(st)-1] > p {
			st = st[:len(st)-1]
		}
		ans[i] = p - st[len(st)-1]
		st = append(st, p)
	}

	return ans
}

func main() {
	prices := []int{8, 4, 6, 2, 3}
	res := finalPrices(prices)

	fmt.Printf("res: %v\n", res)
}

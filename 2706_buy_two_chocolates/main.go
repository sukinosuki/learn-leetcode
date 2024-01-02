package main

import (
	"fmt"
	"sort"
)

func buyChoco(prices []int, money int) int {

	sort.Ints(prices)

	sum := prices[0] + prices[1]
	sub := money - sum
	if sub >= 0 {
		return sub
	}
	return money
	//min := -1
	//
	//for i := 0; i <= len(prices)/2; i++ {
	//	for j := i + 1; j < len(prices); j++ {
	//		sum := prices[i] + prices[j]
	//
	//		if money-sum >= 0 && min < money-sum {
	//			min = money - sum
	//		}
	//
	//	}
	//}
	//
	//if min < 0 {
	//	return money
	//}
	//return min
}

func main() {
	nums := []int{1, 2, 4, 7}
	res := buyChoco(nums, 11)
	//nums := []int{3, 2, 3}
	//nums := []int{1, 2, 2}
	//res := buyChoco(nums, 3)

	fmt.Printf("res %v\n", res)
}

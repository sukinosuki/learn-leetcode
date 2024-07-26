package main

import (
	"fmt"
	"sort"
)

func fillCups(amount []int) int {

	sort.Ints(amount)
	ans := 0
	for amount[2] > 0 {
		amount[1]--
		amount[2]--
		ans++
		sort.Ints(amount)
	}

	return ans
}

func main() {
	// 4
	//amount := []int{1, 4, 2}
	amount := []int{5, 4, 4}
	res := fillCups(amount)
	fmt.Printf("res: %v\n", res)
}

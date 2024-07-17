package main

import (
	"fmt"
	"sort"
)

func minimumCost(cost []int) int {

	sort.Slice(cost, func(i, j int) bool {

		return cost[i] > cost[j]
	})

	ans := 0
	for i := 0; i < len(cost); i += 3 {

		ans += cost[i]
		if i+1 < len(cost) {
			ans += cost[i+1]
		}
	}

	return ans
}

func main() {

	// 5
	//cost := []int{1, 2, 3}
	cost := []int{6, 5, 7, 9, 2, 2}
	res := minimumCost(cost)

	fmt.Printf("res: %v\n", res)
}

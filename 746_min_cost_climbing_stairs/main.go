package main

import "fmt"

func minCostClimbingStairs(cost []int) int {

	n := len(cost)
	minCost := 0

	l2 := n - 1

	for l2 > 0 {
		if cost[l2] >= cost[l2-1] {
			minCost += cost[l2-1]
			l2 -= 2
		} else {
			minCost += cost[l2]
			l2--
		}
	}

	return minCost
}

func main() {
	//cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	//cost := []int{0, 0, 1, 1} // 1
	//cost := []int{0, 1, 2, 2} //2
	//cost := []int{10, 15, 20} // 15
	cost := []int{0, 2, 2, 1} // 15
	res := minCostClimbingStairs(cost)

	fmt.Printf("res: %v\n", res)
}

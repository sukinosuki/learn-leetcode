package main

import "fmt"

func minCostClimbingStairs(cost []int) int {

	n := len(cost)
	db := make([]int, n+1)

	for i := 2; i <= n; i++ {
		db[i] = min(db[i-2]+cost[i-2], db[i-1]+cost[i-1])
	}

	return db[n]
}

func min(n1, n2 int) int {
	if n1 > n2 {
		return n2
	}

	return n1
}

func main() {
	//				0  0    1  2  2  3    3  4  4    5
	//cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	//cost := []int{0, 0, 1, 1} // 1
	//cost := []int{0, 1, 2, 2} //2
	//cost := []int{10, 15, 20} // 15
	//cost := []int{0, 2, 2, 1} // 15
	res := minCostClimbingStairs(cost)

	fmt.Printf("res: %v\n", res)
}

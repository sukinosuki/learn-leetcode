package main

import "fmt"

func minCostClimbingStairs(cost []int) int {

	n := len(cost)
	db := make([]int, n+1)
	for i := 2; i <= n; i++ {

		db[i] = min(db[i-1]+cost[i-1], db[i-2]+cost[i-2])
	}

	return db[n]
}

func min(n1, n2 int) int {
	if n1 < n2 {

		return n1
	}
	return n2
}

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	// db         0    0  1  2  2  3    3  4  4    5  6
	//cost := []int{0, 0, 1, 1} // 1
	// db			0  0  0  0  1
	//cost := []int{0, 1, 2, 2} // 2
	// db			0  0  0  1  2
	//cost := []int{10, 15, 20} // 15
	// db			0   0   10  15
	//cost := []int{0, 2, 2, 1} // 2
	// db			0  0  0  2  2
	res := minCostClimbingStairs(cost)

	fmt.Printf("res: %v\n", res)
}

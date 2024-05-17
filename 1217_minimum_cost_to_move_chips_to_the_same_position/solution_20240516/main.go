package main

import "fmt"

func minCostToMoveChips(position []int) int {

	arr := make([]int, 2)
	for _, v := range position {
		arr[v%2]++
	}

	if arr[0] < arr[1] {
		return arr[0]
	}

	return arr[1]
}

func main() {
	// 1
	//position := []int{1, 10000}

	position := []int{1, 2, 3}
	res := minCostToMoveChips(position)

	fmt.Printf("res: %v\n", res)
}

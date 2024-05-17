package main

import "fmt"

func minCostToMoveChips(position []int) int {

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	position := []int{1, 10000}
	res := minCostToMoveChips(position)

	fmt.Printf("res: %v\n", res)
}
